package chessImager

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/exp/constraints"
	"golang.org/x/image/font/gofont/goregular"
)

//go:embed config/default.json
var defaultSettings string

type renderer interface {
	draw() error
}

// Imager is the main struct that is used to create chess board images
type Imager struct {
	// Used to circumvent a bug in the fogleman/gg package, see
	// SetFontFace/LoadFontFace problem : https://github.com/fogleman/gg/pull/76
	useInternalFont bool
	settings        *Settings
	inverted        bool
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

// LoadSettings loads in a new settings file.
func (i *Imager) LoadSettings(path string) error {
	s, err := loadSettings(path)
	if err != nil {
		return err
	}

	i.settings = s

	return nil
}

// Render renders an image of a chess board based on a FEN string.
func (i *Imager) Render(fen string) (image.Image, error) {
	return i.RenderWithContext(&ImageContext{Fen: fen})
}

// RenderInverted renders an image of an inverted chess board based on a FEN string.
func (i *Imager) RenderInverted(fen string) (image.Image, error) {
	return i.RenderWithContextInverted(&ImageContext{Fen: fen})
}

// RenderWithContext renders an image of a chess board based on an image context.
func (i *Imager) RenderWithContext(ctx *ImageContext) (image.Image, error) {
	i.inverted = false
	return i.renderWithContext(ctx)
}

// RenderWithContextInverted renders an image of an inverted chess board based on an image context.
func (i *Imager) RenderWithContextInverted(ctx *ImageContext) (image.Image, error) {
	i.inverted = true
	return i.renderWithContext(ctx)
}

// renderWithContext renders an image of a chess board based on an image context.
func (i *Imager) renderWithContext(ctx *ImageContext) (image.Image, error) {
	if ok := validateFen(ctx.Fen); !ok {
		return nil, fmt.Errorf("invalid fen: %v", ctx.Fen)
	}

	size, err := i.getBoardSize()
	if err != nil {
		return nil, err
	}
	c := gg.NewContextForImage(image.NewRGBA(size))

	r, err := i.getRenderers(c, ctx)
	if err != nil {
		return nil, err
	}
	for _, rend := range r {
		err = rend.draw()
		if err != nil {
			return nil, err
		}
	}

	return c.Image(), nil
}

// NewContext creates a new image context, which can be used to:
// * Add highlighted squares
// * Add annotations
// * Add moves
func (i *Imager) NewContext(fen string) *ImageContext {
	return &ImageContext{Fen: fen}
}

// SetOrder can be used to set the render order.
func (i *Imager) SetOrder(order []int) error {
	if len(order) == 0 {
		order = []int{0, 1, 2, 3, 4, 5, 6}
	}

	err := i.validateOrder(order)
	if err != nil {
		return err
	}

	i.settings.Order = order

	return nil
}

func (i *Imager) validateOrder(order []int) error {
	if len(order) != 7 {
		return fmt.Errorf("len(order) must be 7")
	}

	var index = map[int]bool{}

	for _, i := range order {
		if i < 0 || i > 6 {
			return fmt.Errorf("invalid renderer index")
		}
		if ok := index[i]; ok {
			return fmt.Errorf("renderer index added twice")
		}
		index[i] = true
	}
	return nil
}

// getRenderers returns a slice of all the renderers in the given order
func (i *Imager) getRenderers(gg *gg.Context, ctx *ImageContext) ([]renderer, error) {
	var result []renderer

	err := i.validateOrder(i.settings.Order)
	if err != nil {
		return nil, err
	}

	renderers := map[int]renderer{
		0: &rendererBorder{Imager: i, ctx: ctx, gg: gg},
		1: &rendererBoard{Imager: i, ctx: ctx, gg: gg},
		2: &rendererRankAndFile{Imager: i, ctx: ctx, gg: gg},
		3: &rendererHighlight{Imager: i, ctx: ctx, gg: gg},
		4: &rendererPiece{Imager: i, ctx: ctx, gg: gg},
		5: &rendererAnnotation{Imager: i, ctx: ctx, gg: gg},
		6: &rendererMoves{Imager: i, ctx: ctx, gg: gg},
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

// hexToRGBA converts a hex string to a color
// #RRGGBBAA or #RRGGBB or RRGGBBAA or RRGGBB
func hexToRGBA(hex string) (col color.RGBA, err error) {
	// Remove leading '#' if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Add missing alpha
	if len(hex) == 6 {
		hex += "FF"
	}

	// Parse the hex values for red, green, blue and alpha
	if len(hex) == 8 {
		_, err = fmt.Sscanf(hex, "%02x%02x%02x%02x", &col.R, &col.G, &col.B, &col.A)
		if err != nil {
			return col, fmt.Errorf("invalid color (%s) : %v", hex, err)
		}
	} else {
		err := errors.New("valid formats : #RRGGBBAA or #RRGGBB or RRGGBBAA or RRGGBB")
		return col, fmt.Errorf("invalid color (%s) : %v", hex, err)
	}

	return col, nil
}

func invert(x int) int {
	return 7 - x
}

func abs[T constraints.Float | constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(v int) int {
	switch {
	case v < 0:
		return -1
	case v == 0:
		return 0
	default:
		return 1
	}
}
