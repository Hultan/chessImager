package chessImager

import (
	"errors"
	"strings"
)

type alg struct {
	pos string
}

func newAlg(pos string) (alg, error) {
	pos = strings.ToLower(pos)
	if len(pos) != 2 {
		return alg{}, errors.New("invalid length of alg")
	}
	if pos[0] < 'a' || pos[0] > 'h' {
		return alg{}, errors.New("invalid character in alg : " + string(pos[0]))
	}
	if pos[1] < '1' || pos[1] > '8' {
		return alg{}, errors.New("invalid character in alg : " + string(pos[1]))
	}
	return alg{pos}, nil
}

func (a alg) coords() (int, int) {
	x, y := int(a.pos[0]-'a'), int(a.pos[1]-'1')
	if settings.Board.Default.Inverted {
		return invert(x), invert(y)
	}
	return x, y
}
