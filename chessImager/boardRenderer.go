package chessImager

import (
	"github.com/fogleman/gg"
)

type boardRenderer struct {
	*Imager
}

func (r *boardRenderer) draw(c *gg.Context, _ ImageSettings) {
	size := r.settings.Board.Size

	// Set background to black color
	c.SetRGBA(toRGBA(r.settings.Board.black))
	c.DrawRectangle(r.border, r.border, float64(size), float64(size))
	c.Fill()

	c.SetRGBA(toRGBA(r.settings.Board.white))

	// Draw light squares
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				c.DrawRectangle(r.getSquareBounds(x, y))
				c.Fill()
			}
		}
	}
}
