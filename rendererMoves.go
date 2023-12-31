package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererMoves struct {
	*Imager
}

func (r *rendererMoves) draw(c *gg.Context) error {
	if r.ctx == nil {
		return nil
	}
	for _, move := range r.ctx.Moves {
		err := r.renderMove(c, move)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererMoves) renderMove(c *gg.Context, move Move) error {
	style := r.getStyle(move)
	var err error

	switch style.Type {
	case MoveTypeDots:
		err = r.renderDottedMove(c, style, move)
	case MoveTypeArrow:
		err = r.renderArrowMove(c, style, move)
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
		return &settings.MoveStyle
	} else {
		return move.Style
	}
}
