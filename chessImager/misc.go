package chessImager

import (
	"fmt"
	"image/color"
	"strings"
)

var pieceMap = map[string]chessPiece{
	"WK": WhiteKing,
	"WQ": WhiteQueen,
	"WR": WhiteRook,
	"WN": WhiteKnight,
	"WB": WhiteBishop,
	"WP": WhitePawn,
	"BK": BlackKing,
	"BQ": BlackQueen,
	"BR": BlackRook,
	"BN": BlackKnight,
	"BB": BlackBishop,
	"BP": BlackPawn,
}

var embeddedPieces = []PieceRectangle{
	{WhiteKing, Rectangle{0, 0, 333, 333}},
	{WhiteQueen, Rectangle{333, 0, 333, 333}},
	{WhiteBishop, Rectangle{666, 0, 333, 333}},
	{WhiteKnight, Rectangle{999, 0, 333, 333}},
	{WhiteRook, Rectangle{1332, 0, 333, 333}},
	{WhitePawn, Rectangle{1665, 0, 333, 333}},
	{BlackKing, Rectangle{0, 333, 333, 333}},
	{BlackQueen, Rectangle{333, 333, 333, 333}},
	{BlackBishop, Rectangle{666, 333, 333, 333}},
	{BlackKnight, Rectangle{999, 333, 333, 333}},
	{BlackRook, Rectangle{1332, 333, 333, 333}},
	{BlackPawn, Rectangle{1665, 333, 333, 333}},
}

type PieceRectangle struct {
	piece chessPiece
	rect  Rectangle
}

// hexToRGBA converts a hex string (#rrggbbaa) to a color
func hexToRGBA(hex string) (col color.RGBA) {
	// Remove the '#' symbol if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex values for red, green, blue and alpha
	fmt.Sscanf(hex, "%02x%02x%02x%02x", &col.R, &col.G, &col.B, &col.A)

	return col
}

func toRGBA(col ColorRGBA) (float64, float64, float64, float64) {
	return float64(col.R) / 255, float64(col.G) / 255, float64(col.B) / 255, float64(col.A) / 255
}

func invert(x, y int) (int, int) {
	return 7 - x, 7 - y
}

func createPieceRectangleSlice(mapPieces [12]ImageMapPiece) []PieceRectangle {
	result := make([]PieceRectangle, len(mapPieces))
	for _, piece := range mapPieces {
		result = append(result, PieceRectangle{
			piece: pieceMap[piece.Piece],
			rect:  piece.Rect,
		})
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(dx int) int {
	if dx < 0 {
		return -1
	}
	if dx == 0 {
		return 0
	}
	return 1
}
