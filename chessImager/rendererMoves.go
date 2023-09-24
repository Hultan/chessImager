package chessImager

import "github.com/fogleman/gg"

type rendererMoves struct {
	*Imager
}

func (r *rendererMoves) draw(c *gg.Context) {
	for _, move := range r.settings.Moves {
		fromX, fromY := r.getSquareBox(r.algToCoords(move.From)).Center()
		toX, toY := r.getSquareBox(r.algToCoords(move.To)).Center()
		c.SetRGBA(toRGBA(r.getStyle(move).Color))
		c.SetLineWidth(3)
		c.DrawLine(fromX, fromY, toX, toY)
		c.Stroke()
	}
}

func (r *rendererMoves) getStyle(move Move) *MoveStyle {
	if move.Style == nil {
		return &r.settings.MoveStyle
	} else {
		return move.Style
	}
}
