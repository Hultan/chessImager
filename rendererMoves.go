package chessImager

import (
	"errors"
	"math"

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

	fromX, fromY, toX, toY, err := r.getFromAndTo(move)
	if err != nil {
		return err
	}

	dx, dy := toX-fromX, toY-fromY
	if dx == 0 && dy == 0 {
		return nil // Ignore no move
	}

	c.SetRGBA(style.Color.toRGBA())

	switch style.Type {
	case MoveTypeDots:
		x, y := fromX, fromY
		if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
			// Render pawn, rook, bishop, king and queen moves (ie straight moves)
			d := max(abs(dx), abs(dy))
			r.renderDottedLine(c, &x, &y, sgn(dx), sgn(dy), d, style)
		} else {
			// Knight type move (or other weird illegal move)
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
	case MoveTypeArrow:
		fx, fy := getSquareBox(fromX, fromY).center()
		rect, err := r.getNextToLast(move)
		if err != nil {
			return err
		}
		tx, ty := rect.center()
		styleBox := rect.shrink(style.Factor)

		if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
			// Render pawn, rook, bishop, king and queen moves (ie straight moves)
			const margin = 5
			dir := r.getDirection(dx, dy)
			length := math.Sqrt((tx-fx)*(tx-fx) + (ty-fy)*(ty-fy))
			if dir%90 != 0 {
				length += rect.Width*2/3*math.Sqrt(2) - styleBox.Width - margin
			} else {
				length += rect.Width/2 - styleBox.Width - margin
			}
			r.renderArrow(c, length, styleBox.Width, fx, fy, dir)
		} else {
			// Knight type move (or other weird illegal move)
			dir, rl := r.getKnightDirection(dx, dy)
			if rl == right {
				r.renderKnightArrowRight(c, rect.Width, styleBox.Width, fx, fy, dir)
			} else {
				r.renderKnightArrowLeft(c, rect.Width, styleBox.Width, fx, fy, dir)
			}
		}
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
	bb := getSquareBox(x, y).shrink(style.Factor)
	cX, cY := bb.center()
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

// Returns next to last square in a move.
// For example if rook moves from d1 to d7, the d6 square will be returned
func (r *rendererMoves) getNextToLast(move Move) (Rectangle, error) {
	fromX, fromY, toX, toY, err := r.getFromAndTo(move)
	if err != nil {
		return Rectangle{}, err
	}

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

func (r *rendererMoves) getFromAndTo(move Move) (int, int, int, int, error) {
	from, err := newAlg(move.From)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	fromX, fromY := from.coords()

	to, err := newAlg(move.To)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	toX, toY := to.coords()

	return fromX, fromY, toX, toY, nil
}

func (r *rendererMoves) getDirection(dx int, dy int) direction {
	switch {
	case dx == 0 && dy < 0:
		return directionSouth
	case dx == 0 && dy > 0:
		return directionNorth
	case dy == 0 && dx < 0:
		return directionWest
	case dy == 0 && dx > 0:
		return directionEast
	case dx > 0 && dy > 0:
		return directionNorthEast
	case dx > 0 && dy < 0:
		return directionSouthEast
	case dx < 0 && dy > 0:
		return directionNorthWest
	case dx < 0 && dy < 0:
		return directionSouthWest
	default:
		panic("invalid direction")
	}
}

func (r *rendererMoves) renderArrow(c *gg.Context, length, width, fx, fy float64, dir direction) {
	c.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	c.MoveTo(fx-width/2, fy)
	c.LineTo(fx-width/2, fy-length)
	c.LineTo(fx-width, fy-length)
	c.LineTo(fx, fy-length-width)
	c.LineTo(fx+width, fy-length)
	c.LineTo(fx+width/2, fy-length)
	c.LineTo(fx+width/2, fy)
	c.LineTo(fx-width/2, fy)
	c.Fill()
	c.RotateAbout(gg.Radians(float64(-dir)), fx, fy)
}

func (r *rendererMoves) renderKnightArrowRight(c *gg.Context, square, width, fx, fy float64, dir direction) {
	length := square * 2

	c.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	c.MoveTo(fx-width/2, fy)
	c.LineTo(fx-width/2, fy-length-width/2)
	c.LineTo(fx+square/2-width, fy-length-width/2)
	c.LineTo(fx+square/2-width, fy-length-width)
	c.LineTo(fx+square/2, fy-length)
	c.LineTo(fx+square/2-width, fy-length+width)
	c.LineTo(fx+square/2-width, fy-length+width/2)
	c.LineTo(fx+width/2, fy-length+width/2)
	c.LineTo(fx+width/2, fy)
	c.LineTo(fx-width/2, fy)
	c.Fill()
	c.RotateAbout(gg.Radians(float64(-dir)), fx, fy)
}

func (r *rendererMoves) renderKnightArrowLeft(c *gg.Context, square, width, fx, fy float64, dir direction) {
	length := square * 2

	c.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	c.MoveTo(fx-width/2, fy)
	c.LineTo(fx-width/2, fy-length+width/2)
	c.LineTo(fx-square/2+width, fy-length+width/2)
	c.LineTo(fx-square/2+width, fy-length+width)
	c.LineTo(fx-square/2, fy-length)
	c.LineTo(fx-square/2+width, fy-length-width)
	c.LineTo(fx-square/2+width, fy-length-width/2)
	c.LineTo(fx+width/2, fy-length-width/2)
	c.LineTo(fx+width/2, fy)
	c.LineTo(fx-width/2, fy)
	c.Fill()
	c.RotateAbout(gg.Radians(float64(-dir)), fx, fy)
}

func (r *rendererMoves) getKnightDirection(dx int, dy int) (direction, leftRight) {
	switch {
	case dx == 1 && dy == 2:
		return directionNorth, right
	case dx == -1 && dy == 2:
		return directionNorth, left
	case dx == 1 && dy == -2:
		return directionSouth, left
	case dx == -1 && dy == -2:
		return directionSouth, right
	case dy == 1 && dx == 2:
		return directionEast, left
	case dy == -1 && dx == 2:
		return directionEast, right
	case dy == 1 && dx == -2:
		return directionWest, right
	case dy == -1 && dx == -2:
		return directionWest, left
	default:
		panic("invalid move")
	}
}
