package chessImager

import (
	"bytes"
	_ "embed"
	"errors"
	"image"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

//go:embed pieces.png
var defaultPieces []byte

var pieces map[chessPiece]image.Image

var pieceMap = map[string]chessPiece{
	"WK": WhiteKing,
	"WQ": WhiteQueen,
	"WR": WhiteRook,
	"WN": WhiteKnight,
	"WB": WhiteBishop,
	"WP": WhitePawn,
	"BK": BlackKing,
	"BQ": BlackQueen,
	"BR": BlackRook,
	"BN": BlackKnight,
	"BB": BlackBishop,
	"BP": BlackPawn,
}

var embeddedPieces = []PieceRectangle{
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

type rendererPiece struct {
	*Imager
}

type PieceRectangle struct {
	piece chessPiece
	rect  Rectangle
}

func (r *rendererPiece) draw(c *gg.Context) error {
	err := r.loadPieces()
	if err != nil {
		return err
	}

	fen := normalizeFEN(r.fen)
	fens := strings.Split(fen, "/")

	for rank, row := range fens {
		for file, piece := range normalizeFENRank(row) {
			if p := letter2Piece[piece]; p != NoPiece {
				c.DrawImage(r.getImageAndPosition(pieces[p], file, rank))
			}
		}
	}

	return nil
}

func (r *rendererPiece) loadPieces() error {
	pieces = make(map[chessPiece]image.Image, 12)

	switch settings.Pieces.Type {
	case PiecesTypeDefault:
		imageMap, _, err := image.Decode(bytes.NewReader(defaultPieces))
		if err != nil {
			return err
		}
		err = r.loadImageMapPieces(imageMap, embeddedPieces)
		if err != nil {
			return err
		}
	case PiecesTypeImages:
		for _, piece := range settings.Pieces.Images.Pieces {
			f, err := os.Open(piece.Path)
			if err != nil {
				return err
			}
			img, _, err := image.Decode(f)
			if err != nil {
				return err
			}

			pieces[pieceMap[strings.ToUpper(piece.Piece)]] = r.resize(img)
		}
	case PiecesTypeImageMap:
		f, err := os.Open(settings.Pieces.ImageMap.Path)
		if err != nil {
			return err
		}
		imageMap, _, err := image.Decode(f)
		if err != nil {
			return err
		}
		pr := r.createPieceRectangleSlice(settings.Pieces.ImageMap.Pieces)
		err = r.loadImageMapPieces(imageMap, pr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererPiece) loadImageMapPieces(imageMap image.Image, items []PieceRectangle) error {
	sub, ok := imageMap.(SubImager)
	if !ok {
		return errors.New("failed to create SubImager. Wrong image type? Try PNG")
	}
	for _, item := range items {
		pieces[item.piece] = r.resize(sub.SubImage(image.Rect(item.rect.ToImageRect())))
	}
	return nil
}

func (r *rendererPiece) createPieceRectangleSlice(mapPieces [12]ImageMapPiece) []PieceRectangle {
	result := make([]PieceRectangle, len(mapPieces))
	for _, piece := range mapPieces {
		result = append(result, PieceRectangle{
			piece: pieceMap[strings.ToUpper(piece.Piece)],
			rect:  piece.Rect,
		})
	}
	return result
}

func (r *rendererPiece) resize(img image.Image) image.Image {
	var pieceSize uint

	board := getBoardBox()
	square := board.Width / 8
	switch settings.Board.Type {
	case BoardTypeDefault:
		pieceSize = uint(square * settings.Pieces.Factor)
	case BoardTypeImage:
		pieceSize = uint(square * settings.Pieces.Factor)
	}
	return resize.Resize(pieceSize, pieceSize, img, resize.Lanczos3)
}

func (r *rendererPiece) getImageAndPosition(img image.Image, x, y int) (image.Image, int, int) {
	board := getBoardBox()
	box := getSquareBox(x, y)
	diff := (int(box.Width) - img.Bounds().Size().Y) / 2

	if settings.Board.Default.Inverted {
		return img, int(box.X) + invert(x)*int(box.Width) + diff, int(box.Y) + invert(y)*int(box.Height) + diff
	}

	return img, int(board.X) + x*int(box.Width) + diff, int(board.Y) + y*int(box.Height) + diff
}
