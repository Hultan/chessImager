package chessImager

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
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

type PieceRectangle struct {
	piece chessPiece
	rect  Rectangle
}

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

// getBoardSize returns a rectangle with the size of the board
// plus the border surrounding it.
func getBoardSize() image.Rectangle {
	switch settings.Board.Type {
	case BoardTypeDefault:
		size := settings.Board.Default.Size + settings.Border.Width*2

		return image.Rectangle{
			Max: image.Point{
				X: size,
				Y: size,
			},
		}
	case BoardTypeImage:
		f, err := os.Open(settings.Board.Image.Path)
		if err != nil {
			panic("error loading board image path: " + err.Error())
		}
		img, _, err := image.Decode(f)
		if err != nil {
			panic("failed to decode board image: " + err.Error())
		}
		return image.Rectangle{
			Max: image.Point{
				X: img.Bounds().Size().X,
				Y: img.Bounds().Size().Y,
			},
		}

	default:
		panic("invalid board type")
	}
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

// loadSettings loads the default settings from a json file
// Path : The path to load the settings from. Leave empty
// for the default settings (config/default.json).
func loadSettings(path string) (*Settings, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	s := &Settings{}
	err = json.NewDecoder(f).Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func loadDefaultSettings() *Settings {
	r := strings.NewReader(defaultSettings)

	s := &Settings{}
	// Ok to panic here, the embedded settings should always be correct
	err := json.NewDecoder(r).Decode(s)
	if err != nil {
		panic(err)
	}

	return s
}
