package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestSetOrderDefault(t *testing.T) {
	filename := "setOrderDefault.png"

	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	ctx := imager.NewContext()
	ctx.AddHighlightEx("e7", nil)

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

func TestSetOrderVariant(t *testing.T) {
	filename := "setOrderVariant.png"

	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Set the rendering order
	err := imager.SetOrder([]int{0, 1, 2, 4, 3, 5, 6})
	if err != nil {
		t.Fatalf("failed to set rendering order : %v", err)
	}

	ctx := imager.NewContext()
	ctx.AddHighlightEx("e7", nil)

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
