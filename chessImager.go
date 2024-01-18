package chessImager

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

//go:embed config/default.json
var defaultSettings string

// Imager is the main struct that is used to create chess board images
type Imager struct {
	// Used to circumvent a bug in the fogleman/gg package, see
	// SetFontFace/LoadFontFace problem : https://github.com/fogleman/gg/pull/76
	useInternalFont bool
	settings        *Settings
}

// NewImager creates a new Imager.
func NewImager() *Imager {
	// We ignore the error here, since the default embedded settings file
	// should always be correct.
	s, _ := loadDefaultSettings()

	return &Imager{settings: s}
}

// NewImagerFromPath creates a new Imager using a user-defined JSON file.
func NewImagerFromPath(path string) (i *Imager, err error) {
	s, err := loadSettings(path)
	if err != nil {
		return nil, err
	}

	return &Imager{settings: s}, nil
}

// Render renders an image of a chess board based on a FEN string.
func (i *Imager) Render(fen string) (image.Image, error) {
	return i.RenderWithContext(&ImageContext{Fen: fen})
}

// RenderWithContext renders an image of a chess board based on an image context.
func (i *Imager) RenderWithContext(ctx *ImageContext) (image.Image, error) {
	size, err := i.getBoardSize()
	if err != nil {
		return nil, err
	}
	c := gg.NewContextForImage(image.NewRGBA(size))

	r, err := i.getRenderers()
	if err != nil {
		return nil, err
	}
	for _, rend := range r {
		err = rend.draw(c, ctx)
		if err != nil {
			return nil, err
		}
	}

	return c.Image(), nil
}

// NewContext creates a new image context, which can be used to:
// * Add the FEN string
// * Add highlighted squares
// * Add annotations
// * Add moves
func (i *Imager) NewContext() *ImageContext {
	return &ImageContext{}
}

// NewContextWithFEN creates a new image context with a FEN string set, which can be used to:
// * Add highlighted squares
// * Add annotations
// * Add moves
func (i *Imager) NewContextWithFEN(fen string) *ImageContext {
	return &ImageContext{Fen: fen}
}

// SetOrder can be used to set the render order.
func (i *Imager) SetOrder(order []int) error {
	if len(order) == 0 {
		order = []int{0, 1, 2, 3, 4, 5, 6}
	}

	if len(order) != 7 {
		return fmt.Errorf("len(order) must be 7")
	}

	i.settings.Order = order

	return nil
}

// getRenderers returns a slice of all the renderers in the given order
func (i *Imager) getRenderers() ([]renderer, error) {
	var result []renderer

	if len(i.settings.Order) != 7 {
		return result, fmt.Errorf("len(order) must be 7")
	}

	renderers := map[int]renderer{
		0: &rendererBorder{i},
		1: &rendererBoard{i},
		2: &rendererRankAndFile{i},
		3: &rendererHighlight{i},
		4: &rendererPiece{i},
		5: &rendererAnnotation{i},
		6: &rendererMoves{i},
	}

	for _, idx := range i.settings.Order {
		r := renderers[idx]
		if r == nil {
			return result, fmt.Errorf("invalid renderer index : %v", idx)
		}
		result = append(result, r)
	}

	return result, nil
}

// getBoardSize returns a rectangle with the size of the board
// plus the border surrounding it.
func (i *Imager) getBoardSize() (image.Rectangle, error) {
	switch i.settings.Board.Type {
	case boardTypeDefault:
		size := i.settings.Board.Default.Size + i.settings.Border.Width*2

		return image.Rectangle{
			Max: image.Point{
				X: size,
				Y: size,
			},
		}, nil
	case boardTypeImage:
		f, err := os.Open(i.settings.Board.Image.Path)
		if err != nil {
			return image.Rectangle{}, fmt.Errorf("failed to load image : %v", err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return image.Rectangle{}, fmt.Errorf("failed to encode image : %v", err)
		}

		return img.Bounds(), nil

	default:
		return image.Rectangle{}, fmt.Errorf("invalid board type : %v", i.settings.Board.Type)
	}
}

func (i *Imager) setFontFace(c *gg.Context, size int) error {
	if i.settings.FontStyle.Path == "" {
		// Use standard font
		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			return err
		}

		face := truetype.NewFace(font, &truetype.Options{Size: float64(size)})
		c.SetFontFace(face)
		i.useInternalFont = true
	} else {
		// Load font specified in config file
		err := c.LoadFontFace(i.settings.FontStyle.Path, float64(size))
		if err != nil {
			return fmt.Errorf("failed to load font face : %v", err)
		}
		i.useInternalFont = false
	}

	return nil
}

func (i *Imager) getBoardBox() Rectangle {
	switch i.settings.Board.Type {
	case boardTypeDefault:
		border := float64(i.settings.Border.Width)
		size := float64(i.settings.Board.Default.Size)

		return Rectangle{
			X:      border,
			Y:      border,
			Width:  size,
			Height: size,
		}
	case boardTypeImage:
		return i.settings.Board.Image.Rect
	default:
		panic("invalid board type")
	}
}

func (i *Imager) getSquareBox(x, y int) Rectangle {
	board := i.getBoardBox()
	square := board.Width / 8

	var dx, dy float64
	switch i.settings.Board.Type {
	case boardTypeDefault:
		border := float64(i.settings.Border.Width)
		dx, dy = border, border
	case boardTypeImage:
		dx, dy = board.X, board.Y
	default:
		panic("invalid board type")
	}

	return Rectangle{
		X:      dx + float64(x)*square,
		Y:      dy + float64(invert(y))*square,
		Width:  square,
		Height: square,
	}
}

// loadSettings loads the settings from a json file
// Path : The path to load the settings from.
func loadSettings(path string) (*Settings, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return decodeSettings(f)
}

// loadDefaultSettings loads the embedded default settings
func loadDefaultSettings() (*Settings, error) {
	r := strings.NewReader(defaultSettings)

	return decodeSettings(r)
}

// decodeSettings decode the string/file and returns a Settings object and an error
func decodeSettings(r io.Reader) (*Settings, error) {
	s := &Settings{}
	err := json.NewDecoder(r).Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
