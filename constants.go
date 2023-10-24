package chessImager

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

type RankAndFileType int

const (
	RankAndFileTypeNone RankAndFileType = iota
	RankAndFileTypeInBorder
	RankAndFileTypeInSquares
)

type PositionType int

const (
	PositionTypeTopLeft PositionType = iota
	PositionTypeTopRight
	PositionTypeBottomRight
	PositionTypeBottomLeft
	PositionTypeMiddle
)

type MoveType int

const (
	MoveTypeDots MoveType = iota
)

type HighlightType int

const (
	HighlightTypeFull HighlightType = iota
	HighlightTypeBorder
	HighlightTypeCircle
	HighlightTypeFilledCircle
	HighlightTypeX
)
