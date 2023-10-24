package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererBoard struct {
	*Imager
}

func (r *rendererBoard) draw(c *gg.Context) error {
	board := r.getBoardBox()

	// Draw the entire board in the black color
	c.SetRGBA(toRGBA(settings.Board.Default.Black))
	c.DrawRectangle(board.Coords())
	c.Fill()

	// Draw the white squares, on top of the black board
	c.SetRGBA(toRGBA(settings.Board.Default.White))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				c.DrawRectangle(getSquareBox(x, y).Coords())
				c.Fill()
			}
		}
	}

	return nil
}

func (r *rendererBoard) getBoardBox() Rectangle {
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      border,
		Y:      border,
		Width:  float64(settings.Board.Default.Size),
		Height: float64(settings.Board.Default.Size),
	}
}
