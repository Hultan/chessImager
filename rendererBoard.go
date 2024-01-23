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
	ctx *ImageContext
	gg  *gg.Context
}

func (r *rendererBoard) draw() error {
	switch r.settings.Board.Type {
	case boardTypeDefault:
		r.drawDefault()
	case boardTypeImage:
		err := r.drawImage()
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid board type")
	}

	return nil
}

func (r *rendererBoard) drawDefault() {
	board := r.getBoardBox()

	// Draw the entire board in the black color
	r.gg.SetRGBA(r.settings.Board.Default.Black.toRGBA())
	r.gg.DrawRectangle(board.coords())
	r.gg.Fill()

	// Draw the white squares, on top of the black board
	r.gg.SetRGBA(r.settings.Board.Default.White.toRGBA())
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (y+x)%2 == 1 {
				r.gg.DrawRectangle(r.getSquareBox(x, y).coords())
				r.gg.Fill()
			}
		}
	}
}

func (r *rendererBoard) drawImage() error {
	f, err := os.Open(r.settings.Board.Image.Path)
	if err != nil {
		return fmt.Errorf("failed to load image : %v", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to encode image : %v", err)
	}

	r.gg.DrawImage(img, 0, 0)

	return nil
}
