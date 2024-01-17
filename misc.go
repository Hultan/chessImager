package chessImager

import (
	"errors"
	"fmt"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/exp/constraints"
)

type renderer interface {
	draw(*gg.Context, *ImageContext) error
}

// hexToRGBA converts a hex string to a color
// #RRGGBBAA or #RRGGBB or RRGGBBAA or RRGGBB
func hexToRGBA(hex string) (col color.RGBA, err error) {
	// Remove leading '#' if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex values for red, green, blue and alpha
	if len(hex) == 8 {
		_, err = fmt.Sscanf(hex, "%02x%02x%02x%02x", &col.R, &col.G, &col.B, &col.A)
		if err != nil {
			return col, fmt.Errorf("invalid color (%s) : %v", hex, err)
		}
	} else if len(hex) == 6 {
		col.A = 255
		_, err = fmt.Sscanf(hex, "%02x%02x%02x", &col.R, &col.G, &col.B)
		if err != nil {
			return col, fmt.Errorf("invalid color (%s) : %v", hex, err)
		}
	} else {
		err := errors.New("valid formats : #RRGGBBAA or #RRGGBB or RRGGBBAA or RRGGBB")
		return col, fmt.Errorf("invalid color (%s) : %v", hex, err)
	}

	return col, nil
}

func invert(x int) int {
	return 7 - x
}

func abs[T constraints.Float | constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(dx int) int {
	switch {
	case dx < 0:
		return -1
	case dx == 0:
		return 0
	default:
		return 1
	}
}
