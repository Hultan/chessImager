package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererHighlight struct {
	*Imager
}

func (r *rendererHighlight) draw(c *gg.Context) {
	if r.ctx == nil {
		return
	}
	for _, high := range r.ctx.Highlight {
		style := r.getStyle(high)
		b := getSquareBox(algToCoords(high.Square))
		c.SetRGBA(toRGBA(style.Color))
		switch style.Type {
		case HighlightFull:
			x, y, w, h := b.Coords()
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		case HighlightBorder:
			x, y, w, h := b.Coords()
			bw := float64(style.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
		case HighlightCircle:
			bb := b.Shrink(style.Factor)
			x, y := bb.Center()
			c.SetLineWidth(float64(style.Width))
			c.DrawCircle(x, y, bb.Width/2)
			c.Stroke()
		case HighlightFilledCircle:
			bb := b.Shrink(style.Factor)
			x, y := bb.Center()
			c.DrawCircle(x, y, bb.Width/2)
			c.Fill()
		case HighlightX:
			bb := b.Shrink(style.Factor)
			x, y, w, h := bb.Coords()
			c.SetLineWidth(float64(style.Width))
			c.DrawLine(x, y, x+w, y+h)
			c.DrawLine(x+w, y, x, y+h)
			c.Stroke()
		default:
			panic("rendererHighlight : oops, why are we here?")
		}
	}
}

func (r *rendererHighlight) getStyle(high HighlightedSquare) *HighlightStyle {
	if high.Style == nil {
		return &settings.HighlightStyle
	} else {
		return high.Style
	}
}
