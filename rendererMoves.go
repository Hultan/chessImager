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
		// Render pawn, rook, bishop, king and queen moves
		r.renderStraightMoves(c, dx, dy, x, y, style)
	} else {
		// Render knight moves (and illegal moves)
		r.renderOtherMoves(c, dx, dy, x, y, style)
	}
}

func (r *rendererMoves) renderStraightMoves(c *gg.Context, dx int, dy int, x int, y int, style *MoveStyle) {
	switch style.Type {
	case MoveTypeDots:
		// Rook type move or bishop type move
		d := max(abs(dx), abs(dy))
		r.renderDottedMove(c, &x, &y, style, d, sgn(dx), sgn(dy))
	case MoveTypeArrow:
		// Not implemented yet
	}
}

func (r *rendererMoves) renderOtherMoves(c *gg.Context, dx int, dy int, x int, y int, style *MoveStyle) {
	switch style.Type {
	case MoveTypeDots:
		r.renderOtherMovesDotted(c, dx, dy, x, y, style)
	case MoveTypeArrow:
		r.renderOtherMovesArrow(c, dx, dy, x, y, style)
	}
}

//
// Arrow move
//

func (r *rendererMoves) renderOtherMovesArrow(c *gg.Context, dx int, dy int, x int, y int, style *MoveStyle) {

}

//
// Dotted move
//

func (r *rendererMoves) renderOtherMovesDotted(c *gg.Context, dx int, dy int, x int, y int, style *MoveStyle) {
	// Horse type move (or other weird illegal move)
	dir := r.getPreferredDirection(dx, dy)
	if dir {
		// abs(dx) > abs(dy) ; vertically first, horizontally second
		r.renderDottedMove(c, &x, &y, style, abs(dy), 0, sgn(dy))
		r.renderDottedMove(c, &x, &y, style, abs(dx), sgn(dx), 0)
	} else {
		// abs(dx) <= abs(dy) ; horizontally first, vertically second
		r.renderDottedMove(c, &x, &y, style, abs(dx), sgn(dx), 0)
		r.renderDottedMove(c, &x, &y, style, abs(dy), 0, sgn(dy))
	}
}

func (r *rendererMoves) renderDottedMove(c *gg.Context, x, y *int, style *MoveStyle, moves, dx, dy int) {
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

//
// Misc
//

func (r *rendererMoves) getStyle(move Move) *MoveStyle {
	if move.Style == nil {
		return &settings.MoveStyle
	} else {
		return move.Style
	}
}
