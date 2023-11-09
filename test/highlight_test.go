package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestHighlight(t *testing.T) {
	filename := "highlight.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewHighlightStyle(chessImager.HighlightTypeFull, "#333333", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("a1", style)

	style, err = ctx.NewHighlightStyle(chessImager.HighlightTypeCircle, "#444444", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("b2", style)

	style, err = ctx.NewHighlightStyle(chessImager.HighlightTypeFilledCircle, "#555555", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("c3", style)

	style, err = ctx.NewHighlightStyle(chessImager.HighlightTypeBorder, "#666666", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("d4", style)

	style, err = ctx.NewHighlightStyle(chessImager.HighlightTypeX, "#777777", 5, 1)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddHighlightEx("e5", style)

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}
