package chessImager

import (
	"errors"
	"fmt"
	"strings"
)

type alg struct {
	pos string

	x, y     int
	status   moveStatus
	inverted bool
}

var algs = map[string]alg{
	"":      {pos: "", status: moveStatusEmpty},
	"0-0":   {pos: "0-0", status: moveStatusKingSideCastling},
	"o-o":   {pos: "o-o", status: moveStatusKingSideCastling},
	"0-0-0": {pos: "0-0-0", status: moveStatusQueenSideCastling},
	"o-o-o": {pos: "o-o-o", status: moveStatusQueenSideCastling},
}

// newAlg calculates coordinates (0-7),(0-7) from a chess position string, like "C5".
// It also handles special cases, like castling and empty strings.
func newAlg(s string, inverted bool) (alg, error) {
	s = strings.ToLower(s)

	fixedAlg, ok := algs[s]
	if ok {
		fixedAlg.pos = s
		return fixedAlg, nil
	}

	// Check illegal moves
	a := alg{pos: s, status: moveStatusIllegal, inverted: inverted}
	if len(s) != 2 {
		return a, errors.New("invalid length of alg")
	} else if s[0] < 'a' || s[0] > 'h' {
		return a, errors.New("invalid character in alg : " + string(s[0]))
	} else if s[1] < '1' || s[1] > '8' {
		return a, errors.New("invalid character in alg : " + string(s[1]))
	}

	// Normal moves
	a.status = moveStatusNormal
	a.x = int(s[0] - 'a')
	a.y = int(s[1] - '1')

	return a, nil
}

func (a alg) coords() (int, int) {
	if a.status != moveStatusNormal {
		// ok to panic here, it's an internal struct, and it is
		// being used wrong!
		panic("not a normal move, check status field")
	}

	if a.inverted {
		return invert(a.x), invert(a.y)
	}
	return a.x, a.y
}

func (a alg) String() string {
	return fmt.Sprintf("move: %s", a.pos)
}
