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

type rendererPiece struct {
	*Imager
	ctx *ImageContext
	gg  *gg.Context
}

type PieceRectangle struct {
	piece chessPiece
	rect  Rectangle
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (r *rendererPiece) draw() error {
	err := r.loadPieces()
	if err != nil {
		return err
	}

	// FEN is validated before rendering starts,
	// so it should be OK here.
	fen := normalizeFEN(r.ctx.Fen)
	fens := strings.Split(fen, "/")

	var inv = r.settings.Board.Default.Inverted
	if r.settings.Board.Type == boardTypeImage {
		inv = r.settings.Board.Image.Inverted
	}

	for rank, row := range fens {
		for file, piece := range row {
			if p := letter2Piece[piece]; p != noPiece {
				r.gg.DrawImage(r.getImageAndPosition(r.ctx.pieces[p], file, rank, inv))
			}
		}
	}

	return nil
}

func (r *rendererPiece) loadPieces() error {
	r.ctx.pieces = make(map[chessPiece]image.Image, 12)

	r.loadPieceMap()
	r.loadEmbeddedPieceMap()

	switch r.settings.Pieces.Type {
	case piecesTypeDefault:
		imageMap, _, err := image.Decode(bytes.NewReader(defaultPieces))
		if err != nil {
			return err
		}
		err = r.loadImageMapPieces(imageMap, r.ctx.embeddedPieces)
		if err != nil {
			return err
		}
	case piecesTypeImages:
		for _, piece := range r.settings.Pieces.Images.Pieces {
			f, err := os.Open(piece.Path)
			if err != nil {
				return err
			}
			img, _, err := image.Decode(f)
			if err != nil {
				return err
			}

			r.ctx.pieces[r.ctx.pieceMap[strings.ToUpper(piece.Piece)]] = r.resize(img)
		}
	case piecesTypeImageMap:
		f, err := os.Open(r.settings.Pieces.ImageMap.Path)
		if err != nil {
			return err
		}
		imageMap, _, err := image.Decode(f)
		if err != nil {
			return err
		}
		pr := r.createPieceRectangleSlice(r.settings.Pieces.ImageMap.Pieces)
		err = r.loadImageMapPieces(imageMap, pr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererPiece) loadImageMapPieces(imageMap image.Image, pr []PieceRectangle) error {
	sub, ok := imageMap.(SubImager)
	if !ok {
		return errors.New("failed to create SubImager. Wrong image type? Try PNG")
	}
	for _, item := range pr {
		r.ctx.pieces[item.piece] = r.resize(sub.SubImage(item.rect.toImageRect()))
	}
	return nil
}

func (r *rendererPiece) createPieceRectangleSlice(mapPieces [12]ImageMapPiece) []PieceRectangle {
	result := make([]PieceRectangle, len(mapPieces))
	for _, piece := range mapPieces {
		result = append(result, PieceRectangle{
			piece: r.ctx.pieceMap[strings.ToUpper(piece.Piece)],
			rect:  piece.Rect,
		})
	}
	return result
}

func (r *rendererPiece) resize(img image.Image) image.Image {
	board := r.getBoardBox()
	pieceSize := uint(board.Width * r.settings.Pieces.Factor / 8)
	return resize.Resize(pieceSize, pieceSize, img, resize.Lanczos3)
}

func (r *rendererPiece) getImageAndPosition(img image.Image, x, y int, inv bool) (image.Image, int, int) {
	board := r.getBoardBox()
	box := r.getSquareBox(x, y)
	diff := (int(box.Width) - img.Bounds().Size().Y) / 2

	if inv {
		return img, int(board.X) + invert(x)*int(box.Width) + diff, int(board.Y) + invert(y)*int(box.Height) + diff
	}

	return img, int(board.X) + x*int(box.Width) + diff, int(board.Y) + y*int(box.Height) + diff
}

func (r *rendererPiece) loadPieceMap() {
	r.ctx.pieceMap = map[string]chessPiece{
		"WK": whiteKing,
		"WQ": whiteQueen,
		"WR": whiteRook,
		"WN": whiteKnight,
		"WB": whiteBishop,
		"WP": whitePawn,
		"BK": blackKing,
		"BQ": blackQueen,
		"BR": blackRook,
		"BN": blackKnight,
		"BB": blackBishop,
		"BP": blackPawn,
	}
}

func (r *rendererPiece) loadEmbeddedPieceMap() {
	r.ctx.embeddedPieces = []PieceRectangle{
		{whiteKing, Rectangle{0, 0, 333, 333}},
		{whiteQueen, Rectangle{333, 0, 333, 333}},
		{whiteBishop, Rectangle{666, 0, 333, 333}},
		{whiteKnight, Rectangle{999, 0, 333, 333}},
		{whiteRook, Rectangle{1332, 0, 333, 333}},
		{whitePawn, Rectangle{1665, 0, 333, 333}},
		{blackKing, Rectangle{0, 333, 333, 333}},
		{blackQueen, Rectangle{333, 333, 333, 333}},
		{blackBishop, Rectangle{666, 333, 333, 333}},
		{blackKnight, Rectangle{999, 333, 333, 333}},
		{blackRook, Rectangle{1332, 333, 333, 333}},
		{blackPawn, Rectangle{1665, 333, 333, 333}},
	}
}
