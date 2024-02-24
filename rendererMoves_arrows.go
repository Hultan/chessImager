package chessImager

import (
	"errors"
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

func (r *rendererMoves) renderArrowMove(style *MoveStyle, move Move) error {
	r.gg.SetRGBA(style.Color.toRGBA())

	from, err := newAlg(move.From, r.inverted)
	if err != nil {
		return err
	}

	to, err := newAlg(move.To, r.inverted)
	if err != nil {
		return err
	}

	switch {
	case from.status == moveStatusIllegal:
		return errors.New(fmt.Sprintf("illegal from move : %s", from))
	case to.status == moveStatusIllegal:
		return errors.New(fmt.Sprintf("illegal to move : %s", to))
	case from.status == moveStatusKingSideCastling && to.status == moveStatusEmpty:
		r.renderCastlingArrow(style, whiteKingSideCastling)
	case from.status == moveStatusQueenSideCastling && to.status == moveStatusEmpty:
		r.renderCastlingArrow(style, whiteQueenSideCastling)
	case from.status == moveStatusEmpty && to.status == moveStatusKingSideCastling:
		r.renderCastlingArrow(style, blackKingSideCastling)
	case from.status == moveStatusEmpty && to.status == moveStatusQueenSideCastling:
		r.renderCastlingArrow(style, blackQueenSideCastling)
	case from.status == moveStatusNormal && to.status == moveStatusNormal:
		err = r.renderNormalMoveArrow(style, move, from, to)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererMoves) renderNormalMoveArrow(style *MoveStyle, move Move, from alg, to alg) error {
	fromX, fromY := from.coords()
	toX, toY := to.coords()
	dx, dy := toX-fromX, toY-fromY
	if dx == 0 && dy == 0 {
		return nil // Ignore no move
	}

	fx, fy := r.getSquareBox(fromX, fromY).center()
	rect, err := r.getNextToLast(move)
	if err != nil {
		return err
	}
	styleBox := rect.shrink(style.Factor)
	tx, ty := rect.center()

	if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
		// Render pawn, rook, bishop, king and queen moves (ie straight moves)
		dir := r.getStraightMoveDirection(dx, dy)
		length := math.Sqrt((tx-fx)*(tx-fx) + (ty-fy)*(ty-fy))
		factor := 1.0
		if dir%90 != 0 {
			factor = math.Sqrt(2)
		}
		length += rect.Width/2*factor - styleBox.Width
		r.renderArrow(length, styleBox.Width, fx, fy, 0, dir)
	} else {
		// Knight type move (or other weird illegal move)
		dir, rl := r.getKnightDirection(dx, dy)
		if rl == right {
			r.renderKnightArrowRight(rect.Width, styleBox.Width, fx, fy, dir)
		} else {
			r.renderKnightArrowLeft(rect.Width, styleBox.Width, fx, fy, dir)
		}
	}
	return nil
}

func (r *rendererMoves) renderCastlingArrow(style *MoveStyle, castling castlingStatus) {
	var kingPos, rookPos string
	var dir1, dir2 = directionEast, directionWest
	var lengthFactor = 1.5

	square := r.getSquareBox(0, 0)
	cdy := style.Padding

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
	king, _ := newAlg(kingPos, r.inverted)
	styleBox := square.shrink(style.Factor)
	fx, fy := r.getSquareBox(king.coords()).center()
	r.renderArrow(square.Width*1.5, styleBox.Width, fx, fy, -cdy, dir1)

	// Render rook castling arrow
	r.gg.SetRGBA(style.Color2.toRGBA())
	rook, _ := newAlg(rookPos, r.inverted)
	fx, fy = r.getSquareBox(rook.coords()).center()
	r.renderArrow(square.Width*lengthFactor, styleBox.Width, fx, fy, -cdy, dir2)
}

func (r *rendererMoves) renderArrow(length, width, fx, fy, dy float64, dir direction) {
	r.gg.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	r.gg.MoveTo(fx-width/2+dy, fy)
	r.gg.LineTo(fx-width/2+dy, fy-length)
	r.gg.LineTo(fx-width+dy, fy-length)
	r.gg.LineTo(fx+dy, fy-length-width)
	r.gg.LineTo(fx+width+dy, fy-length)
	r.gg.LineTo(fx+width/2+dy, fy-length)
	r.gg.LineTo(fx+width/2+dy, fy)
	r.gg.LineTo(fx-width/2+dy, fy)
	r.gg.Fill()
	r.gg.RotateAbout(gg.Radians(float64(-dir)), fx, fy)
}

func (r *rendererMoves) renderKnightArrowRight(square, width, fx, fy float64, dir direction) {
	length := square * 2

	r.gg.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	r.gg.MoveTo(fx-width/2, fy)
	r.gg.LineTo(fx-width/2, fy-length-width/2)
	r.gg.LineTo(fx+square/2-width, fy-length-width/2)
	r.gg.LineTo(fx+square/2-width, fy-length-width)
	r.gg.LineTo(fx+square/2, fy-length)
	r.gg.LineTo(fx+square/2-width, fy-length+width)
	r.gg.LineTo(fx+square/2-width, fy-length+width/2)
	r.gg.LineTo(fx+width/2, fy-length+width/2)
	r.gg.LineTo(fx+width/2, fy)
	r.gg.LineTo(fx-width/2, fy)
	r.gg.Fill()
	r.gg.RotateAbout(gg.Radians(float64(-dir)), fx, fy)
}

func (r *rendererMoves) renderKnightArrowLeft(square, width, fx, fy float64, dir direction) {
	length := square * 2

	r.gg.RotateAbout(gg.Radians(float64(dir)), fx, fy)
	r.gg.MoveTo(fx-width/2, fy)
	r.gg.LineTo(fx-width/2, fy-length+width/2)
	r.gg.LineTo(fx-square/2+width, fy-length+width/2)
	r.gg.LineTo(fx-square/2+width, fy-length+width)
	r.gg.LineTo(fx-square/2, fy-length)
	r.gg.LineTo(fx-square/2+width, fy-length-width)
	r.gg.LineTo(fx-square/2+width, fy-length-width/2)
	r.gg.LineTo(fx+width/2, fy-length-width/2)
	r.gg.LineTo(fx+width/2, fy)
	r.gg.LineTo(fx-width/2, fy)
	r.gg.Fill()
	r.gg.RotateAbout(gg.Radians(float64(-dir)), fx, fy)
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
	from, err := newAlg(move.From, r.inverted)
	if err != nil {
		return Rectangle{}, err
	}

	// We don't need to check to.status here because it has already
	// been checked in the renderArrowMove() function.
	to, err := newAlg(move.To, r.inverted)
	if err != nil {
		return Rectangle{}, err
	}

	dx, dy := to.x-from.x, to.y-from.y

	switch {
	case dx == 0 && dy == 0:
		return Rectangle{}, errors.New("no move") // Ignore no move
	case dx == 0 || dy == 0 || abs(dx) == abs(dy): // Straight moves
		return r.getSquareBox(to.x-sgn(dx), to.y-sgn(dy)), nil
	case abs(dx) == 1 && abs(dy) == 2: // Knight move 1
		return r.getSquareBox(to.x-sgn(dx), to.y), nil
	case abs(dx) == 2 && abs(dy) == 1: // Knight move 2
		return r.getSquareBox(to.x, to.y-sgn(dy)), nil
	default:
		panic("illegal move")
	}
}

func (r *rendererMoves) getStraightMoveDirection(dx int, dy int) direction {
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
