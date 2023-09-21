package chessImager

import (
	"strconv"
	"strings"
)

const validFenPieces = "pbnrqkPBNRQK"
const validFenEmpty = "12345678"

func validateFen(fen string) bool {
	fens := strings.Split(fen, " ")
	if len(fens) < 6 {
		return false
	}
	items := strings.Split(fens[0], "/")
	if len(items) != 8 {
		return false
	}
	for _, item := range items {
		if len(item) < 1 || len(item) > 8 {
			return false
		}
		if checkLength(item) != 8 {
			return false
		}
	}
	return true
}

func checkLength(item string) int {
	l := 0
	for _, c := range item {
		i := strings.Index(validFenPieces, string(c))
		if i >= 0 {
			l++
			continue
		}
		i = strings.Index(validFenEmpty, string(c))
		if i >= 0 {
			l += i + 1
		}
	}
	return l
}

func normalizeFEN(fen string) string {
	// Fen is already validated to be correct here
	normalized := ""
	fens := strings.Split(fen, " ")
	ranks := strings.Split(fens[0], "/")
	for _, rank := range ranks {
		normalized += "/" + normalizeFENRank(rank)
	}

	return normalized[1:]
}

func normalizeFENRank(fenRank string) string {
	normalized := ""
	for _, symbol := range fenRank {
		skip, err := strconv.Atoi(string(symbol))
		if err == nil {
			normalized += strings.Repeat(" ", skip)
		} else {
			normalized += string(symbol)
		}
	}
	return normalized
}

var letter2Piece = map[rune]chessPiece{
	'p': BlackPawn,
	'b': BlackBishop,
	'n': BlackKnight,
	'r': BlackRook,
	'q': BlackQueen,
	'k': BlackKing,
	'P': WhitePawn,
	'B': WhiteBishop,
	'N': WhiteKnight,
	'R': WhiteRook,
	'Q': WhiteQueen,
	'K': WhiteKing,
	' ': NoPiece,
}
