package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererBoard struct {
	*Imager
}

func (r *rendererBoard) draw(c *gg.Context) {
	border := float64(r.settings.Border.Width)
	size := r.settings.Board.Default.Size

	// Set background to black color
	c.SetRGBA(toRGBA(r.settings.Board.Default.Black))
	c.DrawRectangle(border, border, float64(size), float64(size))
	c.Fill()

	c.SetRGBA(toRGBA(r.settings.Board.Default.White))

	// Draw light squares
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				c.DrawRectangle(r.getSquareBox(x, y).Coords())
				c.Fill()
			}
		}
	}
}
