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

type Format int

const (
	Image Format = iota
	PNG
	BMP
)

type RankAndFileType int

const (
	RankAndFileNone RankAndFileType = iota
	RankAndFileInSquares
	RankAndFileInBorder
)

type PositionType int

const (
	TopRight PositionType = iota
	BottomRight
	BottomLeft
	TopLeft
	Middle
)

type MoveType int

const (
	MoveTypeArrow MoveType = iota
	MoveTypeDots
)

type HighlightedSquareType int

const (
	HighlightedSquareFull HighlightedSquareType = iota
	HighlightedSquareBorder
)
