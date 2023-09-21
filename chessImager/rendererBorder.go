package chessImager

import "github.com/fogleman/gg"

type rendererBorder struct {
	*Imager
}

func (r *rendererBorder) draw(c *gg.Context) {
	// Set background color to border color
	c.SetRGBA(toRGBA(r.settings.Border.color))
	c.Clear()
}
