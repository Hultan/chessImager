package chessImager

import (
	"github.com/fogleman/gg"
)

type highlightedSquareRenderer struct {
	*Imager
}

func (r *highlightedSquareRenderer) draw(c *gg.Context, settings ImageSettings) {
	for _, s := range settings.Highlight {
		x, y, w, h := r.getSquareBox(r.algToCoords(s.Square)).Coords()
		c.SetRGBA(toRGBA(s.color))
		if s.Type == HighlightedSquareFull {
			c.DrawRectangle(x, y, w, h)
			c.Fill()
		} else {
			// HighlightedSquareBordered
			bw := float64(s.Width)
			c.SetLineWidth(bw)
			c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
			c.Stroke()
		}
	}
}
