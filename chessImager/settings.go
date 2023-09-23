package chessImager

// TODO : Renderer order
// TODO : Board should be base of a rectangle, even if Board.Type = Default
// This to make the implementation of BoardImage easier.

import "image/color"

// Settings represents general settings for the ChessImager.
// These settings can be applied once, before generating
// images, or be overridden at any point.
// Border: Settings for the border around the chessboard
// Board : Settings for the board
// RankAndFile: If and how the rank and file should be drawn
// Pieces : Piece settings
// Highlight:List of highlighted squares
type Settings struct {
	Border      Border              `json:"border"`
	Board       Board               `json:"board"`
	RankAndFile RankAndFile         `json:"rank_and_file"`
	Pieces      Pieces              `json:"pieces"`
	Highlight   []HighlightedSquare `json:"highlight"`
	Moves       []Move              `json:"moves"`
	Annotations []Annotation        `json:"annotations"`
}

// Border settings for the chessboard
// Width: Width of the border around the chessboard
// Color: Color of the border around the chessboard
type Border struct {
	Width int    `json:"width"`
	Color string `json:"color"`
	color color.Color
}

// Board settings
// Type : 0 = Default, 1 = Image. If Image is set, Border and RankAndFile settings ignored.
// Default : Settings for default drawing of the chessboard
// Image : Settings for using an image of a chessboard as background.
type Board struct {
	Type    BoardType    `json:"type"`
	Default BoardDefault `json:"default"`
	Image   BoardImage   `json:"image"`
}

// BoardDefault represents settings for how the board should be rendered when Board.Type=0 (default).
// Inverted : If false, white will be on bottom.
// Size : Size of the board excluding the border. Normally this value should be divisible by 8.
// White : The color of the light squares
// Black : The color of the dark squares
type BoardDefault struct {
	Inverted bool   `json:"inverted"`
	Size     int    `json:"size"`
	White    string `json:"white"`
	white    color.Color
	Black    string `json:"black"`
	black    color.Color
}

// BoardImage represents settings for rendering the background image of a chessboard (Board.Type=1)
// If you are using BoardImage, BoardDefault will be ignored.
// Inverted : If false, white will be on bottom.
// Path : Path to the background image of a chessboard
// Board : Rectangle that defines where the board is positioned on the image
// Size : The size of the piece images, will be centered in the squares
type BoardImage struct {
	Inverted bool      `json:"inverted"`
	Path     string    `json:"path"`
	Board    Rectangle `json:"board"`
	Size     int       `json:"size"`
}

// RankAndFile defines how the rank and file indicators should be drawn.
// Important: Only used when Board.Type = Default
// Type : 0 = None, 1 = InSquares, 2 = InBorder
// Color : Font color to use
// Size : Font size to use
type RankAndFile struct {
	Type  RankAndFileType `json:"type"`
	Color string          `json:"color"`
	color color.Color
	Size  int `json:"size"`
}

// HighlightedSquare defines how highlighted squares should be drawn.
// Square : The square to be highlighted (ex "f3")
// Color : The highlight color
// Type: 0 = Full square is highlighted, 1 = Only a border around the square is highlighted
// Width : Width of the border (if Type = 1)
type HighlightedSquare struct {
	Square string `json:"square"`
	Color  string `json:"color"`
	color  color.RGBA
	Type   HighlightedSquareType `json:"type"`
	Width  int                   `json:"width"`
}

// Pieces represents settings of how to draw pieces
// Factor : Resize factor for pieces, default = 1 (=100%), pieces will be scaled up or down by factor
// Type: 0 = Embedded pieces, 1 = Images, 2 ImageMap
// Images : Only used if Type=1
// ImageMap : Only used if Type=2
type Pieces struct {
	Factor   float64    `json:"factor"`
	Type     PiecesType `json:"type"`
	Images   Images     `json:"images"`
	ImageMap ImageMap   `json:"image_map"`
}

// Images represents settings for Pieces.Type=1, where each piece is stored as its own image
// Pieces : List of 12 pieces objects containing piece identifier and a path
type Images struct {
	Pieces [12]ImagePiece `json:"pieces"`
}

// ImagePiece represents a single piece.
// Piece : Type of piece, must be one of : "WK","WQ","WR","WN","WB","WP","BK","BQ","BR","BN","BB","BP"
// Path : Path to the pieces image.
type ImagePiece struct {
	Piece string `json:"piece"`
	Path  string `json:"path"`
}

// ImageMap represents settings for Pieces.Type=2, where all 12 pieces are in one image
// Path : Path to the image
// Pieces : 12 pieces and their rectangles that define the pieces in the image
type ImageMap struct {
	Path   string            `json:"path"`
	Pieces [12]ImageMapPiece `json:"pieces"`
}

// ImageMapPiece represents one piece in an image map
// Piece : Type of piece, must be one of : "WK","WQ","WR","WN","WB","WP","BK","BQ","BR","BN","BB","BP"
// Rect : A rectangle that defines where in the image map the piece is located
type ImageMapPiece struct {
	Piece string    `json:"piece"`
	Rect  Rectangle `json:"rect"`
}

// Annotation represents the settings for one annotation
// Square : The square to annotate, ex "f4"
// Text : Extremely short annotation text (usually !,!!,?,??,#...)
// Style : The annotation style
type Annotation struct {
	Square string          `json:"square"`
	Text   string          `json:"text"`
	Style  AnnotationStyle `json:"style"`
}

// AnnotationStyle represents the style for one annotation
// Position : 0=TopRight, 1=BottomRight, 2=BottomLeft, 3=TopLeft, 4=Middle
// Size : Size of annotation
// BackgroundColor : The background color (RGBA, ex "#FF0000FF"
// ForegroundColor : The foreground color (RGBA, ex "#FF0000FF"
// BorderColor : The border color (RGBA, ex "#FF0000FF"
// BorderWidth : The border width
type AnnotationStyle struct {
	Position        PositionType `json:"position"`
	Size            int          `json:"size"`
	BackgroundColor color.Color  `json:"background_color"`
	ForegroundColor color.Color  `json:"foreground_color"`
	BorderColor     color.Color  `json:"border_color"`
	BorderWidth     int          `json:"border_width"`
}

// Move represents a single move arrow on the chessboard.
// From : The from position of the move
// To : The to position of the move
// Color : The color of the arrow
// Type : The arrow type, 0=arrow, 1=dotted
type Move struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Color string `json:"color"`
	color color.Color
	Type  ArrowType `json:"type"`
}
