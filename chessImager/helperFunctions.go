package chessImager

import (
	"fmt"
	"image/color"
	"strings"
)

// hexToRGBA converts a hex string (#rrggbbaa) to a color
func hexToRGBA(hex string) color.RGBA {
	// Remove the '#' symbol if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex values for red, green, blue and alpha
	var r, g, b, a uint8
	fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)

	return color.RGBA{R: r, G: g, B: b, A: a}
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
