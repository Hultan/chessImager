package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererHighlight struct {
	*Imager
	ctx *ImageContext
	gg  *gg.Context
}

func (r *rendererHighlight) draw() error {
	if r.ctx == nil {
		return nil
	}

	for _, high := range r.ctx.Highlight {
		square, err := newAlg(high.Square, r.inverted)
		if err != nil {
			return err
		}
		b := r.getSquareBox(square.coords())

		style := r.getStyle(high)
		r.gg.SetRGBA(style.Color.toRGBA())

		switch style.Type {
		case HighlightTypeFull:
			r.highlightFull(b)
		case HighlightTypeBorder:
			r.highlightBorder(b, style)
		case HighlightTypeCircle:
			r.highlightCircle(b, style)
		case HighlightTypeFilledCircle:
			r.highlightCircleFilled(b, style)
		case HighlightTypeX:
			r.highlightX(b, style)
		default:
			return errors.New("invalid highlight type")
		}
	}

	return nil
}

func (r *rendererHighlight) highlightX(b Rectangle, style *HighlightStyle) {
	bb := b.shrink(style.Factor)
	x, y, w, h := bb.coords()
	r.gg.SetLineWidth(float64(style.Width))
	r.gg.DrawLine(x, y, x+w, y+h)
	r.gg.DrawLine(x+w, y, x, y+h)
	r.gg.Stroke()
}

func (r *rendererHighlight) highlightCircleFilled(b Rectangle, style *HighlightStyle) {
	bb := b.shrink(style.Factor)
	x, y := bb.center()
	r.gg.DrawCircle(x, y, bb.Width/2)
	r.gg.Fill()
}

func (r *rendererHighlight) highlightCircle(b Rectangle, style *HighlightStyle) {
	bb := b.shrink(style.Factor)
	x, y := bb.center()
	r.gg.SetLineWidth(float64(style.Width))
	r.gg.DrawCircle(x, y, bb.Width/2)
	r.gg.Stroke()
}

func (r *rendererHighlight) highlightBorder(b Rectangle, style *HighlightStyle) {
	x, y, w, h := b.coords()
	bw := float64(style.Width)
	r.gg.SetLineWidth(bw)
	r.gg.DrawRectangle(x+bw/2, y+bw/2, w-bw, h-bw)
	r.gg.Stroke()
}

func (r *rendererHighlight) highlightFull(b Rectangle) {
	x, y, w, h := b.coords()
	r.gg.DrawRectangle(x, y, w, h)
	r.gg.Fill()
}

func (r *rendererHighlight) getStyle(high HighlightedSquare) *HighlightStyle {
	if high.Style == nil {
		return &r.settings.HighlightStyle
	} else {
		return high.Style
	}
}
