package chessImager

type BoardType int

const (
	BoardTypeDefault BoardType = iota
	BoardTypeImage
)

type PiecesType int

const (
	PiecesTypeDefault PiecesType = iota
	PiecesTypeImages
	PiecesTypeImageMap
)

type chessPiece int

const (
	WhitePawn chessPiece = iota
	WhiteBishop
	WhiteKnight
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackBishop
	BlackKnight
	BlackRook
	BlackQueen
	BlackKing
	NoPiece
)

type RankAndFileType int

const (
	RankAndFileNone RankAndFileType = iota
	RankAndFileInBorder
	RankAndFileInSquares
)

type PositionType int

const (
	TopRight PositionType = iota
	BottomRight
	BottomLeft
	TopLeft
	Middle
)

type ArrowType int

const (
	MoveTypeArrow ArrowType = iota
	MoveTypeDots
)

type HighlightedSquareType int

const (
	HighlightedSquareFull HighlightedSquareType = iota
	HighlightedSquareBorder
)
