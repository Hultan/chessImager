package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererAnnotation struct {
	*Imager
}

func (r *rendererAnnotation) draw(c *gg.Context) {
	if r.ctx == nil {
		return
	}
	for _, annotation := range r.ctx.Annotations {
		rect := r.getAnnotationRectangle(annotation)

		// Draw annotation circle

		style := r.getStyle(annotation)
		x, y := rect.Center()
		c.SetRGBA(toRGBA(style.BorderColor))
		c.DrawCircle(x, y, rect.Width/2)
		c.Fill()
		c.SetRGBA(toRGBA(style.BackgroundColor))
		c.DrawCircle(x, y, rect.Width/2-float64(r.getStyle(annotation).BorderWidth))
		c.Fill()

		// Draw annotation text
		c.SetRGBA(toRGBA(style.FontColor))
		setFontFace(c, r.getStyle(annotation).FontSize)
		c.DrawStringAnchored(annotation.Text, x, y, 0.5, 0.5)
	}
}

func (r *rendererAnnotation) getAnnotationRectangle(annotation Annotation) Rectangle {
	rect := getSquareBox(algToCoords(annotation.Square))

	size := float64(r.getStyle(annotation).Size)
	space := 2.0
	switch r.getStyle(annotation).Position {
	case PositionTopLeft:
		return Rectangle{
			X:      rect.X + space,
			Y:      rect.Y + space,
			Width:  size,
			Height: size,
		}
	case PositionTopRight:
		return Rectangle{
			X:      rect.X + rect.Width - size - space,
			Y:      rect.Y + space,
			Width:  size,
			Height: size,
		}
	case PositionBottomLeft:
		return Rectangle{
			X:      rect.X + space,
			Y:      rect.Y + rect.Height - size - space,
			Width:  size,
			Height: size,
		}
	case PositionBottomRight:
		return Rectangle{
			X:      rect.X + rect.Width - size - space,
			Y:      rect.Y + rect.Height - size - space,
			Width:  size,
			Height: size,
		}
	case PositionMiddle:
		return Rectangle{
			X:      rect.X + (rect.Width-size)/2,
			Y:      rect.Y + (rect.Height-size)/2,
			Width:  size,
			Height: size,
		}
	}

	return Rectangle{}
}

func (r *rendererAnnotation) getStyle(annotation Annotation) *AnnotationStyle {
	if annotation.Style == nil {
		return &settings.AnnotationStyle
	} else {
		return annotation.Style
	}
}
