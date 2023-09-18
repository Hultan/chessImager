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

func toRGBA(col color.Color) (float64, float64, float64, float64) {
	r, g, b, a := col.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}
