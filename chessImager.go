package chessImager

import (
	_ "embed"
	"errors"
	"fmt"
	"image"

	"github.com/fogleman/gg"
)

//go:embed config/default.json
var defaultSettings string

type Imager struct {
	fen string
	ctx *Context
}

var settings *Settings

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

	i = &Imager{}

	return
}

func (i *Imager) Render(fen string) (image.Image, error) {
	return i.RenderEx(fen, nil)
}

func (i *Imager) RenderEx(fen string, ctx *Context) (image.Image, error) {
	var err error

	if !validateFen(fen) {
		return nil, fmt.Errorf("invalid fen : %v", fen)
	}

	i.fen = fen
	i.ctx = ctx
	c := gg.NewContextForImage(image.NewRGBA(getBoardSize()))

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

func (i *Imager) NewContext() *Context {
	return &Context{}
}

func (i *Imager) SetOrder(order []int) error {
	if len(order) != 7 {
		return errors.New("len(order) must be 7")
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
		return result, errors.New("len(order) must be 7")
	}

	for _, idx := range settings.Order {
		r := renderers[idx]
		if r == nil {
			return result, fmt.Errorf("no renderer with index : %d", idx)
		}
		result = append(result, r)
	}

	return result, nil
}
