package chessImager

import (
	_ "embed"
	"image"

	"github.com/fogleman/gg"
)

//go:embed pieces.png
var defaultPieces []byte

var pieces map[chessPiece]image.Image

type pieceRenderer struct {
	*Imager
}

func (r *pieceRenderer) draw(c *gg.Context, _ ImageSettings) {
	c.DrawImage(r.getImageAndPosition(BlackRook, 0, 7))
	c.DrawImage(r.getImageAndPosition(BlackKnight, 1, 7))
	c.DrawImage(r.getImageAndPosition(BlackBishop, 2, 7))
	c.DrawImage(r.getImageAndPosition(BlackQueen, 3, 7))
	c.DrawImage(r.getImageAndPosition(BlackKing, 4, 7))
	c.DrawImage(r.getImageAndPosition(BlackBishop, 5, 7))
	c.DrawImage(r.getImageAndPosition(BlackKnight, 6, 7))
	c.DrawImage(r.getImageAndPosition(BlackRook, 7, 7))
	for i := 0; i < 8; i++ {
		c.DrawImage(r.getImageAndPosition(BlackPawn, i, 6))
		c.DrawImage(r.getImageAndPosition(WhitePawn, i, 1))
	}
	c.DrawImage(r.getImageAndPosition(WhiteRook, 0, 0))
	c.DrawImage(r.getImageAndPosition(WhiteKnight, 1, 0))
	c.DrawImage(r.getImageAndPosition(WhiteBishop, 2, 0))
	c.DrawImage(r.getImageAndPosition(WhiteQueen, 3, 0))
	c.DrawImage(r.getImageAndPosition(WhiteKing, 4, 0))
	c.DrawImage(r.getImageAndPosition(WhiteBishop, 5, 0))
	c.DrawImage(r.getImageAndPosition(WhiteKnight, 6, 0))
	c.DrawImage(r.getImageAndPosition(WhiteRook, 7, 0))
}

func (r *pieceRenderer) getImageAndPosition(piece chessPiece, x, y int) (image.Image, int, int) {
	square := r.settings.Board.Size / 8
	border := r.settings.Board.Border.Width

	return pieces[piece], border + x*square, border + (7-y)*square
}
