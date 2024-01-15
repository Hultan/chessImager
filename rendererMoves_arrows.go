package chessImager

import (
	"errors"
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

func (r *rendererMoves) renderArrowMove(c *gg.Context, style *MoveStyle, move Move) error {
	c.SetRGBA(style.Color.toRGBA())

	from, err := newAlg(move.From)
	if err != nil {
		return err
	}

	to, err := newAlg(move.To)
	if err != nil {
		return err
	}

	switch {
	case from.status == moveStatusIllegal:
		return errors.New(fmt.Sprintf("illegal move : %s", from))
	case to.status == moveStatusIllegal:
		return errors.New(fmt.Sprintf("illegal move : %s", to))
	case from.status == moveStatusKingSideCastling && to.status == moveStatusEmpty:
		r.renderCastlingArrow(c, style, whiteKingSideCastling)
	case from.status == moveStatusQueenSideCastling && to.status == moveStatusEmpty:
		r.renderCastlingArrow(c, style, whiteQueenSideCastling)
	case from.status == moveStatusEmpty && to.status == moveStatusKingSideCastling:
		r.renderCastlingArrow(c, style, blackKingSideCastling)
	case from.status == moveStatusEmpty && to.status == moveStatusQueenSideCastling:
		r.renderCastlingArrow(c, style, blackQueenSideCastling)
	case from.status == moveStatusNormal && to.status == moveStatusNormal:
		err = r.renderNormalMoveArrow(c, style, move, from, to)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererMoves) renderNormalMoveArrow(c *gg.Context, style *MoveStyle, move Move, from alg, to alg) error {
	fromX, fromY := from.coords(settings.Board.Default.Inverted)
	toX, toY := to.coords(settings.Board.Default.Inverted)
	dx, dy := toX-fromX, toY-fromY
	if dx == 0 && dy == 0 {
		return nil // Ignore no move
	}

	fx, fy := getSquareBox(fromX, fromY).center()
	rect, err := r.getNextToLast(move)
	if err != nil {
		return err
	}
	styleBox := rect.shrink(style.Factor)
	tx, ty := rect.center()

	if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
		// Render pawn, rook, bishop, king and queen moves (ie straight moves)
		dir := r.getDirection(dx, dy)
		length := math.Sqrt((tx-fx)*(tx-fx) + (ty-fy)*(ty-fy))
		if dir%90 != 0 {
			length += rect.Width/2*math.Sqrt(2) - styleBox.Width
		} else {
			length += rect.Width/2 - styleBox.Width
		}
		r.renderArrow(c, length, styleBox.Width, fx, fy, 0, dir)
	} else {
		// Knight type move (or other weird illegal move)
		dir, rl := r.getKnightDirection(dx, dy)
		if rl == right {
			r.renderKnightArrowRight(c, rect.Width, styleBox.Width, fx, fy, dir)
		} else {
			r.renderKnightArrowLeft(c, rect.Width, styleBox.Width, fx, fy, dir)
		}
	}
	return nil
}

func (r *rendererMoves) renderCastlingArrow(c *gg.Context, style *MoveStyle, castling castlingStatus) {
	var kingPos, rookPos string
	var dir1, dir2 = directionEast, directionWest
	var lengthFactor = 1.5

	square := getSquareBox(0, 0)
	cdy := square.shrink(style.Factor).Width/2 + style.Padding

	switch castling {
	case whiteKingSideCastling:
		kingPos, rookPos = "E1", "H1"
	case whiteQueenSideCastling:
		kingPos, rookPos = "E1", "A1"
		dir1, dir2 = directionWest, directionEast
		lengthFactor = 2.5
		cdy *= -1
	case blackKingSideCastling:
		kingPos, rookPos = "E8", "H8"
	case blackQueenSideCastling:
		kingPos, rookPos = "E8", "A8"
		dir1, dir2 = directionWest, directionEast
		lengthFactor = 2.5
		cdy *= -1
	}

	// Render king castling arrow
	king, _ := newAlg(kingPos)
	styleBox := square.shrink(style.Factor)
	fx, fy := getSquareBox(king.coords(settings.Board.Default.Inverted)).center()
	r.renderArrow(c, square.Width*1.5, styleBox.Width, fx, fy, -cdy, dir1)

	// Render rook castling arrow
	c.SetRGBA(style.Color2.toRGBA())
	rook, _ := newAlg(rookPos)
	fx, fy = getSquareBox(rook.coords(settings.Board.Default.Inverted)).center()
	r.renderArrow(c, square.Width*lengthFactor, styleBox.Width, fx, fy, -cdy, dir2)
}

func (r *rendererMoves) renderArrow(c *gg.Context, length, width, fx, fy, dy float64, dir direction) {
	c.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	c.MoveTo(fx-width/2+dy, fy)
	c.LineTo(fx-width/2+dy, fy-length)
	c.LineTo(fx-width+dy, fy-length)
	c.LineTo(fx+dy, fy-length-width)
	c.LineTo(fx+width+dy, fy-length)
	c.LineTo(fx+width/2+dy, fy-length)
	c.LineTo(fx+width/2+dy, fy)
	c.LineTo(fx-width/2+dy, fy)
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

// Returns next to last square in a move.
// For example if rook moves from d1 to d7, the d6 square will be returned
func (r *rendererMoves) getNextToLast(move Move) (Rectangle, error) {
	// We don't need to check from.status here because it has already
	// been checked in the renderArrowMove() function.
	from, err := newAlg(move.From)
	if err != nil {
		return Rectangle{}, err
	}

	// We don't need to check to.status here because it has already
	// been checked in the renderArrowMove() function.
	to, err := newAlg(move.To)
	if err != nil {
		return Rectangle{}, err
	}

	dx, dy := to.x-from.x, to.y-from.y

	switch {
	case dx == 0 && dy == 0:
		return Rectangle{}, errors.New("no move") // Ignore no move
	case dx == 0 || dy == 0 || abs(dx) == abs(dy): // Straight moves
		return getSquareBox(to.x-sgn(dx), to.y-sgn(dy)), nil
	case abs(dx) == 1 && abs(dy) == 2: // Knight move 1
		return getSquareBox(to.x-sgn(dx), to.y), nil
	case abs(dx) == 2 && abs(dy) == 1: // Knight move 2
		return getSquareBox(to.x, to.y-sgn(dy)), nil
	default:
		panic("illegal move")
	}
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
