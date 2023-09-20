package chessImager

import "github.com/fogleman/gg"

type rendererBorder struct {
	*Imager
}

func (r *rendererBorder) draw(c *gg.Context, _ ImageSettings) {
	// Set background color to border color
	c.SetRGBA(toRGBA(r.settings.Board.Border.color))
	c.Clear()
}
