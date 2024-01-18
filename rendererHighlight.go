package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererHighlight struct {
	*Imager
}

func (r *rendererHighlight) draw(c *gg.Context, ctx *ImageContext) error {
	if ctx == nil {
		return nil
	}

	for _, high := range ctx.Highlight {
		square, err := newAlg(high.Square, r.settings.Board.Default.Inverted)
		if err != nil {
			return err
		}
		b := r.getSquareBox(square.coords())

		style := r.getStyle(high)
		c.SetRGBA(style.Color.toRGBA())

		switch style.Type {
		case HighlightTypeFull:
			r.highlightFull(c, b)
		case HighlightTypeBorder:
			r.highlightBorder(c, b, style)
		case HighlightTypeCircle:
			r.highlightCircle(c, b, style)
		case HighlightTypeFilledCircle:
			r.highlightCircleFilled(c, b, style)
		case HighlightTypeX:
			r.highlightX(c, b, style)
		default:
			return errors.New("invalid highlight type")
		}
	}

	return nil
}

func (r *rendererHighlight) highlightX(c *gg.Context, b Rectangle, style *HighlightStyle) {
	bb := b.shrink(style.Factor)
	x, y, w, h := bb.coords()
	c.SetLineWidth(float64(style.Width))
	c.DrawLine(x, y, x+w, y+h)
	c.DrawLine(x+w, y, x, y+h)
	c.Stroke()
}

func (r *rendererHighlight) highlightCircleFilled(c *gg.Context, b Rectangle, style *HighlightStyle) {
	bb := b.shrink(style.Factor)
	x, y := bb.center()
	c.DrawCircle(x, y, bb.Width/2)
	c.Fill()
}

func (r *rendererHighlight) highlightCircle(c *gg.Context, b Rectangle, style *HighlightStyle) {
	bb := b.shrink(style.Factor)
	x, y := bb.center()
	c.SetLineWidth(float64(style.Width))
	c.DrawCircle(x, y, bb.Width/2)
	c.Stroke()
}

func (r *rendererHighlight) highlightBorder(c *gg.Context, b Rectangle, style *HighlightStyle) {
	x, y, w, h := b.coords()
	bw := float64(style.Width)
	c.SetLineWidth(bw)
	c.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
	c.Stroke()
}

func (r *rendererHighlight) highlightFull(c *gg.Context, b Rectangle) {
	x, y, w, h := b.coords()
	c.DrawRectangle(x, y, w, h)
	c.Fill()
}

func (r *rendererHighlight) getStyle(high HighlightedSquare) *HighlightStyle {
	if high.Style == nil {
		return &r.settings.HighlightStyle
	} else {
		return high.Style
	}
}
