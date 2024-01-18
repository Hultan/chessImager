package chessImager

import "fmt"

//
// ImageContext is used for advanced chess images
// (advanced images are images that includes a FEN
// string and optionally highlighted squares,
// annotations and/or moves.
//

type ImageContext struct {
	Fen         string
	Highlight   []HighlightedSquare
	Moves       []Move
	Annotations []Annotation
}

// SetFEN adds a FEN string to the ImageContext object
func (c *ImageContext) SetFEN(fen string) error {
	if !validateFen(fen) {
		return fmt.Errorf("invalid FEN : %v", fen)
	}

	c.Fen = fen
	return nil
}

// AddHighlight adds a new highlighted square.
func (c *ImageContext) AddHighlight(square string) *ImageContext {
	c.Highlight = append(c.Highlight, HighlightedSquare{Square: square})

	return c
}

// AddHighlightWithStyle adds a new highlighted square with a specific style.
func (c *ImageContext) AddHighlightWithStyle(square string, style *HighlightStyle) *ImageContext {
	c.Highlight = append(c.Highlight, HighlightedSquare{Square: square, Style: style})

	return c
}

// AddAnnotation adds a new annotation.
func (c *ImageContext) AddAnnotation(square, text string) *ImageContext {
	c.Annotations = append(c.Annotations, Annotation{Square: square, Text: text})

	return c
}

// AddAnnotationWithStyle adds a new annotation with a specific style.
func (c *ImageContext) AddAnnotationWithStyle(square, text string, style *AnnotationStyle) *ImageContext {
	c.Annotations = append(c.Annotations, Annotation{Square: square, Text: text, Style: style})

	return c
}

// AddMove adds a move.
func (c *ImageContext) AddMove(from, to string) *ImageContext {
	c.Moves = append(c.Moves, Move{From: from, To: to})

	return c
}

// AddMoveWithStyle adds a move with a specific style.
func (c *ImageContext) AddMoveWithStyle(from, to string, style *MoveStyle) *ImageContext {
	c.Moves = append(c.Moves, Move{From: from, To: to, Style: style})

	return c
}

// NewHighlightStyle creates a new highlight style.
func (c *ImageContext) NewHighlightStyle(typ HighlightType, color string, width int, factor float64) (*HighlightStyle, error) {
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
func (c *ImageContext) NewAnnotationStyle(pos PositionType, size, fontSize, borderWidth int, bgc, fc,
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
func (c *ImageContext) NewMoveStyle(typ MoveType, color string, color2 string, factor float64) (*MoveStyle, error) {
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
