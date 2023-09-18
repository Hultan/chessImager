package chessImager

import "image/color"

// ImageSettings represents the settings for one chess board image.
// If a FEN is specified, the manual setup in Setup is ignored.
type ImageSettings struct {
	Highlight   []HighlightedSquare // List of highlighted squares
	Setup       [8]string           // ex "rnb kbnr" valid chars = "rnbqkpRNBQKP "
	FEN         string              // FEN string (instead of Setup)
	Moves       []Move              // List of marked moves
	Annotations []Annotation        // List of annotations
}

type Annotation struct {
	Square string          // Square position : "a6" or "c5"
	Text   string          // Annotation text, ex "!", "??", "#"
	Style  AnnotationStyle // Annotation style, can be provided globally
}

type AnnotationStyle struct {
	Position        PositionType // TopLeft, TopRight, etc
	Size            int          // Size of the annotation symbol
	BackgroundColor color.Color  // Color of the background
	ForegroundColor color.Color  // Color of the foreground
	BorderColor     color.Color  // Color of the border
}

type Move struct {
	From, To string      // [a-hA-H][1-8] example "A1", "g4", ...
	Color    color.Color // Color of the move arrow
	Type     MoveType    // Type of move arrow
}

type HighlightedSquare struct {
	Square string                // [a-hA-H][1-8] example "A1", "g4", ...
	Color  string                // Color of the marked square
	color  color.RGBA            // Color of the marked square
	Width  int                   // Width of the border (only used if Type = HighlightedSquareBorder)
	Type   HighlightedSquareType // Type of marked square
}
