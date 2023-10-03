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
			x, y := b.Center()
			w := float64(style.Width)
			c.DrawCircle(x, y, w)
			c.Fill()
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
