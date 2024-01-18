package chessImager

import (
	"errors"
	"fmt"
	"image"
	"os"

	"github.com/fogleman/gg"
)

type rendererBoard struct {
	*Imager
}

func (r *rendererBoard) draw(c *gg.Context, _ *ImageContext) error {
	switch r.settings.Board.Type {
	case boardTypeDefault:
		r.drawDefault(c)
	case boardTypeImage:
		err := r.drawImage(c)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid board type")
	}

	return nil
}

func (r *rendererBoard) drawDefault(c *gg.Context) {
	board := r.getBoardBox()

	// Draw the entire board in the black color
	c.SetRGBA(r.settings.Board.Default.Black.toRGBA())
	c.DrawRectangle(board.coords())
	c.Fill()

	// Draw the white squares, on top of the black board
	c.SetRGBA(r.settings.Board.Default.White.toRGBA())
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				c.DrawRectangle(r.getSquareBox(x, y).coords())
				c.Fill()
			}
		}
	}
}

func (r *rendererBoard) drawImage(c *gg.Context) error {
	f, err := os.Open(r.settings.Board.Image.Path)
	if err != nil {
		return fmt.Errorf("failed to load image : %v", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to encode image : %v", err)
	}

	c.DrawImage(img, 0, 0)

	return nil
}
