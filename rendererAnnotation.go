package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererAnnotation struct {
	*Imager
}

func (r *rendererAnnotation) draw(c *gg.Context) error {
	if r.ctx == nil {
		return nil
	}
	for _, annotation := range r.ctx.Annotations {
		rect, err := r.getAnnotationRectangle(annotation)
		if err != nil {
			return err
		}

		// Draw annotation circle
		style := r.getStyle(annotation)
		x, y := rect.center()
		c.SetRGBA(style.BorderColor.toRGBA())
		c.DrawCircle(x, y, rect.Width/2)
		c.Fill()
		c.SetRGBA(style.BackgroundColor.toRGBA())
		c.DrawCircle(x, y, rect.Width/2-float64(r.getStyle(annotation).BorderWidth))
		c.Fill()

		// Draw annotation text
		c.SetRGBA(style.FontColor.toRGBA())
		err = r.setFontFace(settings.FontStyle.Path, c, r.getStyle(annotation).FontSize)
		if err != nil {
			return err
		}
		if r.useInternalFont {
			y -= 3 // SetFontFace/LoadFontFace problem : https://github.com/fogleman/gg/pull/76
		}
		c.DrawStringAnchored(annotation.Text, x, y, 0.5, 0.5)
	}

	return nil
}

func (r *rendererAnnotation) getAnnotationRectangle(annotation Annotation) (Rectangle, error) {
	a, err := newAlg(annotation.Square)
	if err != nil {
		return Rectangle{}, err
	}

	rect := settings.getSquareBox(a.coords(settings.Board.Default.Inverted))
	style := r.getStyle(annotation)
	size := float64(style.Size)
	space := 2.0

	switch style.Position {
	case PositionTypeTopLeft:
		return Rectangle{
			X:      rect.X + space,
			Y:      rect.Y + space,
			Width:  size,
			Height: size,
		}, nil
	case PositionTypeTopRight:
		return Rectangle{
			X:      rect.X + rect.Width - size - space,
			Y:      rect.Y + space,
			Width:  size,
			Height: size,
		}, nil
	case PositionTypeBottomLeft:
		return Rectangle{
			X:      rect.X + space,
			Y:      rect.Y + rect.Height - size - space,
			Width:  size,
			Height: size,
		}, nil
	case PositionTypeBottomRight:
		return Rectangle{
			X:      rect.X + rect.Width - size - space,
			Y:      rect.Y + rect.Height - size - space,
			Width:  size,
			Height: size,
		}, nil
	case PositionTypeMiddle:
		return Rectangle{
			X:      rect.X + (rect.Width-size)/2,
			Y:      rect.Y + (rect.Height-size)/2,
			Width:  size,
			Height: size,
		}, nil
	default:
		return Rectangle{}, errors.New("invalid position type")
	}
}

func (r *rendererAnnotation) getStyle(annotation Annotation) *AnnotationStyle {
	if annotation.Style == nil {
		return &settings.AnnotationStyle
	} else {
		return annotation.Style
	}
}
