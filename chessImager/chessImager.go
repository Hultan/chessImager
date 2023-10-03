package chessImager

import (
	"encoding/json"
	"fmt"
	"image"
	"os"

	"github.com/fogleman/gg"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type renderer interface {
	draw(*gg.Context)
}

type Imager struct {
	fen string
}

var settings *Settings

// NewImager creates a new Imager.
func NewImager() *Imager {
	return &Imager{}
}

func (i *Imager) Render(fen string) image.Image {
	return i.RenderEx(fen, nil)
}

func (i *Imager) RenderEx(fen string, ctx *Context) image.Image {
	if !validateFen(fen) {
		panic(fmt.Errorf("invalid fen : %v", fen))
	}

	// Handle settings
	var err error
	if ctx == nil {
		settings, err = loadSettings("")
		if err != nil {
			panic(err)
		}
	} else {
		settings = ctx.settings
	}

	i.fen = fen
	c := gg.NewContextForImage(image.NewRGBA(getBoardSize()))

	r := getRenderers(i, settings.Order)
	for _, rend := range r {
		rend.draw(c)
	}

	return c.Image()
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
		for i := 0; i < len(renderers); i++ {
			result = append(result, renderers[i])
		}
	} else {
		for _, o := range order {
			result = append(result, renderers[o])
		}
	}

	return result
}

// loadSettings loads the default settings from a json file
// Path : The path to load the settings from. Leave empty
// for the default settings (config/default.json).
func loadSettings(path string) (*Settings, error) {
	p := "config/default.json"
	if path != "" {
		p = path
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	settings := &Settings{}
	err = json.NewDecoder(f).Decode(settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
