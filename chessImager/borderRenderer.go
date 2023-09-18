package chessImager

import "github.com/fogleman/gg"

type borderRenderer struct {
	*Imager
}

func (r *borderRenderer) draw(c *gg.Context, _ ImageSettings) {
	// Set background color to border color
	c.SetRGBA(toRGBA(r.settings.Board.Border.color))
	c.Clear()
}
