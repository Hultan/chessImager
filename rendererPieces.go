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

type rendererPiece struct {
	*Imager
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

			pieces[pieceMap[piece.Piece]] = r.resize(img)
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
		return errors.New("failed to create SubImager. Wrong image type? Try PGN")
	}
	for _, item := range items {
		pieces[item.piece] = r.resize(sub.SubImage(image.Rect(item.rect.ToRect())))
	}
	return nil
}

func (r *rendererPiece) createPieceRectangleSlice(mapPieces [12]ImageMapPiece) []PieceRectangle {
	result := make([]PieceRectangle, len(mapPieces))
	for _, piece := range mapPieces {
		result = append(result, PieceRectangle{
			piece: pieceMap[piece.Piece],
			rect:  piece.Rect,
		})
	}
	return result
}

func (r *rendererPiece) resize(img image.Image) image.Image {
	var pieceSize uint

	switch settings.Board.Type {
	case BoardTypeDefault:
		pieceSize = uint(float64(settings.Board.Default.Size/8) * settings.Pieces.Factor)
	case BoardTypeImage:
		// Ok to panic here, not yet implemented
		panic("Not implemented!")
	}
	return resize.Resize(pieceSize, pieceSize, img, resize.Lanczos3)
}

func (r *rendererPiece) getImageAndPosition(img image.Image, x, y int) (image.Image, int, int) {
	square := settings.Board.Default.Size / 8
	border := settings.Border.Width
	diff := (square - img.Bounds().Size().Y) / 2

	if settings.Board.Default.Inverted {
		return img, border + invert(x)*square + diff, border + invert(y)*square + diff
	}

	return img, border + x*square + diff, border + y*square + diff
}
