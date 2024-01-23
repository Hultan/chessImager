package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererMoves struct {
	*Imager
	ctx *ImageContext
	gg  *gg.Context
}

func (r *rendererMoves) draw() error {
	if r.ctx == nil {
		return nil
	}

	for _, move := range r.ctx.Moves {
		err := r.renderMove(move)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererMoves) renderMove(move Move) error {
	var err error
	style := r.getStyle(move)

	switch style.Type {
	case MoveTypeDots:
		err = r.renderDottedMove(style, move)
	case MoveTypeArrow:
		err = r.renderArrowMove(style, move)
	default:
		err = errors.New("illegal move type")
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *rendererMoves) getStyle(move Move) *MoveStyle {
	if move.Style == nil {
		return &r.settings.MoveStyle
	} else {
		return move.Style
	}
}
