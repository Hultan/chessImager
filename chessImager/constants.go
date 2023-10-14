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
	MoveTypeDots  MoveType = iota // Dots from and to
	moveTypeArrow                 // Arrow from -> to, NOT IMPLEMENTED YET!
)

type HighlightType int

const (
	HighlightFull HighlightType = iota
	HighlightBorder
	HighlightCircle
	HighlightFilledCircle
	HighlightX
)
