package chessImager

import (
	"errors"

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

	switch style.Type {
	case MoveTypeDots:
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
		// TODO : How to handle castling?
	case MoveTypeArrow:
		var rect Rectangle
		fx, fy := getSquareBox(fromX, fromY).Center()
		rect, err = r.getNextToLast(move)
		if err != nil {
			return err
		}
		tx, ty := rect.Center()
		box := rect.Shrink(style.Factor)

		c.SetLineWidth(box.Width)
		c.DrawLine(fx, fy, tx, ty)
		c.Stroke()

		r.renderArrowHead(c, rect, rect.Width, box.Width, dx, dy)
	default:
		return errors.New("illegal move type")
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

// Render arrow head for all types of moves
func (r *rendererMoves) renderArrowHead(c *gg.Context, rect Rectangle, square, width float64, dx, dy int) {
	switch {
	case dx == 0 && dy > 0:
		// Vertical move up (queen, rook, king, pawn)
		r.renderArrowHeadInDirection(c, rect, square, width, directionNorth)
		return
	case dx == 0 && dy < 0:
		// Vertical move down (queen, rook, king)
		r.renderArrowHeadInDirection(c, rect, square, width, directionSouth)
		return
	case dx < 0 && dy == 0:
		// Horizontal move left (queen, rook, king)
		r.renderArrowHeadInDirection(c, rect, square, width, directionWest)
		return
	case dx > 0 && dy == 0:
		// Horizontal move right (queen, rook, king)
		r.renderArrowHeadInDirection(c, rect, square, width, directionEast)
		return
	case abs(dx) == abs(dy):
		// Bishop move, king (diagonal), queen (diagonal) or pawn (when capturing)
		switch {
		case dx > 0 && dy > 0:
			// NE
			r.renderArrowHeadInDirection(c, rect, square, width, directionNorthEast)
		case dx > 0 && dy < 0:
			// SE
			r.renderArrowHeadInDirection(c, rect, square, width, directionSouthEast)
		case dx < 0 && dy > 0:
			// NW
			r.renderArrowHeadInDirection(c, rect, square, width, directionNorthWest)
		case dx < 0 && dy < 0:
			// SW
			r.renderArrowHeadInDirection(c, rect, square, width, directionSouthWest)
		}
		return
	case false:
		// Castling, how to handle

	// Knight moves
	case (dx == 2 && dy == 1) || (dx == -2 && dy == 1):
		r.renderArrowHeadInDirection(c, rect, square, width, directionNorth)
	case (dx == 1 && dy == 2) || (dx == 1 && dy == -2):
		r.renderArrowHeadInDirection(c, rect, square, width, directionEast)
	case (dx == 2 && dy == -1) || (dx == -2 && dy == -1):
		r.renderArrowHeadInDirection(c, rect, square, width, directionSouth)
	case (dx == -1 && dy == 2) || (dx == -1 && dy == -2):
		r.renderArrowHeadInDirection(c, rect, square, width, directionWest)

	default:
		// Illegal moves
		return
	}
}

func (r *rendererMoves) renderArrowHeadInDirection(c *gg.Context, rect Rectangle, square, width float64, dir direction) {
	cx, cy := rect.Center()

	// Rotate to draw in correct angle
	c.RotateAbout(gg.Radians(float64(dir)), cx, cy)

	// Draw line
	c.SetLineWidth(abs(width))
	c.DrawLine(cx, cy, cx, cy-square/2+width)
	c.Stroke()

	// Draw arrow head (triangle part)
	c.SetLineWidth(1)
	c.MoveTo(cx+width, cy-square/2+width)
	c.LineTo(cx-width, cy-square/2+width)
	c.LineTo(cx, cy-square/2)
	c.ClosePath()
	c.Fill()

	// Rotate back
	c.RotateAbout(gg.Radians(float64(-dir)), cx, cy)
}

// Returns next to last square in a move.
// For example if rook moves from d1 to d7, the d6 square will be returned
func (r *rendererMoves) getNextToLast(move Move) (Rectangle, error) {
	from, err := newAlg(move.From)
	if err != nil {
		return Rectangle{}, err
	}
	fromX, fromY := from.coords()

	to, err := newAlg(move.To)
	if err != nil {
		return Rectangle{}, err
	}
	toX, toY := to.coords()

	dx, dy := toX-fromX, toY-fromY

	if dx == 0 && dy == 0 {
		return Rectangle{}, errors.New("no move") // Ignore no move
	} else if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
		// Straight moves
		return getSquareBox(toX-sgn(dx), toY-sgn(dy)), nil
	} else if abs(dx) == 1 && abs(dy) == 2 {
		// Knight move 1
		return getSquareBox(toX-sgn(dx), toY), nil
	} else if abs(dx) == 2 && abs(dy) == 1 {
		// Knight move 2
		return getSquareBox(toX, toY-sgn(dy)), nil
	}

	return Rectangle{}, errors.New("invalid move")
}
