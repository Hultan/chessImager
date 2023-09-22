package chessImager

import (
	"bytes"
	_ "embed"
	"image"
	"os"
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
	r.loadPieces()

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

func getEmbeddedRectangles() []PieceRectangle {
	return []PieceRectangle{
		{WhiteKing, Rectangle{0, 0, 333, 333}},
		{WhiteQueen, Rectangle{333, 0, 333, 333}},
		{WhiteBishop, Rectangle{666, 0, 333, 333}},
		{WhiteKnight, Rectangle{999, 0, 333, 333}},
		{WhiteRook, Rectangle{1332, 0, 333, 333}},
		{WhitePawn, Rectangle{1665, 0, 333, 333}},
		{BlackKing, Rectangle{0, 333, 333, 333}},
		{BlackQueen, Rectangle{333, 333, 333, 333}},
		{BlackBishop, Rectangle{666, 333, 333, 333}},
		{BlackKnight, Rectangle{999, 333, 333, 333}},
		{BlackRook, Rectangle{1332, 333, 333, 333}},
		{BlackPawn, Rectangle{1665, 333, 333, 333}},
	}
}

func (r *rendererPiece) loadPieces() {
	pieces = make(map[chessPiece]image.Image, 12)

	switch r.settings.Pieces.Type {
	case PiecesTypeDefault:
		pr := getEmbeddedRectangles()
		imageMap, _, err := image.Decode(bytes.NewReader(defaultPieces))
		if err != nil {
			panic(err)
		}
		r.loadImageMapPieces(imageMap, pr)
	case PiecesTypeImages:
		for _, piece := range r.settings.Pieces.Images.Pieces {
			f, err := os.Open(piece.Path)
			if err != nil {
				panic(err)
			}
			img, _, err := image.Decode(f)
			if err != nil {
				panic(err)
			}

			pieces[stringToChessPiece(piece.Piece)] = r.resize(img)
		}
	case PiecesTypeImageMap:
		f, err := os.Open(r.settings.Pieces.ImageMap.Path)
		if err != nil {
			panic(err)
		}
		imageMap, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		pr := imageMapPieceToPieceRectangle(r.settings.Pieces.ImageMap.Pieces)
		r.loadImageMapPieces(imageMap, pr)
	}
}

func imageMapPieceToPieceRectangle(mapPieces [12]ImageMapPiece) []PieceRectangle {
	var result []PieceRectangle
	for _, piece := range mapPieces {
		result = append(result, PieceRectangle{
			piece: stringToChessPiece(piece.Piece),
			rect:  piece.Rect,
		})
	}
	return result
}

func (r *rendererPiece) loadImageMapPieces(imageMap image.Image, items []PieceRectangle) {
	sub, ok := imageMap.(SubImager)
	if !ok {
		panic("Failed to create SubImager")
	}
	for _, item := range items {
		pieces[item.piece] = r.resize(sub.SubImage(image.Rect(item.rect.ToRect())))
	}
}

func (r *rendererPiece) resize(img image.Image) image.Image {
	var square uint
	switch r.settings.Board.Type {
	case BoardTypeDefault:
		square = uint(r.settings.Board.Default.Size) / 8
	case BoardTypeImage:
		panic("Not implemented!")
	}
	return resize.Resize(square, square, img, resize.Lanczos3)
}
