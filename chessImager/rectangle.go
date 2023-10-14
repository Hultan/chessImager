package chessImager

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (r Rectangle) Coords() (float64, float64, float64, float64) {
	return r.X, r.Y, r.Width, r.Height
}

func (r Rectangle) Shrink(f float64) Rectangle {
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

func (r Rectangle) ToRect() (int, int, int, int) {
	return int(r.X), int(r.Y), int(r.X + r.Width), int(r.Y + r.Height)
}

func (r Rectangle) Center() (float64, float64) {
	return r.X + r.Width/2, r.Y + r.Height/2
}
