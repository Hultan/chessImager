package chessImager

import "errors"

//
// Context is used for advanced chess images
// (images that includes highlighted squares,
// annotations and/or moves)
//

type Context struct {
	settings *Settings

	Highlight   []HighlightedSquare
	Moves       []Move
	Annotations []Annotation
}

func NewContext() (*Context, error) {
	return NewContextFromPath("")
}

func NewContextFromPath(path string) (*Context, error) {
	s, err := loadSettings(path)
	if err != nil {
		return nil, err
	}
	return &Context{settings: s}, nil
}

func (c *Context) SetOrder(order []int) error {
	if len(order) != 7 {
		return errors.New("len(order) must be 7")
	}

	c.settings.Order = order
	return nil
}

func (c *Context) AddHighlight(square string) {
	c.Highlight = append(c.Highlight, HighlightedSquare{Square: square})
}

func (c *Context) AddHighlightEx(square string, style *HighlightStyle) {
	c.Highlight = append(c.Highlight, HighlightedSquare{Square: square, Style: style})
}

func (c *Context) AddAnnotation(square, text string) {
	c.Annotations = append(c.Annotations, Annotation{Square: square, Text: text})
}

func (c *Context) AddAnnotationEx(square, text string, style *AnnotationStyle) {
	c.Annotations = append(c.Annotations, Annotation{Square: square, Text: text, Style: style})
}

func (c *Context) AddMove(from, to string) {
	c.Moves = append(c.Moves, Move{From: from, To: to})
}

func (c *Context) AddMoveEx(from, to string, style *MoveStyle) {
	c.Moves = append(c.Moves, Move{From: from, To: to, Style: style})
}

func (c *Context) NewHighlightStyle(typ HighlightType, color string, width int) (*HighlightStyle, error) {
	col, err := hexToRGBA(color)
	if err != nil {
		return nil, err
	}
	return &HighlightStyle{
		Type:  typ,
		Color: ColorRGBA{col},
		Width: width,
	}, nil
}

func (c *Context) NewAnnotationStyle(pos PositionType, size, fontSize, borderWidth int, bgc, fc,
	bc string) (*AnnotationStyle, error) {

	fCol, err := hexToRGBA(fc)
	if err != nil {
		return nil, err
	}

	bgCol, err := hexToRGBA(bgc)
	if err != nil {
		return nil, err
	}

	bCol, err := hexToRGBA(bc)
	if err != nil {
		return nil, err
	}

	return &AnnotationStyle{
		Position:        pos,
		Size:            size,
		FontColor:       ColorRGBA{fCol},
		FontSize:        fontSize,
		BackgroundColor: ColorRGBA{bgCol},
		BorderColor:     ColorRGBA{bCol},
		BorderWidth:     borderWidth,
	}, nil
}

func (c *Context) NewMoveStyle(typ MoveType, color string, factor float64) (*MoveStyle, error) {
	col, err := hexToRGBA(color)
	if err != nil {
		return nil, err
	}

	return &MoveStyle{
		Color:  ColorRGBA{col},
		Type:   typ,
		Factor: factor,
	}, nil
}
