package chessImager

import "github.com/fogleman/gg"

type rendererBorder struct {
	*Imager
}

func (r *rendererBorder) draw(c *gg.Context, _ *ImageContext) error {
	if r.settings.Board.Type == boardTypeImage {
		return nil
	}

	// Set background color to border color
	c.SetRGBA(r.settings.Border.Color.toRGBA())
	c.Clear()

	return nil
}
