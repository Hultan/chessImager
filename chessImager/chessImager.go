package chessImager

import (
	"bytes"
	"encoding/json"
	"image"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type renderer interface {
	draw(*gg.Context)
}

type Imager struct {
	settings *Settings
}

func NewImager() (*Imager, error) {
	settings, err := GetSettings()
	if err != nil {
		panic(err)
	}
	i := &Imager{settings}
	i.loadEmbeddedPieces()
	return i, nil
}

func (i *Imager) GetImage(fen string) image.Image {
	return i.GetImageEx(fen, nil)
}

func (i *Imager) GetImageEx(fen string, s *Settings) image.Image {
	// TODO: Validate and normalize FEN
	var settings *Settings
	var err error

	if s == nil {
		settings, err = GetSettings()
		if err != nil {
			panic(err)
		}
	} else {
		settings = s
	}

	convertColors(settings)
	i.settings = settings

	c := gg.NewContextForImage(image.NewRGBA(i.getBoardSize()))

	r := getRenderers(i)
	for _, rend := range r {
		rend.draw(c)
	}

	return c.Image()
}

func (i *Imager) loadEmbeddedPieces() {
	pieces = make(map[chessPiece]image.Image, 12)
	img, _, err := image.Decode(bytes.NewReader(defaultPieces))
	if err != nil {
		panic(err)
	}
	sub, ok := img.(SubImager)
	if !ok {
		panic("Failed to create SubImager")
	}
	pieces[WhiteKing] = i.resize(sub.SubImage(image.Rect(0, 0, 333, 333)))
	pieces[WhiteQueen] = i.resize(sub.SubImage(image.Rect(333, 0, 666, 333)))
	pieces[WhiteBishop] = i.resize(sub.SubImage(image.Rect(666, 0, 999, 333)))
	pieces[WhiteKnight] = i.resize(sub.SubImage(image.Rect(999, 0, 1332, 333)))
	pieces[WhiteRook] = i.resize(sub.SubImage(image.Rect(1332, 0, 1665, 333)))
	pieces[WhitePawn] = i.resize(sub.SubImage(image.Rect(1665, 0, 1998, 333)))
	pieces[BlackKing] = i.resize(sub.SubImage(image.Rect(0, 333, 333, 666)))
	pieces[BlackQueen] = i.resize(sub.SubImage(image.Rect(333, 333, 666, 666)))
	pieces[BlackBishop] = i.resize(sub.SubImage(image.Rect(666, 333, 999, 666)))
	pieces[BlackKnight] = i.resize(sub.SubImage(image.Rect(999, 333, 1332, 666)))
	pieces[BlackRook] = i.resize(sub.SubImage(image.Rect(1332, 333, 1665, 666)))
	pieces[BlackPawn] = i.resize(sub.SubImage(image.Rect(1665, 333, 1998, 666)))
}

func (i *Imager) resize(img image.Image) image.Image {
	var square uint
	if i.settings.Board.Type == BoardTypeDefault {
		square = uint(i.settings.Board.Default.Size) / 8
	}
	return resize.Resize(square, square, img, resize.Lanczos3)
}

// getBoardSize returns a rectangle with the size of the board
// plus the border surrounding it.
func (i *Imager) getBoardSize() image.Rectangle {
	size := i.settings.Board.Default.Size + i.settings.Border.Width*2

	return image.Rectangle{
		Max: image.Point{
			X: size,
			Y: size,
		},
	}
}

func (i *Imager) algToCoords(alg string) (int, int) {
	alg = strings.ToLower(alg)
	if len(alg) != 2 {
		panic("invalid length of alg")
	}
	if alg[0] < 'a' || alg[0] > 'h' {
		panic("invalid character in alg : " + string(alg[0]))
	}
	if alg[1] < '1' || alg[1] > '8' {
		panic("invalid character in alg : " + string(alg[1]))
	}
	x, y := int(alg[0]-'a'), int(alg[1]-'1')
	if i.settings.Board.Default.Inverted {
		return invert(x, y)
	}
	return x, y
}

func (i *Imager) getRankBox(rank int) Rectangle {
	square := float64(i.settings.Board.Default.Size) / 8
	border := float64(i.settings.Border.Width)

	return Rectangle{
		X:      0,
		Y:      border + float64(7-rank)*square,
		Width:  border,
		Height: square,
	}
}

func (i *Imager) getFileBox(file int) Rectangle {
	square := float64(i.settings.Board.Default.Size) / 8
	border := float64(i.settings.Border.Width)

	return Rectangle{
		X:      border + float64(7-file)*square,
		Y:      border + 8*square,
		Width:  square,
		Height: border - 4, // Adjustment for letter g
	}
}

func (i *Imager) getSquareBox(x, y int) Rectangle {
	square := float64(i.settings.Board.Default.Size) / 8
	border := float64(i.settings.Border.Width)

	return Rectangle{
		X:      border + float64(x)*square,
		Y:      border + float64(7-y)*square,
		Width:  square,
		Height: square,
	}
}

// getRenderers returns a slice of all the renderers (in order of their importance).
func getRenderers(i *Imager) []renderer {
	return []renderer{
		&rendererBorder{i},
		&rendererBoard{i},
		&rendererRankAndFile{i},
		&rendererHighlightedSquare{i},
		&rendererPiece{i},
	}
}

// GetSettings loads the default settings from a json file
func GetSettings() (*Settings, error) {
	f, err := os.Open("config/default.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	settings := &Settings{}
	err = json.NewDecoder(f).Decode(settings)
	if err != nil {
		return nil, err
	}

	convertColors(settings)

	return settings, nil
}
