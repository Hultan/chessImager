package chessImager

import (
	"errors"
	"fmt"

	"github.com/fogleman/gg"
)

const borderLimit = 10

type rendererRankAndFile struct {
	*Imager
}

type RankFile struct {
	box  Rectangle
	text string
}

func (r *rendererRankAndFile) draw(c *gg.Context) error {
	if settings.Board.Type == BoardTypeImage {
		return nil
	}

	var dx, dy float64 // InSquare adjustments

	square := float64(settings.Board.Default.Size) / 8
	border := float64(settings.Border.Width)
	size := settings.RankAndFile.FontSize

	switch settings.RankAndFile.Type {
	case RankAndFileTypeNone:
		return nil
	case RankAndFileTypeInBorder:
		// Don't bother drawing ranks and files when to border is too thin
		if border < borderLimit {
			return nil
		}
	case RankAndFileTypeInSquares:
		// Don't bother drawing ranks and files when to border is too thin
		if border < borderLimit {
			return nil
		}
		dx, dy = (square-border)/2, -border
	default:
		return errors.New("invalid rank and file type")
	}

	c.SetRGBA(toRGBA(settings.RankAndFile.FontColor))
	err := setFontFace(c, size)
	if err != nil {
		return err
	}
	r.drawRanksAndFiles(c, dx, dy)

	return nil
}

func (r *rendererRankAndFile) drawRanksAndFiles(c *gg.Context, dx, dy float64) {
	rfBoxes := r.getRFBoxes()

	for _, rfBox := range rfBoxes {
		tw, th := c.MeasureString(rfBox.text)
		x := rfBox.box.X + (rfBox.box.Width-tw)/2
		y := rfBox.box.Y + (rfBox.box.Height-th)/2 + th
		c.DrawString(rfBox.text, x+dx, y+dy)
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
		box.Height -= 3 // Adjust the height a little to handle the letter g
		rf = append(rf, RankFile{box: box, text: text})
	}
	return rf
}

func (r *rendererRankAndFile) getRankText(n int) string {
	if settings.Board.Default.Inverted {
		return fmt.Sprintf("%d", 8-n)
	} else {
		return fmt.Sprintf("%d", n+1)
	}
}

func (r *rendererRankAndFile) getFileText(n int) string {
	if settings.Board.Default.Inverted {
		return fmt.Sprintf("%c", 'a'+n)
	} else {
		return fmt.Sprintf("%c", 'h'-n)
	}
}

func (r *rendererRankAndFile) getRankBox(rank int) Rectangle {
	square := float64(settings.Board.Default.Size) / 8
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      0,
		Y:      border + float64(invert(rank))*square,
		Width:  border,
		Height: square,
	}
}

func (r *rendererRankAndFile) getFileBox(file int) Rectangle {
	square := float64(settings.Board.Default.Size) / 8
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      border + float64(invert(file))*square,
		Y:      border + 8*square,
		Width:  square,
		Height: border,
	}
}
