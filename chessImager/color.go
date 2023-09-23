package chessImager

import (
	"encoding/json"
	"fmt"
	"image/color"
)

type ColorRGBA struct {
	color.RGBA
}

func (c *ColorRGBA) MarshalJSON() ([]byte, error) {
	// Encode the color.RGBA as a hexadecimal string
	hexColor := fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
	return json.Marshal(hexColor)
}

func (c *ColorRGBA) UnmarshalJSON(data []byte) error {
	var hexColor string
	if err := json.Unmarshal(data, &hexColor); err != nil {
		return err
	}

	// Parse the hexadecimal string and set it to the color.RGBA
	_, err := fmt.Sscanf(hexColor, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	if err != nil {
		return err
	}

	return nil
}
