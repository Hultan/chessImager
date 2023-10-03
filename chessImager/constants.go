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
	PositionTopLeft PositionType = iota
	PositionTopRight
	PositionBottomRight
	PositionBottomLeft
	PositionMiddle
)

type MoveType int

const (
	MoveTypeArrow  MoveType = iota // Arrow from -> to
	MoveTypeLine                   // Line from -> to
	MoveTypeDots                   // Dots from and to
	MoveTypeDotsEx                 // Dots from -> to
)

type HighlightType int

const (
	HighlightFull HighlightType = iota
	HighlightBorder
	HighlightCircle
)
