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

	r := getRenderers(i)
	for _, rend := range r {
		rend.draw(c)
	}

	return c.Image()
}

// getRenderers returns a slice of all the renderers (in order of their importance).
func getRenderers(i *Imager) []renderer {
	return []renderer{
		&rendererBorder{i},
		&rendererBoard{i},
		&rendererRankAndFile{i},
		&rendererHighlightedSquare{i},
		&rendererPiece{i},
		&rendererAnnotation{i},
		&rendererMoves{i},
	}
}
