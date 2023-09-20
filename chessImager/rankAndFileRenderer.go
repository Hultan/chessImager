package chessImager

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type rankAndFileRenderer struct {
	*Imager
}

func (r *rankAndFileRenderer) draw(c *gg.Context, _ ImageSettings) {
	var dx, dy float64
	square := float64(r.settings.Board.Size) / 8
	border := float64(r.settings.Board.Border.Width)
	size := r.settings.Board.RankAndFile.Size
	color := r.settings.Board.RankAndFile.color

	switch r.settings.Board.RankAndFile.Type {
	case RankAndFileNone:
		return
	case RankAndFileInBorder:
		if border < 10 {
			return
		}
	case RankAndFileInSquares:
		// TODO : Should use r.getSquareBounds() instead
		if border < 10 {
			return
		}
		dx, dy = (square-border)/2, -border*1.15
	}

	// Set font face
	c.SetRGBA(toRGBA(color))
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(size),
	})
	c.SetFontFace(face)

	r.drawRanksAndFiles(c, border, square, dx, dy)
}

func (r *rankAndFileRenderer) drawRanksAndFiles(c *gg.Context, border, square, dx, dy float64) {
	var text string

	// Rank
	for i := 0; i < 8; i++ {
		if r.settings.Board.Inverted {
			text = fmt.Sprintf("%d", i+1)
		} else {
			text = fmt.Sprintf("%d", 8-i)
		}
		tw, th := c.MeasureString(text)
		x := (border - tw) / 2
		y := border + square*float64(i+1) - (square-th)/2
		c.DrawString(text, x+dx, y+dy)
	}
	// File
	for i := 0; i < 8; i++ {
		if r.settings.Board.Inverted {
			text = fmt.Sprintf("%c", 'h'-i)
		} else {
			text = fmt.Sprintf("%c", 'a'+i)
		}
		tw, th := c.MeasureString(text)
		x := border + square*float64(i) + (square-tw)/2
		y := border*1.85 + square*8 - (border-th)/2
		c.DrawString(text, x+dx, y+dy)
	}
}
