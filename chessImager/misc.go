package chessImager

import "image/color"

type Position struct {
	Top  int
	Left int
}

type Size struct {
	Width  int
	Height int
}

type Rectangle struct {
	Position
	Size
}

type Border struct {
	Width int
	Color string
	color color.Color
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
