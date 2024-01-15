package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererHighlight struct {
	*Imager
}

func (r *rendererHighlight) draw(c *gg.Context) error {
	if r.ctx == nil {
		return nil
	}

	for _, high := range r.ctx.Highlight {
		a, err := newAlg(high.Square)
		if err != nil {
			return err
		}
		b := settings.getSquareBox(a.coords(settings.Board.Default.Inverted))

		style := r.getStyle(high)
		c.SetRGBA(style.Color.toRGBA())

		switch style.Type {
		case HighlightTypeFull:
			x, y, w, h := b.coords()
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		case HighlightTypeBorder:
			x, y, w, h := b.coords()
			bw := float64(style.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
		case HighlightTypeCircle:
			bb := b.shrink(style.Factor)
			x, y := bb.center()
			c.SetLineWidth(float64(style.Width))
			c.DrawCircle(x, y, bb.Width/2)
			c.Stroke()
		case HighlightTypeFilledCircle:
			bb := b.shrink(style.Factor)
			x, y := bb.center()
			c.DrawCircle(x, y, bb.Width/2)
			c.Fill()
		case HighlightTypeX:
			bb := b.shrink(style.Factor)
			x, y, w, h := bb.coords()
			c.SetLineWidth(float64(style.Width))
			c.DrawLine(x, y, x+w, y+h)
			c.DrawLine(x+w, y, x, y+h)
			c.Stroke()
		default:
			return errors.New("invalid highlight type")
		}
	}

	return nil
}

func (r *rendererHighlight) getStyle(high HighlightedSquare) *HighlightStyle {
	if high.Style == nil {
		return &settings.HighlightStyle
	} else {
		return high.Style
	}
}
