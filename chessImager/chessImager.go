package chessImager

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
)

type Imager struct {
	fen string
	ctx *Context
}

var settings *Settings

// NewImager creates a new Imager.
func NewImager() *Imager {
	return &Imager{}
}

func (i *Imager) Render(fen string) (image.Image, error) {
	return i.RenderEx(fen, nil)
}

func (i *Imager) RenderEx(fen string, ctx *Context) (image.Image, error) {
	if !validateFen(fen) {
		return nil, fmt.Errorf("invalid fen : %v", fen)
	}

	// Handle settings
	var err error
	if ctx == nil {
		settings, err = loadSettings("")
		if err != nil {
			return nil, err
		}
	} else {
		settings = ctx.settings
	}

	i.fen = fen
	i.ctx = ctx
	c := gg.NewContextForImage(image.NewRGBA(getBoardSize()))

	r := getRenderers(i, settings.Order)
	for _, rend := range r {
		rend.draw(c)
	}

	return c.Image(), nil
}

func (i *Imager) NewContext() (*Context, error) {
	return i.NewContextFromPath("")
}

func (i *Imager) NewContextFromPath(path string) (*Context, error) {
	s, err := loadSettings(path)
	if err != nil {
		return nil, err
	}
	return &Context{settings: s}, nil
}

// getRenderers returns a slice of all the renderers in the given order
func getRenderers(i *Imager, order []int) []renderer {
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

	if order == nil {
		order = []int{0, 1, 2, 3, 4, 5, 6}
	}

	for _, i := range order {
		r := renderers[i]
		if r == nil {
			panic(fmt.Errorf("no renderer with index : %d", i))
		}
		result = append(result, r)
	}

	return result
}
