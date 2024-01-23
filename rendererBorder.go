package chessImager

import "github.com/fogleman/gg"

type rendererBorder struct {
	*Imager
	ctx *ImageContext
	gg  *gg.Context
}

func (r *rendererBorder) draw() error {
	if r.settings.Board.Type == boardTypeImage {
		return nil
	}

	// Set background color to border color
	r.gg.SetRGBA(r.settings.Border.Color.toRGBA())
	r.gg.Clear()

	return nil
}
