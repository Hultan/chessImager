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
		style := r.getStyle(high)
		a, err := newAlg(high.Square)
		if err != nil {
			return err
		}
		b := getSquareBox(a.coords())
		c.SetRGBA(toRGBA(style.Color))
		switch style.Type {
		case HighlightTypeFull:
			x, y, w, h := b.Coords()
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		case HighlightTypeBorder:
			x, y, w, h := b.Coords()
			bw := float64(style.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
		case HighlightTypeCircle:
			bb := b.Shrink(style.Factor)
			x, y := bb.Center()
			c.SetLineWidth(float64(style.Width))
			c.DrawCircle(x, y, bb.Width/2)
			c.Stroke()
		case HighlightTypeFilledCircle:
			bb := b.Shrink(style.Factor)
			x, y := bb.Center()
			c.DrawCircle(x, y, bb.Width/2)
			c.Fill()
		case HighlightTypeX:
			bb := b.Shrink(style.Factor)
			x, y, w, h := bb.Coords()
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
