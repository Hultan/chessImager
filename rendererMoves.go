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
		err := r.renderMove(c, move)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererMoves) renderMove(c *gg.Context, move Move) error {
	style := r.getStyle(move)

	from, err := newAlg(move.From)
	if err != nil {
		return err
	}
	fromX, fromY := from.coords()

	to, err := newAlg(move.To)
	if err != nil {
		return err
	}
	toX, toY := to.coords()

	dx, dy := toX-fromX, toY-fromY
	x, y := fromX, fromY

	if dx == 0 && dy == 0 {
		return nil // Ignore no move
	}

	c.SetRGBA(style.Color.toRGBA())
	if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
		// Render pawn, rook, bishop, king and queen moves (ie straight moves)
		d := max(abs(dx), abs(dy))
		r.renderDottedLine(c, &x, &y, sgn(dx), sgn(dy), d, style)
	} else {
		// Horse type move (or other weird illegal move)
		if abs(dx) > abs(dy) {
			// abs(dx) > abs(dy) ; vertically first, horizontally second
			r.renderDottedLine(c, &x, &y, 0, sgn(dy), abs(dy), style)
			r.renderDottedLine(c, &x, &y, sgn(dx), 0, abs(dx), style)
		} else {
			// abs(dx) <= abs(dy) ; horizontally first, vertically second
			r.renderDottedLine(c, &x, &y, sgn(dx), 0, abs(dx), style)
			r.renderDottedLine(c, &x, &y, 0, sgn(dy), abs(dy), style)
		}
	}

	return nil
}

func (r *rendererMoves) renderDottedLine(c *gg.Context, x, y *int, dx, dy, moves int, style *MoveStyle) {
	for i := 0; i < moves; i++ {
		r.renderDotInSquare(c, *x, *y, style)
		*x += dx
		*y += dy
	}
}

func (r *rendererMoves) renderDotInSquare(c *gg.Context, x, y int, style *MoveStyle) {
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
