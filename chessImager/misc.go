package chessImager

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r Rectangle) Coords() (float64, float64, float64, float64) {
	return r.X, r.Y, r.Width, r.Height
}
