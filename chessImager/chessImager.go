package chessImager

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"os"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type renderer interface {
	draw(*gg.Context, ImageSettings)
}

type Imager struct {
	settings *Settings
}

func NewImager() (*Imager, error) {
	s, err := loadDefaultSettings()
	if err != nil {
		panic(err)
	}

	i := &Imager{settings: s}
	i.loadPieces()
	return i, nil
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (i *Imager) loadPieces() {
	pieces = make(map[chessPiece]image.Image, 16)
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

	//saveImage(WhitePawn)
	//saveImage(WhiteBishop)
	//saveImage(WhiteKnight)
	//saveImage(WhiteRook)
	//saveImage(WhiteQueen)
	//saveImage(WhiteKing)
	//saveImage(BlackPawn)
	//saveImage(BlackBishop)
	//saveImage(BlackKnight)
	//saveImage(BlackRook)
	//saveImage(BlackQueen)
	//saveImage(BlackKing)
}

func (i *Imager) resize(img image.Image) image.Image {
	square := uint(i.settings.Board.Size) / 8
	return resize.Resize(square, square, img, resize.Lanczos3)
}
func saveImage(piece chessPiece) {
	path := getPath(piece)
	f, _ := os.Create(path)
	png.Encode(f, pieces[piece])
	f.Close()
}

func getPath(piece chessPiece) string {
	switch piece {
	case WhiteKing:
		return "/home/per/temp/whiteKing.png"
	case WhiteQueen:
		return "/home/per/temp/whiteQueen.png"
	case WhiteBishop:
		return "/home/per/temp/whiteBishop.png"
	case WhiteKnight:
		return "/home/per/temp/whiteKnight.png"
	case WhiteRook:
		return "/home/per/temp/whiteRook.png"
	case WhitePawn:
		return "/home/per/temp/whitePawn.png"
	case BlackKing:
		return "/home/per/temp/blackKing.png"
	case BlackQueen:
		return "/home/per/temp/blackQueen.png"
	case BlackBishop:
		return "/home/per/temp/blackBishop.png"
	case BlackKnight:
		return "/home/per/temp/blackKnight.png"
	case BlackRook:
		return "/home/per/temp/blackRook.png"
	case BlackPawn:
		return "/home/per/temp/blackPawn.png"
	default:
		return ""
	}
}

func (i *Imager) GetImage(settings ImageSettings) *image.RGBA {
	convertColorsForMove(&settings)

	r := getRenderers(i)

	im := image.NewRGBA(i.getSize())
	c := gg.NewContextForImage(im)

	for _, rend := range r {
		rend.draw(c, settings)
	}
	err := c.SavePNG("/home/per/temp/img.png")
	if err != nil {
		return nil
	}

	return im
}

// getSize returns a rectangle with the size of the board
// plus the border surrounding it.
func (i *Imager) getSize() image.Rectangle {
	size := i.settings.Board.Size + i.settings.Board.Border.Width*2

	return image.Rectangle{
		Max: image.Point{
			X: size,
			Y: size,
		},
	}
}

func (i *Imager) algToCoords(alg string) (int, int) {
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
	if i.settings.Board.Inverted {
		return invert(x, y)
	}
	return x, y
}

func (i *Imager) getRankBox(rank int) Rectangle {
	square := float64(i.settings.Board.Size) / 8
	border := float64(i.settings.Board.Border.Width)

	return Rectangle{
		X:      0,
		Y:      border + float64(7-rank)*square,
		Width:  border,
		Height: square,
	}
}

func (i *Imager) getFileBox(file int) Rectangle {
	square := float64(i.settings.Board.Size) / 8
	border := float64(i.settings.Board.Border.Width)

	return Rectangle{
		X:      border + float64(7-file)*square,
		Y:      border + 8*square,
		Width:  square,
		Height: border - 4, // Adjustment for letter g
	}
}

func (i *Imager) getSquareBox(x, y int) Rectangle {
	square := float64(i.settings.Board.Size) / 8
	border := float64(i.settings.Board.Border.Width)

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

// loadDefaultSettings loads the default settings from a json file
func loadDefaultSettings() (*Settings, error) {
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

// convertColors converts all color strings "#FF00BBFF" to color.RGBA
func convertColors(settings *Settings) {
	settings.Board.white = hexToRGBA(settings.Board.White)
	settings.Board.black = hexToRGBA(settings.Board.Black)
	settings.Board.Border.color = hexToRGBA(settings.Board.Border.Color)
	settings.Board.RankAndFile.color = hexToRGBA(settings.Board.RankAndFile.Color)
}

// convertColorsForMove converts all color strings "#FF00BBFF" to color.RGBA
func convertColorsForMove(settings *ImageSettings) {
	for i := range settings.Highlight {
		settings.Highlight[i].color = hexToRGBA(settings.Highlight[i].Color)
	}
}
