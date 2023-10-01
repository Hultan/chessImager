package chessImager

import (
	"fmt"
	"image"

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

func (i *Imager) GetImage(fen string) image.Image {
	return i.GetImageEx(fen, nil)
}

func (i *Imager) GetImageEx(fen string, s *Settings) image.Image {
	if !validateFen(fen) {
		panic(fmt.Errorf("invalid fen : %v", fen))
	}

	// Handle settings
	var err error
	if s == nil {
		settings, err = GetSettings("")
		if err != nil {
			panic(err)
		}
	} else {
		settings = s
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
		3: &rendererHighlightedSquare{i},
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
