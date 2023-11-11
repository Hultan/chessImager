package chessImager

import "github.com/fogleman/gg"

type rendererBorder struct {
	*Imager
}

func (r *rendererBorder) draw(c *gg.Context) error {
	if settings.Board.Type == boardTypeImage {
		return nil
	}

	// Set background color to border color
	c.SetRGBA(settings.Border.Color.toRGBA())
	c.Clear()

	return nil
}
