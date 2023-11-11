package chessImager

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
	MoveTypeArrow
)

type HighlightType int

const (
	HighlightTypeFull HighlightType = iota
	HighlightTypeBorder
	HighlightTypeCircle
	HighlightTypeFilledCircle
	HighlightTypeX
)

type rankFileType int

const (
	rank rankFileType = iota
	file
)

type chessPiece int

const (
	whitePawn chessPiece = iota
	whiteBishop
	whiteKnight
	whiteRook
	whiteQueen
	whiteKing
	blackPawn
	blackBishop
	blackKnight
	blackRook
	blackQueen
	blackKing
	noPiece
)

type boardType int

const (
	boardTypeDefault boardType = iota
	boardTypeImage
)

type piecesType int

const (
	piecesTypeDefault piecesType = iota
	piecesTypeImages
	piecesTypeImageMap
)

type rankAndFileType int

const (
	rankAndFileTypeNone rankAndFileType = iota
	rankAndFileTypeInBorder
	rankAndFileTypeInSquares
)

type direction int

const (
	directionNorth direction = iota * 45
	directionNorthEast
	directionEast
	directionSouthEast
	directionSouth
	directionSouthWest
	directionWest
	directionNorthWest
)
