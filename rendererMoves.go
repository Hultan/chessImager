package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererMoves struct {
	*Imager
}

func (r *rendererMoves) draw(c *gg.Context) error {
	if r.ctx == nil {
		return nil
	}
	for _, move := range r.ctx.Moves {
		r.renderMove(c, move)
	}

	return nil
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
		// Render pawn, rook, bishop, king and queen moves
		// Rook type move or bishop type move
		d := max(abs(dx), abs(dy))
		r.renderDottedLine(c, &x, &y, sgn(dx), sgn(dy), d, style)
	} else {
		// Horse type move (or other weird illegal move)
		dir := r.getPreferredDirection(dx, dy)
		if dir {
			// abs(dx) > abs(dy) ; vertically first, horizontally second
			r.renderDottedLine(c, &x, &y, 0, sgn(dy), abs(dy), style)
			r.renderDottedLine(c, &x, &y, sgn(dx), 0, abs(dx), style)
		} else {
			// abs(dx) <= abs(dy) ; horizontally first, vertically second
			r.renderDottedLine(c, &x, &y, sgn(dx), 0, abs(dx), style)
			r.renderDottedLine(c, &x, &y, 0, sgn(dy), abs(dy), style)
		}
	}
}

func (r *rendererMoves) renderDottedLine(c *gg.Context, x, y *int, dx, dy, moves int, style *MoveStyle) {
	for i := 0; i < moves; i++ {
		r.renderDot(c, *x, *y, style)
		*x += dx
		*y += dy
	}
}

func (r *rendererMoves) renderDot(c *gg.Context, x, y int, style *MoveStyle) {
	bb := getSquareBox(x, y).Shrink(style.Factor)
	cX, cY := bb.Center()
	c.DrawCircle(cX, cY, bb.Width/2)
	c.Fill()
}

func (r *rendererMoves) getPreferredDirection(dx, dy int) bool {
	if abs(dx) <= abs(dy) {
		return false
	}

	return true
}

func (r *rendererMoves) getStyle(move Move) *MoveStyle {
	if move.Style == nil {
		return &settings.MoveStyle
	} else {
		return move.Style
	}
}
