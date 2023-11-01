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
	// If Board.Type is BoardTypeImage then we should not draw ranks and files
	if settings.Board.Type == BoardTypeImage {
		return nil
	}

	// If the user has chosen to not render rank and file, then return
	if settings.RankAndFile.Type == RankAndFileTypeNone {
		return nil
	}

	// Don't bother drawing ranks and files when to border is too thin
	border := float64(settings.Border.Width)
	if border < borderLimit {
		return nil
	}

	fontSize := settings.RankAndFile.FontSize
	c.SetRGBA(settings.RankAndFile.FontColor.toRGBA())
	err := setFontFace(c, fontSize)
	if err != nil {
		return err
	}

	switch settings.RankAndFile.Type {
	case RankAndFileTypeInBorder:
		r.drawRanksAndFiles(c, 0, 0)
	case RankAndFileTypeInSquares:
		// TODO : Needs better implementation
		square := float64(getBoardBox().Width) / 8
		dx, dy := (square-border)/2-5, -border-5
		r.drawRanksAndFiles(c, dx, dy)
	default:
		return errors.New("invalid rank and file type")
	}

	return nil
}

func (r *rendererRankAndFile) drawRanksAndFiles(c *gg.Context, dx, dy float64) {
	rfBoxes := r.getRFBoxes()

	for _, rfBox := range rfBoxes {
		tw, th := c.MeasureString(rfBox.text)
		x := rfBox.box.X + (rfBox.box.Width-tw)/2
		// We are adjusting by 2 pixels here because of bug in MeasureString?
		y := rfBox.box.Y + (rfBox.box.Height-th)/2 + th - 2
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
		return fmt.Sprintf("%c", 'H'-n)
	} else {
		return fmt.Sprintf("%c", 'A'+n)
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
		X:      border + float64(file)*square,
		Y:      border + 8*square,
		Width:  square,
		Height: border,
	}
}
