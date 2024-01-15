package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererBoard struct {
	*Imager
}

func (r *rendererBoard) draw(c *gg.Context) error {
	switch r.settings.Board.Type {
	case boardTypeDefault:
		r.drawDefault(c)
	case boardTypeImage:
		r.drawImage(c)
	default:
		return errors.New("invalid board type")
	}

	return nil
}

func (r *rendererBoard) drawDefault(c *gg.Context) {
	board := r.settings.getBoardBox()

	// Draw the entire board in the black color
	c.SetRGBA(r.settings.Board.Default.Black.toRGBA())
	c.DrawRectangle(board.coords())
	c.Fill()

	// Draw the white squares, on top of the black board
	c.SetRGBA(r.settings.Board.Default.White.toRGBA())
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				c.DrawRectangle(r.settings.getSquareBox(x, y).coords())
				c.Fill()
			}
		}
	}
}

func (r *rendererBoard) drawImage(c *gg.Context) {
	c.DrawImage(r.boardImage, 0, 0)
}
