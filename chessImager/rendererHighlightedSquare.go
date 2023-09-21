package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererHighlightedSquare struct {
	*Imager
}

func (r *rendererHighlightedSquare) draw(c *gg.Context) {
	for _, high := range r.settings.Highlight {
		x, y, w, h := r.getSquareBox(r.algToCoords(high.Square)).Coords()
		c.SetRGBA(toRGBA(high.color))
		switch high.Type {
		case HighlightedSquareFull:
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		case HighlightedSquareBorder:
			bw := float64(high.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
		default:
			panic("rendererHighlightedSquare : oops, why are we here?")
		}
	}
}
