package chessImager

import (
	"strconv"
	"strings"
)

const validChars = "pbnrqkPBNRQK12345678"

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

func validateFen(fen string) bool {
	fens := strings.Split(fen, " ")
	if len(fens) < 6 {
		return false
	}

	// We only care about the first section (board section) of the fen string.
	items := strings.Split(fens[0], "/")
	if len(items) != 8 {
		return false
	}

	for _, item := range items {
		// Check valid characters
		if !checkValidChars(item) {
			return false
		}

		// Check valid length (after normalization)
		if checkLength(item) != 8 {
			return false
		}
	}
	return true
}

func checkValidChars(item string) bool {
	for _, c := range item {
		if strings.Index(validChars, string(c)) == -1 {
			return false
		}
	}

	return true
}

func checkLength(item string) int {
	length := 0

	for _, char := range item {
		if char >= '1' && char <= '8' {
			// If the character is a digit (1-8), add its integer value to the length
			length += int(char - '0')
		} else if char != '/' {
			// If the character is not a digit or '/', increment the length by 1
			length++
		}
	}

	return length
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
