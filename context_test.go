package chessImager

import (
	"testing"
)

func TestMovesDots(t *testing.T) {
	filename := "moves.png"

	imager := NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewMoveStyle(MoveTypeDots, "#333333", "#333333", 0.5)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddMoveEx("a1", "a7", style)
	ctx.AddMoveEx("b1", "e4", style)
	ctx.AddMoveEx("f1", "g3", style)
	ctx.AddMoveEx("g5", "e6", style)

	// Render the image
	const fen = "8/r7/4n3/8/4q3/6n1/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesKnight(t *testing.T) {
	filename := "movesKnight.png"

	imager := NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("f6", "g8").AddMove("f6", "h7")
	ctx.AddMove("f6", "h5").AddMove("f6", "g4")
	ctx.AddMove("f6", "e4").AddMove("f6", "d5")
	ctx.AddMove("f6", "d7").AddMove("f6", "e8")

	// Render the image
	const fen = "8/8/5n2/8/8/8/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesBishop(t *testing.T) {
	filename := "movesBishop.png"

	imager := NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("a1", "f6").AddMove("d8", "f6")
	ctx.AddMove("h8", "f6").AddMove("h4", "f6")

	// Render the image
	const fen = "8/8/5B2/8/8/8/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesRook(t *testing.T) {
	filename := "movesRook.png"

	imager := NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("f1", "f6").AddMove("f8", "f6")
	ctx.AddMove("a6", "f6").AddMove("h6", "f6")

	// Render the image
	const fen = "8/8/5R2/8/8/8/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesCastlingKing(t *testing.T) {
	filename := "movesCastling.png"

	imager := NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("0-0", "").AddMove("", "0-0")

	// Render the image
	const fen = "5rk1/8/8/8/8/8/8/5rk1 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesCastlingQueen(t *testing.T) {
	filename := "movesCastling2.png"

	imager := NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("0-0-0", "").AddMove("", "0-0-0")

	// Render the image
	const fen = "2kr4/8/8/8/8/8/8/2kr4 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesCastlingKing_Dots(t *testing.T) {
	filename := "movesCastlingDots.png"

	imager := NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewMoveStyle(MoveTypeDots, "#33FF33", "#3333FF", 0.2)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddMoveEx("0-0", "", style).AddMoveEx("", "0-0", style)

	// Render the image
	const fen = "5rk1/8/8/8/8/8/8/5rk1 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesCastlingQueen_Dots(t *testing.T) {
	filename := "movesCastling2Dots.png"

	imager := NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewMoveStyle(MoveTypeDots, "#33FF33", "#3333FF", 0.2)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddMoveEx("0-0-0", "", style).AddMoveEx("", "0-0-0", style)

	// Render the image
	const fen = "2kr4/8/8/8/8/8/8/2kr4 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestHighlight(t *testing.T) {
	filename := "highlight.png"

	imager := NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewHighlightStyle(HighlightTypeFull, "#333333", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("a1", style)

	style, err = ctx.NewHighlightStyle(HighlightTypeCircle, "#444444", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("b2", style)

	style, err = ctx.NewHighlightStyle(HighlightTypeFilledCircle, "#555555", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("c3", style)

	style, err = ctx.NewHighlightStyle(HighlightTypeBorder, "#666666", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("d4", style)

	style, err = ctx.NewHighlightStyle(HighlightTypeX, "#777777", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("e5", style)

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestAnnotation(t *testing.T) {
	filename := "annotations.png"

	imager := NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewAnnotationStyle(PositionTypeTopLeft, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("D1", "!", style)

	style, err = ctx.NewAnnotationStyle(PositionTypeTopRight, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("E1", "!", style)

	style, err = ctx.NewAnnotationStyle(PositionTypeBottomRight, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("F1", "!", style)

	style, err = ctx.NewAnnotationStyle(PositionTypeBottomLeft, 20, 17, 1, "#FFFFFFFF", "#000000FF",
		"#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("G1", "!", style)

	style, err = ctx.NewAnnotationStyle(PositionTypeMiddle, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("H1", "!", style)

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}