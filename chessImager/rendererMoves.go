package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererMoves struct {
	*Imager
}

func (r *rendererMoves) draw(c *gg.Context) {
	for _, move := range r.settings.Moves {
		factor := r.getStyle(move).Factor / 2
		if factor == 0 {
			factor = 0.15
		}
		fromX, fromY := r.algToCoords(move.From)
		toX, toY := r.algToCoords(move.To)
		dx, dy := toX-fromX, toY-fromY

		if dx == 0 && dy == 0 {
			continue // Ignore no move
		}

		square := float64(r.settings.Board.Default.Size) / 8
		c.SetRGBA(toRGBA(r.getStyle(move).Color))
		if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
			// Rook type move or bishop type move
			d := max(abs(dx), abs(dy))
			x, y := fromX, fromY
			for i := 0; i < d; i++ {
				cX, cY := r.getSquareBox(x, y).Center()
				c.DrawCircle(cX, cY, square*factor)
				c.Fill()
				x += sgn(dx)
				y += sgn(dy)
			}
		} else {
			// Horse type move (or other weird illegal move)
			//x, y := fromX, fromY
			//cX, cY := r.getSquareBox(x, y).Center()
			//c.DrawCircle(cX, cY, square*factor)
			//c.Fill()

			x, y := fromX, fromY
			for i := 0; i <= abs(dy); i++ {
				cX, cY := r.getSquareBox(x, y).Center()
				c.DrawCircle(cX, cY, square*0.15)
				c.Fill()
				y += sgn(dy)
			}
			y++
			for i := 0; i < abs(dx)-1; i++ {
				x += sgn(dx)
				cX, cY := r.getSquareBox(x, y).Center()
				c.DrawCircle(cX, cY, square*0.15)
				c.Fill()
			}
		}
	}
}

func (r *rendererMoves) getStyle(move Move) *MoveStyle {
	if move.Style == nil {
		return &r.settings.MoveStyle
	} else {
		return move.Style
	}
}
