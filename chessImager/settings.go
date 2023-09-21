package chessImager

import "image/color"

// Settings represents general settings for the ChessImager.
// These settings can be applied once, before generating
// images, or be overridden at any point.
type Settings struct {
	Border      Border              // Width and Color (zero width = no border)
	Board       Board               // Board settings
	RankAndFile RankAndFile         // If and how the rank and file should be drawn
	Pieces      Pieces              // Piece settings
	Highlight   []HighlightedSquare // List of highlighted squares
}

type Border struct {
	Width int
	Color string
	color color.Color
}

type Board struct {
	Type    BoardType    // 0 = default, 1 = BoardImage
	Default BoardDefault // Default board settings
	Image   BoardImage   // Board image settings, if Image.Path is set, the rest is ignored except for inverted
}

type BoardDefault struct {
	Inverted bool        // False : White bottom; True : Black bottom
	Size     int         // The width and height of the chess board (excluding border). Normally divisible by 8.
	White    string      // White color
	white    color.Color // White color
	Black    string      // Black color
	black    color.Color // Black color
}

type BoardImage struct {
	Inverted bool      // False : White bottom; True : Black bottom
	Path     string    // Path to an image of a chessboard, if specified, the rest of the Board settings is ignored
	Board    Rectangle // The rectangle in the image that represents the board
	Size     int       // The size of the piece images, will be centered in the squares
}

type RankAndFile struct {
	Type  RankAndFileType // None, InSquares, InBorder
	Color string          // The font color of the rank and file text
	color color.Color     // The font color of the rank and file text
	Size  int             // The font size of the rank and file text
}

type HighlightedSquare struct {
	Square string                // [a-hA-H][1-8] example "A1", "g4", ...
	Color  string                // Color of the marked square
	color  color.RGBA            // Color of the marked square
	Width  int                   // Width of the border (only used if Type = HighlightedSquareBorder)
	Type   HighlightedSquareType // Type of marked square
}

// Pieces
// If Type = 0 (Default), the embedded pieces will be used
// If Type = 1 (Images), the Paths slice containing 12 paths to individual pics will be used
// If Type = 2 (ImageMapPath), the ImageMapPath and ImageMapCoords will be used
// If a FEN is specified, the manual setup in Setup is ignored.
type Pieces struct {
	Type PiecesType // 0=Default, 1=Images, 2=ImageMap

	Images   Images
	ImageMap ImageMap
}

type Images struct {
	Paths [12]string // List of paths to chess piece images
}

type ImageMap struct {
	Path   string        // Path to an image containing the pieces
	Coords [12]Rectangle // List of rectangles in the map for each piece
}

type Moves struct {
	Moves []Move // List of marked moves
}

type Annotations struct {
	Annotations []Annotation // List of annotations
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
	Type     ArrowType   // Type of move arrow
}
