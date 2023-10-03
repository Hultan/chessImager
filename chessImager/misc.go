package chessImager

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
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
func hexToRGBA(hex string) (col color.RGBA, err error) {
	// Remove the '#' symbol if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex values for red, green, blue and alpha
	_, err = fmt.Sscanf(hex, "%02x%02x%02x%02x", &col.R, &col.G, &col.B, &col.A)
	if err != nil {
		return col, fmt.Errorf("invalid color (%s) : %v", hex, err)
	}

	return col, nil
}

func toRGBA(col ColorRGBA) (float64, float64, float64, float64) {
	return float64(col.R) / 255, float64(col.G) / 255, float64(col.B) / 255, float64(col.A) / 255
}

func invert(x int) int {
	return 7 - x
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

// getBoardSize returns a rectangle with the size of the board
// plus the border surrounding it.
func getBoardSize() image.Rectangle {
	size := settings.Board.Default.Size + settings.Border.Width*2

	return image.Rectangle{
		Max: image.Point{
			X: size,
			Y: size,
		},
	}
}

func algToCoords(alg string) (int, int) {
	alg = strings.ToLower(alg)
	if len(alg) != 2 {
		panic("invalid length of alg")
	}
	if alg[0] < 'a' || alg[0] > 'h' {
		panic("invalid character in alg : " + string(alg[0]))
	}
	if alg[1] < '1' || alg[1] > '8' {
		panic("invalid character in alg : " + string(alg[1]))
	}
	x, y := int(alg[0]-'a'), int(alg[1]-'1')
	if settings.Board.Default.Inverted {
		return invert(x), invert(y)
	}
	return x, y
}

func getRankBox(rank int) Rectangle {
	square := float64(settings.Board.Default.Size) / 8
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      0,
		Y:      border + float64(invert(rank))*square,
		Width:  border,
		Height: square,
	}
}

func getFileBox(file int) Rectangle {
	square := float64(settings.Board.Default.Size) / 8
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      border + float64(invert(file))*square,
		Y:      border + 8*square,
		Width:  square,
		Height: border,
	}
}

func getSquareBox(x, y int) Rectangle {
	square := float64(settings.Board.Default.Size) / 8
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      border + float64(x)*square,
		Y:      border + float64(invert(y))*square,
		Width:  square,
		Height: square,
	}
}

func getBoardBox() Rectangle {
	border := float64(settings.Border.Width)

	return Rectangle{
		X:      border,
		Y:      border,
		Width:  float64(settings.Board.Default.Size),
		Height: float64(settings.Board.Default.Size),
	}
}

func setFontFace(c *gg.Context, size int) {
	path := "roboto.ttf"
	if settings.FontStyle.Path != "" {
		path = settings.FontStyle.Path
	}

	err := c.LoadFontFace(path, float64(size))
	if err != nil {
		panic(fmt.Errorf("failed to load font face : %v", err))
	}
}
