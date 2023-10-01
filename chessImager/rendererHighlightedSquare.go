package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererHighlightedSquare struct {
	*Imager
}

func (r *rendererHighlightedSquare) draw(c *gg.Context) {
	for _, high := range r.settings.Highlight {
		style := r.getStyle(high)
		b := r.getSquareBox(r.algToCoords(high.Square))
		c.SetRGBA(toRGBA(style.Color))
		switch style.Type {
		case HighlightedSquareFull:
			x, y, w, h := b.Coords()
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		case HighlightedSquareBorder:
			x, y, w, h := b.Coords()
			bw := float64(style.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
		case HighlightedSquareCircle:
			x, y := b.Center()
			w := float64(style.Width)
			c.DrawCircle(x, y, w)
			c.Fill()
		default:
			panic("rendererHighlightedSquare : oops, why are we here?")
		}
	}
}

func (r *rendererHighlightedSquare) getStyle(high HighlightedSquare) *HighlightedSquareStyle {
	if high.Style == nil {
		return &r.settings.HighlightedSquareStyle
	} else {
		return high.Style
	}
}
