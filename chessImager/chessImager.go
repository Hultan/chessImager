package chessImager

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/fogleman/gg"
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
	return i, nil
}

func (i *Imager) GetImage(fen string) image.Image {
	return i.GetImageEx(fen, nil)
}

func (i *Imager) GetImageEx(fen string, s *Settings) image.Image {
	if !validateFen(fen) {
		panic(fmt.Errorf("invalid fen : %v", fen))
	}

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

	//convertColors(settings)
	i.settings = settings

	c := gg.NewContextForImage(image.NewRGBA(i.getBoardSize()))

	r := getRenderers(i, fen)
	for _, rend := range r {
		rend.draw(c)
	}

	return c.Image()
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
		Height: border - 3, // Vertical adjustment for letter g
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

func (i *Imager) setFontFace(c *gg.Context, size int) {
	// Set font face
	err := c.LoadFontFace("roboto.ttf", float64(size))
	if err != nil {
		panic("")
	}
	//font, err := truetype.Parse(goregular.TTF)
	//if err != nil {
	//	panic("")
	//}
	//face := truetype.NewFace(font, &truetype.Options{
	//	Size: float64(size),
	//})
	//c.SetFontFace(face)
}

// getRenderers returns a slice of all the renderers (in order of their importance).
func getRenderers(i *Imager, fen string) []renderer {
	return []renderer{
		&rendererBorder{i},
		&rendererBoard{i},
		&rendererRankAndFile{i},
		&rendererHighlightedSquare{i},
		&rendererPiece{i, fen},
		&rendererAnnotation{i},
		&rendererMoves{i},
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

	//convertColors(settings)

	return settings, nil
}
