package chessImager

import (
	"bytes"
	_ "embed"
	"image"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

//go:embed pieces.png
var defaultPieces []byte

var pieces map[chessPiece]image.Image

type rendererPiece struct {
	*Imager
	fen string
}

func (r *rendererPiece) draw(c *gg.Context) {
	if r.settings.Pieces.Type == PiecesTypeDefault {
		r.loadEmbeddedPieces()
	} else {
		panic("not implemented")
	}

	fen := normalizeFEN(r.fen)
	fens := strings.Split(fen, "/")

	for rank, row := range fens {
		for file, piece := range normalizeFENRank(row) {
			if p := letter2Piece[piece]; p != NoPiece {
				c.DrawImage(r.getImageAndPosition(p, file, rank))
			}
		}
	}
}

func (r *rendererPiece) getImageAndPosition(piece chessPiece, x, y int) (image.Image, int, int) {
	square := r.settings.Board.Default.Size / 8
	border := r.settings.Border.Width

	return pieces[piece], border + x*square, border + y*square
}

// TODO : Replace with loadImageMapPieces
func (r *rendererPiece) loadEmbeddedPieces() {
	pieces = make(map[chessPiece]image.Image, 12)
	img, _, err := image.Decode(bytes.NewReader(defaultPieces))
	if err != nil {
		panic(err)
	}
	sub, ok := img.(SubImager)
	if !ok {
		panic("Failed to create SubImager")
	}
	pieces[WhiteKing] = r.resize(sub.SubImage(image.Rect(0, 0, 333, 333)))
	pieces[WhiteQueen] = r.resize(sub.SubImage(image.Rect(333, 0, 666, 333)))
	pieces[WhiteBishop] = r.resize(sub.SubImage(image.Rect(666, 0, 999, 333)))
	pieces[WhiteKnight] = r.resize(sub.SubImage(image.Rect(999, 0, 1332, 333)))
	pieces[WhiteRook] = r.resize(sub.SubImage(image.Rect(1332, 0, 1665, 333)))
	pieces[WhitePawn] = r.resize(sub.SubImage(image.Rect(1665, 0, 1998, 333)))
	pieces[BlackKing] = r.resize(sub.SubImage(image.Rect(0, 333, 333, 666)))
	pieces[BlackQueen] = r.resize(sub.SubImage(image.Rect(333, 333, 666, 666)))
	pieces[BlackBishop] = r.resize(sub.SubImage(image.Rect(666, 333, 999, 666)))
	pieces[BlackKnight] = r.resize(sub.SubImage(image.Rect(999, 333, 1332, 666)))
	pieces[BlackRook] = r.resize(sub.SubImage(image.Rect(1332, 333, 1665, 666)))
	pieces[BlackPawn] = r.resize(sub.SubImage(image.Rect(1665, 333, 1998, 666)))
}

func (r *rendererPiece) resize(img image.Image) image.Image {
	var square uint
	if r.settings.Board.Type == BoardTypeDefault {
		square = uint(r.settings.Board.Default.Size) / 8
	}
	return resize.Resize(square, square, img, resize.Lanczos3)
}
