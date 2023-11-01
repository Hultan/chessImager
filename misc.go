package chessImager

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type renderer interface {
	draw(*gg.Context) error
}

// hexToRGBA converts a hex string to a color
// #RRGGBBAA or #RRGGBB or RRGGBBAA or RRGGBB
func hexToRGBA(hex string) (col color.RGBA, err error) {
	// Remove the '#' symbol if it exists
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

func toRGBA(col ColorRGBA) (float64, float64, float64, float64) {
	return float64(col.R) / 255, float64(col.G) / 255, float64(col.B) / 255, float64(col.A) / 255
}

func invert(x int) int {
	return 7 - x
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

func getBoardBox() Rectangle {
	switch settings.Board.Type {
	case BoardTypeDefault:
		border := float64(settings.Border.Width)
		size := float64(settings.Board.Default.Size)

		return Rectangle{
			X:      border,
			Y:      border,
			Width:  size,
			Height: size,
		}
	case BoardTypeImage:
		return settings.Board.Image.Rect
	default:
		panic("invalid board type")
	}
}

func getSquareBox(x, y int) Rectangle {
	board := getBoardBox()
	square := board.Width / 8

	if settings.Board.Type == BoardTypeDefault {
		border := float64(settings.Border.Width)

		return Rectangle{
			X:      border + float64(x)*square,
			Y:      border + float64(invert(y))*square,
			Width:  square,
			Height: square,
		}
	} else {
		return Rectangle{
			X:      board.X + float64(x)*square,
			Y:      board.Y + float64(invert(y))*square,
			Width:  square,
			Height: square,
		}
	}
}

func validateAlg(alg string) error {
	alg = strings.ToLower(alg)
	if len(alg) != 2 {
		return errors.New("invalid length of alg")
	}
	if alg[0] < 'a' || alg[0] > 'h' {
		return errors.New("invalid character in alg : " + string(alg[0]))
	}
	if alg[1] < '1' || alg[1] > '8' {
		return errors.New("invalid character in alg : " + string(alg[1]))
	}
	return nil
}
func algToCoords(alg string) (int, int) {
	alg = strings.ToLower(alg)
	x, y := int(alg[0]-'a'), int(alg[1]-'1')
	if settings.Board.Default.Inverted {
		return invert(x), invert(y)
	}
	return x, y
}

func setFontFace(c *gg.Context, size int) error {
	if settings.FontStyle.Path == "" {
		// Use standard font
		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			log.Fatal(err)
		}

		face := truetype.NewFace(font, &truetype.Options{Size: float64(size)})
		c.SetFontFace(face)
		useInternalFont = true
	} else {
		// Load font specified in config file
		err := c.LoadFontFace(settings.FontStyle.Path, float64(size))
		if err != nil {
			return fmt.Errorf("failed to load font face : %v", err)
		}
		useInternalFont = false
	}

	return nil
}
