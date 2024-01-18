package chessImager

import (
	"errors"
	"fmt"

	"github.com/fogleman/gg"
)

func (r *rendererMoves) renderDottedMove(c *gg.Context, style *MoveStyle, move Move) error {
	c.SetRGBA(style.Color.toRGBA())

	from, err := newAlg(move.From, r.settings.Board.Default.Inverted)
	if err != nil {
		return err
	}

	to, err := newAlg(move.To, r.settings.Board.Default.Inverted)
	if err != nil {
		return err
	}

	switch {
	case from.status == moveStatusIllegal:
		return errors.New(fmt.Sprintf("illegal move : %s", from))
	case to.status == moveStatusIllegal:
		return errors.New(fmt.Sprintf("illegal move : %s", to))
	case from.status == moveStatusKingSideCastling && to.status == moveStatusEmpty:
		r.renderCastlingDottedLine(c, whiteKingSideCastling, style)
	case from.status == moveStatusQueenSideCastling && to.status == moveStatusEmpty:
		r.renderCastlingDottedLine(c, whiteQueenSideCastling, style)
	case from.status == moveStatusEmpty && to.status == moveStatusKingSideCastling:
		r.renderCastlingDottedLine(c, blackKingSideCastling, style)
	case from.status == moveStatusEmpty && to.status == moveStatusQueenSideCastling:
		r.renderCastlingDottedLine(c, blackQueenSideCastling, style)
	case from.status == moveStatusNormal && to.status == moveStatusNormal:
		err = r.renderNormalDottedMove(c, style, from, to)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererMoves) renderNormalDottedMove(c *gg.Context, style *MoveStyle, from, to alg) error {
	fromX, fromY := from.coords()
	toX, toY := to.coords()

	dx, dy := toX-fromX, toY-fromY
	if dx == 0 && dy == 0 {
		return nil // Ignore no move
	}

	x, y := fromX, fromY
	if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
		// Render pawn, rook, bishop, king and queen moves (ie straight moves)
		d := max(abs(dx), abs(dy))
		r.renderDottedLine(c, &x, &y, sgn(dx), sgn(dy), d, 0, style)
	} else {
		// Knight type move (or other weird illegal move)
		if abs(dx) > abs(dy) {
			// abs(dx) > abs(dy) ; vertically first, horizontally second
			r.renderDottedLine(c, &x, &y, 0, sgn(dy), abs(dy), 0, style)
			r.renderDottedLine(c, &x, &y, sgn(dx), 0, abs(dx), 0, style)
		} else {
			// abs(dx) <= abs(dy) ; horizontally first, vertically second
			r.renderDottedLine(c, &x, &y, sgn(dx), 0, abs(dx), 0, style)
			r.renderDottedLine(c, &x, &y, 0, sgn(dy), abs(dy), 0, style)
		}
	}
	return nil
}

func (r *rendererMoves) renderCastlingDottedLine(c *gg.Context, castling castlingStatus, style *MoveStyle) {
	var kx, ky, rx, ry int
	var dx, dy, rookMoves = 1, 0, 3

	switch castling {
	case whiteKingSideCastling:
		kx, ky = 4, 0 // E1
		rx, ry = 7, 0 // H1
	case whiteQueenSideCastling:
		kx, ky = 4, 0 // E1
		rx, ry = 0, 0 // A1
		dx = -1
		rookMoves = 4
	case blackKingSideCastling:
		kx, ky = 4, 7 // E8
		rx, ry = 7, 7 // H8
	case blackQueenSideCastling:
		kx, ky = 4, 7 // E8
		rx, ry = 0, 7 // A8
		dx = -1
		rookMoves = 4
	}
	// Calculate the castling dy
	cdy := r.getSquareBox(0, 0).shrink(style.Factor).Width/2 + style.Padding

	c.SetRGBA(style.Color.toRGBA())
	r.renderDottedLine(c, &kx, &ky, dx, dy, 3, -cdy, style)

	c.SetRGBA(style.Color2.toRGBA())
	r.renderDottedLine(c, &rx, &ry, dx*-1, dy, rookMoves, cdy, style)
}

func (r *rendererMoves) renderDottedLine(c *gg.Context, x, y *int, dx, dy, moves int, cdy float64, style *MoveStyle) {
	for i := 0; i < moves; i++ {
		r.renderDotInSquare(c, *x, *y, cdy, style)
		*x += dx
		*y += dy
	}
}

func (r *rendererMoves) renderDotInSquare(c *gg.Context, x, y int, cdy float64, style *MoveStyle) {
	bb := r.getSquareBox(x, y).shrink(style.Factor)
	cX, cY := bb.center()
	c.DrawCircle(cX, cY+cdy, bb.Width/2)
	c.Fill()
}
