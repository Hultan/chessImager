package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestSimpleExample(t *testing.T) {
	filename := "simple.png"

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := chessImager.NewImager().Render(fen)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestMediumExample(t *testing.T) {
	filename := "medium.png"

	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Create a new image context
	ctx := imager.NewContext()

	// Highlight square e7
	// Annotate square e7 with "!!"
	// Show move e1-e7
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("e1", "e7")

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := imager.RenderEx(fen, ctx)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestAdvancedExample(t *testing.T) {
	filename := "advanced.png"

	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Set the rendering order
	err := imager.SetOrder([]int{0, 1, 2, 3, 5, 4, 6})
	if err != nil {
		t.Fatalf("failed to set rendering order : %v", err)
	}

	// Create a new image context
	ctx := imager.NewContext()

	// Create a highlight style, for the square e7
	hs, err := ctx.NewHighlightStyle(
		chessImager.HighlightTypeFilledCircle, // Highlight type
		"#88E57C",                             // Highlight color
		4,                                     // Highlight circle width
		0.9,                                   // Highlight factor (not used for this Type)
	)
	if err != nil {
		t.Fatalf("failed to create highlight style : %v", err)
	}

	// Create an annotation style, for the square e7
	as, err := ctx.NewAnnotationStyle(
		chessImager.PositionTypeTopLeft, // Position
		25, 20, 1,                       // Size, font size, border width
		"#E8E57C", "#000000", "#FFFFFF", // Background, font, border color
	)
	if err != nil {
		t.Fatalf("failed to create annotation style : %v", err)
	}

	// Create a move style, for the move e1-e7
	ms, err := ctx.NewMoveStyle(
		chessImager.MoveTypeDots, // Move type
		"#9D6B5EFF",              // Dot color
		"#9D6B5EFF",              // Dot color 2
		0.2,                      // Dot size
	)
	if err != nil {
		t.Fatalf("failed to create move style : %v", err)
	}

	// Highlight the e7 square, annotate e7 as a brilliant move (!!) and
	// show move e1-e7.
	ctx.AddHighlightEx("e7", hs).AddAnnotationEx("e7", "!!", as).AddMoveEx("e1", "e7", ms)

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := imager.RenderEx(fen, ctx)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestOtherExample(t *testing.T) {
	filename := "other.png"

	// Create a new imager using your custom JSON file
	imager, err := chessImager.NewImagerFromPath("data/other.json")
	if err != nil {
		t.Fatalf("failed to create imager : %v", err)
	}

	// Create a new image context
	ctx := imager.NewContext()

	// Highlight the e7 square, annotate e7 as a brilliant move (!!) and
	// show move e1-e7.
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("e1", "e7")

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := imager.RenderEx(fen, ctx)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}
