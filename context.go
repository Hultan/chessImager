package chessImager

//
// Context is used for advanced chess images
// (images that includes highlighted squares,
// annotations and/or moves)
//

type Context struct {
	Highlight   []HighlightedSquare
	Moves       []Move
	Annotations []Annotation
}

// AddHighlight adds a new highlighted square.
func (c *Context) AddHighlight(square string) *Context {
	c.Highlight = append(c.Highlight, HighlightedSquare{Square: square})

	return c
}

// AddHighlightEx adds a new highlighted square with a specific style.
func (c *Context) AddHighlightEx(square string, style *HighlightStyle) *Context {
	c.Highlight = append(c.Highlight, HighlightedSquare{Square: square, Style: style})

	return c
}

// AddAnnotation adds a new annotation.
func (c *Context) AddAnnotation(square, text string) *Context {
	c.Annotations = append(c.Annotations, Annotation{Square: square, Text: text})

	return c
}

// AddAnnotationEx adds a new annotation with a specific style.
func (c *Context) AddAnnotationEx(square, text string, style *AnnotationStyle) *Context {
	c.Annotations = append(c.Annotations, Annotation{Square: square, Text: text, Style: style})

	return c
}

// AddMove adds a move.
func (c *Context) AddMove(from, to string) *Context {
	c.Moves = append(c.Moves, Move{From: from, To: to})

	return c
}

// AddMoveEx adds a move with a specific style.
func (c *Context) AddMoveEx(from, to string, style *MoveStyle) *Context {
	c.Moves = append(c.Moves, Move{From: from, To: to, Style: style})

	return c
}

// NewHighlightStyle creates a new highlight style.
func (c *Context) NewHighlightStyle(typ HighlightType, color string, width int, factor float64) (*HighlightStyle, error) {
	col, err := hexToRGBA(color)
	if err != nil {
		return nil, err
	}
	return &HighlightStyle{
		Type:   typ,
		Color:  ColorRGBA{col},
		Width:  width,
		Factor: factor,
	}, nil
}

// NewAnnotationStyle creates a new annotation style.
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

// NewMoveStyle creates a new move style.
func (c *Context) NewMoveStyle(typ MoveType, color string, color2 string, factor float64) (*MoveStyle, error) {
	col, err := hexToRGBA(color)
	if err != nil {
		return nil, err
	}

	col2, err := hexToRGBA(color2)
	if err != nil {
		return nil, err
	}

	return &MoveStyle{
		Color:  ColorRGBA{col},
		Color2: ColorRGBA{col2},
		Type:   typ,
		Factor: factor,
	}, nil
}
