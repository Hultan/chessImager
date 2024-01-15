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
	typ  rankFileType
	box  Rectangle
	text string
}

func (r *rendererRankAndFile) draw(c *gg.Context) error {
	// If Board.Type is BoardTypeImage then we should not draw ranks and files
	if r.settings.Board.Type == boardTypeImage {
		return nil
	}

	// If the user has chosen to not render rank and file, then return
	if r.settings.RankAndFile.Type == rankAndFileTypeNone {
		return nil
	}

	// Don't bother drawing ranks and files when to border is too thin
	border := float64(r.settings.Border.Width)
	if border < borderLimit {
		return nil
	}

	fontSize := r.settings.RankAndFile.FontSize
	c.SetRGBA(r.settings.RankAndFile.FontColor.toRGBA())
	err := r.setFontFace(r.settings.FontStyle.Path, c, fontSize)
	if err != nil {
		return err
	}

	switch r.settings.RankAndFile.Type {
	case rankAndFileTypeInBorder:
		r.drawRanksAndFiles(c, 0, 0)
	case rankAndFileTypeInSquares:
		const padding = 3
		square := float64(r.settings.getBoardBox().Width) / 8
		diff := (square - float64(fontSize) - padding) / 2
		r.drawRanksAndFiles(c, diff, diff)
	default:
		return errors.New("invalid rank and file type")
	}

	return nil
}

func (r *rendererRankAndFile) drawRanksAndFiles(c *gg.Context, dx, dy float64) {
	rfBoxes := r.getRFBoxes()

	var diff float64
	if r.useInternalFont {
		diff -= 2
	}

	for _, rfBox := range rfBoxes {
		tw, th := c.MeasureString(rfBox.text)
		x := rfBox.box.X + (rfBox.box.Width-tw)/2
		// We are adjusting by 2 pixels here because of bug in MeasureString?
		y := rfBox.box.Y + (rfBox.box.Height-th)/2 + th + diff
		if rfBox.typ == rank {
			c.DrawString(rfBox.text, x-dx, y-dy)
		} else {
			c.DrawString(rfBox.text, x+dx, y+dy)
		}
	}
}

func (r *rendererRankAndFile) getRFBoxes() []RankFile {
	var rf []RankFile
	var box Rectangle

	for i := 0; i < 8; i++ {
		// Ranks
		text := r.getRankText(i)
		if r.settings.RankAndFile.Type == rankAndFileTypeInBorder {
			box = r.getRankBox(i)
		} else {
			box = r.settings.getSquareBox(0, i)
		}
		rf = append(rf, RankFile{box: box, text: text, typ: rank})

		// Files
		text = r.getFileText(i)
		if r.settings.RankAndFile.Type == rankAndFileTypeInBorder {
			box = r.getFileBox(i)
		} else {
			box = r.settings.getSquareBox(i, 0)
		}
		rf = append(rf, RankFile{box: box, text: text, typ: file})
	}

	return rf
}

func (r *rendererRankAndFile) getRankText(n int) string {
	if r.settings.Board.Default.Inverted {
		return fmt.Sprintf("%d", 8-n)
	} else {
		return fmt.Sprintf("%d", n+1)
	}
}

func (r *rendererRankAndFile) getFileText(n int) string {
	if r.settings.Board.Default.Inverted {
		return fmt.Sprintf("%c", 'H'-n)
	} else {
		return fmt.Sprintf("%c", 'A'+n)
	}
}

func (r *rendererRankAndFile) getRankBox(rank int) Rectangle {
	square := float64(r.settings.getBoardBox().Width) / 8
	border := float64(r.settings.Border.Width)

	return Rectangle{
		X:      0,
		Y:      border + float64(invert(rank))*square,
		Width:  border,
		Height: square,
	}
}

func (r *rendererRankAndFile) getFileBox(file int) Rectangle {
	square := float64(r.settings.getBoardBox().Width) / 8
	border := float64(r.settings.Border.Width)

	return Rectangle{
		X:      border + float64(file)*square,
		Y:      border + 8*square,
		Width:  square,
		Height: border,
	}
}
