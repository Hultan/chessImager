package chessImager

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

//go:embed config/default.json
var defaultSettings string

type Imager struct {
	fen string
	ctx *Context
}

var settings *Settings
var boardImage image.Image

// Used to circumvent a bug in the fogleman/gg package, see
// SetFontFace/LoadFontFace problem : https://github.com/fogleman/gg/pull/76
var useInternalFont = true

// NewImager creates a new Imager.
func NewImager() *Imager {
	settings = loadDefaultSettings()
	return &Imager{}
}

// NewImagerFromPath creates a new Imager using a user-defined JSON file.
func NewImagerFromPath(path string) (i *Imager, err error) {
	settings, err = loadSettings(path)
	if err != nil {
		return nil, err
	}

	if err = validateSettings(); err != nil {
		return nil, err
	}

	i = &Imager{}

	return
}

// Render renders an image of a chess board based on a FEN string.
func (i *Imager) Render(fen string) (image.Image, error) {
	return i.RenderEx(fen, nil)
}

// RenderEx renders an image of a chess board based on a FEN string and a context.
func (i *Imager) RenderEx(fen string, ctx *Context) (image.Image, error) {
	var err error

	if !validateFen(fen) {
		return nil, fmt.Errorf("invalid fen : %v", fen)
	}

	i.fen = fen
	i.ctx = ctx
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
		err = rend.draw(c)
		if err != nil {
			return nil, err
		}
	}

	return c.Image(), nil
}

// NewContext creates a new context, which can be used to:
// * Add highlighted squares
// * Add annotations
// * Add moves
func (i *Imager) NewContext() *Context {
	return &Context{}
}

// SetOrder can be used to set the render order.
func (i *Imager) SetOrder(order []int) error {
	if len(order) == 0 {
		order = []int{0, 1, 2, 3, 4, 5, 6}
	}

	if len(order) != 7 {
		return fmt.Errorf("len(order) must be 7")
	}

	settings.Order = order

	return nil
}

// getRenderers returns a slice of all the renderers in the given order
func (i *Imager) getRenderers() ([]renderer, error) {
	var result []renderer

	renderers := map[int]renderer{
		0: &rendererBorder{i},
		1: &rendererBoard{i},
		2: &rendererRankAndFile{i},
		3: &rendererHighlight{i},
		4: &rendererPiece{i},
		5: &rendererAnnotation{i},
		6: &rendererMoves{i},
	}

	if len(settings.Order) != 7 {
		return result, fmt.Errorf("len(order) must be 7")
	}

	for _, idx := range settings.Order {
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
	switch settings.Board.Type {
	case boardTypeDefault:
		size := settings.Board.Default.Size + settings.Border.Width*2

		return image.Rectangle{
			Max: image.Point{
				X: size,
				Y: size,
			},
		}, nil
	case boardTypeImage:
		return boardImage.Bounds(), nil

	default:
		return image.Rectangle{}, fmt.Errorf("invalid board type : %v", settings.Board.Type)
	}
}

// loadSettings loads the settings from a json file
// Path : The path to load the settings from.
func loadSettings(path string) (*Settings, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	s := &Settings{}
	err = json.NewDecoder(f).Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// loadDefaultSettings loads the embedded default settings
func loadDefaultSettings() *Settings {
	r := strings.NewReader(defaultSettings)

	s := &Settings{}
	// Ok to panic here, the embedded settings should always be correct
	err := json.NewDecoder(r).Decode(s)
	if err != nil {
		panic(err)
	}

	return s
}

// validateSettings validates some of the values in the JSON file
func validateSettings() error {
	if settings.Board.Type == boardTypeImage {
		if err := tryLoadImage(settings.Board.Image.Path, &boardImage); err != nil {
			return err
		}
	}

	if settings.Pieces.Type == piecesTypeImageMap {
		var img image.Image
		if err := tryLoadImage(settings.Pieces.ImageMap.Path, &img); err != nil {
			return err
		}
	}

	if settings.Pieces.Type == piecesTypeImages {
		var img image.Image
		for _, p := range settings.Pieces.Images.Pieces {
			if err := tryLoadImage(p.Path, &img); err != nil {
				return err
			}
		}
	}

	if settings.FontStyle.Path != "" {
		if err := tryLoadFile(settings.FontStyle.Path); err != nil {
			return err
		}
	}

	return nil
}

// tryLoadImage tries to load the specified image, makes sure it exists,
// and is an image.
func tryLoadImage(path string, img *image.Image) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to load image : %v", err)
	}
	defer f.Close()

	*img, _, err = image.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to encode image : %v", err)
	}

	return nil
}

// tryLoadFile tries to load the specified file, makes sure it exists.
func tryLoadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to file (%s) : %v", path, err)
	}
	defer f.Close()

	return nil
}
