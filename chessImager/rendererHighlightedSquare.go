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
		x, y, w, h := r.getSquareBox(r.algToCoords(high.Square)).Coords()
		c.SetRGBA(toRGBA(style.Color))
		switch style.Type {
		case HighlightedSquareFull:
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		case HighlightedSquareBorder:
			bw := float64(style.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
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
