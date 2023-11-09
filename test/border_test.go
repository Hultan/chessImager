package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestBorder(t *testing.T) {
	filename := "border.png"

	imager, err := chessImager.NewImagerFromPath("data/border.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := imager.Render(fen)
	if err != nil {
		t.Fatalf("Failed to render chess board: %v", err)
	}

	compareImages(t, filename, &img)
}
