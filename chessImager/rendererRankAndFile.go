package chessImager

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type rendererRankAndFile struct {
	*Imager
}

type RankFile struct {
	box  Rectangle
	text string
}

func (r *rendererRankAndFile) draw(c *gg.Context, _ ImageSettings) {
	var dx, dy float64 // InSquare adjustments

	square := float64(r.settings.Board.Size) / 8
	border := float64(r.settings.Board.Border.Width)
	size := r.settings.Board.RankAndFile.Size

	switch r.settings.Board.RankAndFile.Type {
	case RankAndFileNone:
		return
	case RankAndFileInBorder:
		if border < 10 {
			return
		}
	case RankAndFileInSquares:
		if border < 10 {
			return
		}
		dx, dy = (square-border)/2, -border
	}

	c.SetRGBA(toRGBA(r.settings.Board.RankAndFile.color))
	r.setFontFace(c, size)
	r.drawRanksAndFiles(c, dx, dy)
}

func (r *rendererRankAndFile) drawRanksAndFiles(c *gg.Context, dx, dy float64) {
	rf := r.getRFBoxes()

	for _, r := range rf {
		tw, th := c.MeasureString(r.text)
		x := r.box.X + (r.box.Width-tw)/2
		y := r.box.Y + (r.box.Height-th)/2 + th
		c.DrawString(r.text, x+dx, y+dy)
	}
}

func (r *rendererRankAndFile) getRFBoxes() []RankFile {
	var rf []RankFile

	for i := 0; i < 8; i++ {
		// Ranks
		text := r.getRankText(i)
		box := r.getRankBox(i)
		rf = append(rf, RankFile{box: box, text: text})

		// Files
		text = r.getFileText(i)
		box = r.getFileBox(i)
		rf = append(rf, RankFile{box: box, text: text})
	}
	return rf
}

func (r *rendererRankAndFile) getRankText(n int) string {
	if r.settings.Board.Inverted {
		return fmt.Sprintf("%d", 8-n)
	} else {
		return fmt.Sprintf("%d", n+1)
	}
}

func (r *rendererRankAndFile) getFileText(n int) string {
	if r.settings.Board.Inverted {
		return fmt.Sprintf("%c", 'a'+n)
	} else {
		return fmt.Sprintf("%c", 'h'-n)
	}
}

func (r *rendererRankAndFile) setFontFace(c *gg.Context, size int) {
	// Set font face
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(size),
	})
	c.SetFontFace(face)
}
