package chessImager

import (
	"encoding/json"
	"image"
	"os"

	"github.com/fogleman/gg"
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
	return &Imager{settings: s}, nil
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
		return i.invert(x, y)
	}
	return x, y
}

// TODO : implement a getRankBox and getFileBox function
// and then simplify rankAndFileRenderer.
// TODO : remove Imager.square and border
// TODO : getSquareBounds should return a Rectangle
// TODO : Rectangle should have the toCoords() method
func (i *Imager) getSquareBounds(x, y int) (float64, float64, float64, float64) {
	square := float64(i.settings.Board.Size) / 8
	border := float64(i.settings.Board.Border.Width)

	return border + float64(x)*square, border + float64(7-y)*square, square, square
}

func (i *Imager) invert(x, y int) (int, int) {
	return 7 - x, 7 - y
}

// getRenderers returns a slice of all the renderers (in order of their importance).
func getRenderers(i *Imager) []renderer {
	return []renderer{
		&borderRenderer{i},
		&boardRenderer{i},
		&rankAndFileRenderer{i},
		&highlightedSquareRenderer{i},
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

// convertColors converts all color strings "#FF00BBFF" to color.RGBA
func convertColorsForMove(settings *ImageSettings) {
	for i := range settings.Highlight {
		settings.Highlight[i].color = hexToRGBA(settings.Highlight[i].Color)
	}
}
