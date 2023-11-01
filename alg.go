package chessImager

import (
	"errors"
	"strings"
)

type alg struct {
	pos string
}

func newAlg(s string) (alg, error) {
	s = strings.ToLower(s)
	if len(s) != 2 {
		return alg{}, errors.New("invalid length of alg")
	}
	if s[0] < 'a' || s[0] > 'h' {
		return alg{}, errors.New("invalid character in alg : " + string(s[0]))
	}
	if s[1] < '1' || s[1] > '8' {
		return alg{}, errors.New("invalid character in alg : " + string(s[1]))
	}
	return alg{s}, nil
}

func (a alg) coords() (int, int) {
	x, y := int(a.pos[0]-'a'), int(a.pos[1]-'1')
	if settings.Board.Default.Inverted {
		return invert(x), invert(y)
	}
	return x, y
}
