package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererBoard struct {
	*Imager
}

func (r *rendererBoard) draw(c *gg.Context) error {
	switch settings.Board.Type {
	case BoardTypeDefault:
		r.drawDefault(c)
	case BoardTypeImage:
		r.drawImage(c)
	default:
		return errors.New("invalid board type")
	}

	return nil
}

func (r *rendererBoard) drawDefault(c *gg.Context) {
	board := getBoardBox()

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
}

func (r *rendererBoard) drawImage(c *gg.Context) {
	c.DrawImage(boardImage, 0, 0)
}
