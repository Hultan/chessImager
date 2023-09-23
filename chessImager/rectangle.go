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

func (r Rectangle) ToRect() (int, int, int, int) {
	return int(r.X), int(r.Y), int(r.X + r.Width), int(r.Y + r.Height)
}
