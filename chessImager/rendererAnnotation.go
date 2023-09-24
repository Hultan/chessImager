package chessImager

import (
	"github.com/fogleman/gg"
)

type rendererAnnotation struct {
	*Imager
}

func (r *rendererAnnotation) draw(c *gg.Context) {
	for _, annotation := range r.settings.Annotations {
		rect := r.getAnnotationRectangle(annotation)

		// Draw annotation circle
		style := r.getStyle(annotation)
		c.SetRGBA(toRGBA(style.BorderColor))
		c.DrawCircle(rect.X+rect.Width/2, rect.Y+rect.Height/2, rect.Width/2)
		c.Fill()
		c.SetRGBA(toRGBA(style.BackgroundColor))
		c.DrawCircle(rect.X+rect.Width/2, rect.Y+rect.Height/2, rect.Width/2-float64(r.getStyle(annotation).BorderWidth))
		c.Fill()

		// Draw annotation text
		c.SetRGBA(toRGBA(style.ForegroundColor))
		r.setFontFace(c, r.getStyle(annotation).FontSize)
		tw, th := c.MeasureString(annotation.Text)
		x := rect.X + (rect.Width-tw)/2
		y := rect.Y + (rect.Height-th)/2 + th - 1
		c.DrawString(annotation.Text, x, y)
	}
}

func (r *rendererAnnotation) getAnnotationRectangle(annotation Annotation) Rectangle {
	rect := r.getSquareBox(r.algToCoords(annotation.Square))

	size := float64(r.getStyle(annotation).Size)
	space := 2.0
	switch r.getStyle(annotation).Position {
	case TopLeft:
		return Rectangle{
			X:      rect.X + space,
			Y:      rect.Y + space,
			Width:  size,
			Height: size,
		}
	case TopRight:
		return Rectangle{
			X:      rect.X + rect.Width - size - space,
			Y:      rect.Y + space,
			Width:  size,
			Height: size,
		}
	case BottomLeft:
		return Rectangle{
			X:      rect.X + space,
			Y:      rect.Y + rect.Height - size - space,
			Width:  size,
			Height: size,
		}
	case BottomRight:
		return Rectangle{
			X:      rect.X + rect.Width - size - space,
			Y:      rect.Y + rect.Height - size - space,
			Width:  size,
			Height: size,
		}
	case Middle:
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
		return &r.settings.AnnotationStyle
	} else {
		return annotation.Style
	}
}
