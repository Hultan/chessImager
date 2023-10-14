package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererMoves struct {
	*Imager
}

func (r *rendererMoves) draw(c *gg.Context) {
	if r.ctx == nil {
		return
	}
	for _, move := range r.ctx.Moves {
		r.renderMove(c, move)
	}
}

func (r *rendererMoves) renderMove(c *gg.Context, move Move) {
	style := r.getStyle(move)
	fromX, fromY := algToCoords(move.From)
	toX, toY := algToCoords(move.To)
	dx, dy := toX-fromX, toY-fromY
	x, y := fromX, fromY

	if dx == 0 && dy == 0 {
		return // Ignore no move
	}

	c.SetRGBA(toRGBA(style.Color))
	if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
		// Rook type move or bishop type move
		d := max(abs(dx), abs(dy))
		for i := 0; i < d; i++ {
			r.highlightBox(c, x, y, style)
			x += sgn(dx)
			y += sgn(dy)
		}
	} else {
		// Horse type move (or other weird illegal move)
		for i := 0; i <= abs(dy); i++ {
			r.highlightBox(c, x, y, style)
			y += sgn(dy)
		}
		y -= sgn(dy)
		for i := 0; i < abs(dx)-1; i++ {
			x += sgn(dx)
			r.highlightBox(c, x, y, style)
		}
	}
}

func (r *rendererMoves) highlightBox(c *gg.Context, x, y int, style *MoveStyle) {
	bb := getSquareBox(x, y).Shrink(style.Factor)
	cX, cY := bb.Center()
	c.DrawCircle(cX, cY, bb.Width/2)
	c.Fill()
}

func (r *rendererMoves) getStyle(move Move) *MoveStyle {
	if move.Style == nil {
		return &settings.MoveStyle
	} else {
		return move.Style
	}
}
