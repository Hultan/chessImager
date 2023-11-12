package chessImager

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"strings"
)

type ColorRGBA struct {
	color.RGBA
}

func (c *ColorRGBA) MarshalJSON() ([]byte, error) {
	// Encode the color.RGBA as a hexadecimal string
	hexColor := fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
	return json.Marshal(hexColor)
}

func (c *ColorRGBA) UnmarshalJSON(data []byte) (err error) {
	var hexColor string

	if err = json.Unmarshal(data, &hexColor); err != nil {
		return err
	}

	// Catch invalid colors
	if len(hexColor) < 6 || len(hexColor) > 9 {
		return errors.New("invalid color")
	}

	// Remove the # and add the alpha if needed
	hexColor = strings.TrimPrefix(hexColor, "#")
	if len(hexColor) == 6 {
		hexColor += "FF"
	}

	// Parse the hexadecimal string and set it to the color.RGBA
	_, err = fmt.Sscanf(hexColor, "%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	if err != nil {
		return err
	}

	return nil
}

func (c *ColorRGBA) toRGBA() (float64, float64, float64, float64) {
	return float64(c.R) / 255, float64(c.G) / 255, float64(c.B) / 255, float64(c.A) / 255
}
