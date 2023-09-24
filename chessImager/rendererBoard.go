package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererBoard struct {
	*Imager
}

func (r *rendererBoard) draw(c *gg.Context) {
	board := r.getDefaultBoardBox()

	// Set background to black color
	c.SetRGBA(toRGBA(r.settings.Board.Default.Black))
	c.DrawRectangle(board.Coords())
	c.Fill()

	// Draw light squares
	c.SetRGBA(toRGBA(r.settings.Board.Default.White))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				c.DrawRectangle(r.getSquareBox(x, y).Coords())
				c.Fill()
			}
		}
	}
}
