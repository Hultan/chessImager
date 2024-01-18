package chessImager

import "image"

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (r Rectangle) coords() (float64, float64, float64, float64) {
	return r.X, r.Y, r.Width, r.Height
}

func (r Rectangle) shrink(f float64) Rectangle {
	wf := r.Width * (1 - f) / 2
	hf := r.Height * (1 - f) / 2

	rr := Rectangle{
		X:      r.X + wf,
		Y:      r.Y + hf,
		Width:  r.Width - wf*2,
		Height: r.Height - hf*2,
	}

	return rr
}

// TODO : Should return an image.Rectangle instead?

func (r Rectangle) toImageRect() image.Rectangle {
	return image.Rect(int(r.X), int(r.Y), int(r.X+r.Width), int(r.Y+r.Height))
}

func (r Rectangle) center() (float64, float64) {
	return r.X + r.Width/2, r.Y + r.Height/2
}
