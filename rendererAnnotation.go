package chessImager

import (
	"errors"

	"github.com/fogleman/gg"
)

type rendererAnnotation struct {
	*Imager
	ctx *ImageContext
	gg  *gg.Context
}

func (r *rendererAnnotation) draw() error {
	if r.ctx == nil {
		return nil
	}
	for _, annotation := range r.ctx.Annotations {
		rect, err := r.getAnnotationRectangle(annotation)
		if err != nil {
			return err
		}

		r.drawAnnotationCircle(annotation, rect)

		err = r.drawAnnotationText(annotation, rect)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rendererAnnotation) drawAnnotationText(annotation Annotation, rect Rectangle) error {
	x, y := rect.center()
	style := r.getStyle(annotation)
	err := r.setFontFace(r.gg, style.FontSize)

	r.gg.SetRGBA(style.FontColor.toRGBA())
	if err != nil {
		return err
	}
	if r.useInternalFont {
		y -= 3 // SetFontFace/LoadFontFace problem : https://github.com/fogleman/gg/pull/76
	}
	r.gg.DrawStringAnchored(annotation.Text, x, y, 0.5, 0.5)

	return nil
}

func (r *rendererAnnotation) drawAnnotationCircle(annotation Annotation, rect Rectangle) {
	style := r.getStyle(annotation)
	x, y := rect.center()
	r.gg.SetRGBA(style.BorderColor.toRGBA())
	r.gg.DrawCircle(x, y, rect.Width/2)
	r.gg.Fill()
	r.gg.SetRGBA(style.BackgroundColor.toRGBA())
	r.gg.DrawCircle(x, y, rect.Width/2-float64(r.getStyle(annotation).BorderWidth))
	r.gg.Fill()
}

func (r *rendererAnnotation) getAnnotationRectangle(annotation Annotation) (Rectangle, error) {
	square, err := newAlg(annotation.Square, r.inverted)
	if err != nil {
		return Rectangle{}, err
	}

	rect := r.getSquareBox(square.coords())
	style := r.getStyle(annotation)
	size := float64(style.Size)
	const space = 2.0

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
		return &r.settings.AnnotationStyle
	} else {
		return annotation.Style
	}
}
