package chessImager

import (
	"fmt"
	"image/color"
	"strings"
)

// hexToRGBA converts a hex string (#rrggbbaa) to a color
func hexToRGBA(hex string) (col color.RGBA) {
	// Remove the '#' symbol if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex values for red, green, blue and alpha
	// TODO : Handle error
	fmt.Sscanf(hex, "%02x%02x%02x%02x", &col.R, &col.G, &col.B, &col.A)

	return col
}

// convertColors converts all color strings "#FF00BBFF" to color.RGBA
func convertColors(settings *Settings) {
	settings.Board.Default.white = hexToRGBA(settings.Board.Default.White)
	settings.Board.Default.black = hexToRGBA(settings.Board.Default.Black)
	settings.Border.color = hexToRGBA(settings.Border.Color)
	settings.RankAndFile.color = hexToRGBA(settings.RankAndFile.Color)
	for i := range settings.Highlight {
		settings.Highlight[i].color = hexToRGBA(settings.Highlight[i].Color)
	}
}

func toRGBA(col color.Color) (float64, float64, float64, float64) {
	r, g, b, a := col.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
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
