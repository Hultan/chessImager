package chessImager

import (
	"image/color"
)

// Settings represents general settings for the ChessImager.
// These settings can be applied once, before generating
// images, or be overridden at any point.
type Settings struct {
	Board  Board  // Board settings
	Pieces Pieces // Piece settings
}

type Board struct {
	Image       BoardImage  // Board image settings, if Image.Path is set, the rest is ignored
	Size        int         // The width and height of the chess board (excluding border). Normally divisible by 8.
	Inverted    bool        // False : White bottom; True : Black bottom
	RankAndFile RankAndFile // If and how the rank and file should be drawn
	White       string      // White color
	white       color.Color // White color
	Black       string      // Black color
	black       color.Color // Black color
	Border      Border      // Width and Color (zero width = no border)
}

type BoardImage struct {
	Path          string    // Path to an image of a chessboard, if specified, the rest of the Board settings is ignored
	Board         Rectangle // The rectangle in the image that represents the board
	PieceDistance Size      // The distance between the top left corner of each square and
	// the top left corner of the piece image
}

type RankAndFile struct {
	Type  RankAndFileType // None, InSquares, InBorder
	Color string          // The font color of the rank and file text
	color color.Color     // The font color of the rank and file text
	Size  int             // The font size of the rank and file text
}

// Pieces
// * If ImageMapPath and ImageMapCoords are specified, they will be used
// * If all 12 paths are specified, and ImageMapPath is empty, they will be used
// * Otherwise the default images will be used
type Pieces struct {
	Paths          [12]string    // List of paths to chess piece images
	ImageMapPath   string        // Path to an image containing the pieces
	ImageMapCoords [12]Rectangle // List of rectangles in the map for each piece
}
