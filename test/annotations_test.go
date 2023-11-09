package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestAnnotation(t *testing.T) {
	filename := "annotations.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewAnnotationStyle(chessImager.PositionTypeTopLeft, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("D1", "!", style)

	style, err = ctx.NewAnnotationStyle(chessImager.PositionTypeTopRight, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("E1", "!", style)

	style, err = ctx.NewAnnotationStyle(chessImager.PositionTypeBottomRight, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("F1", "!", style)

	style, err = ctx.NewAnnotationStyle(chessImager.PositionTypeBottomLeft, 20, 17, 1, "#FFFFFFFF", "#000000FF",
		"#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("G1", "!", style)

	style, err = ctx.NewAnnotationStyle(chessImager.PositionTypeMiddle, 20, 17, 1, "#FFFFFFFF", "#000000FF", "#FF0000FF")
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddAnnotationEx("H1", "!", style)

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}
