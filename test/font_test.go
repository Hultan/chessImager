package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestFontStandard(t *testing.T) {
	filename := "fontStandard.png"

	imager, err := chessImager.NewImagerFromPath("data/fontStandard.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestFontRoboto(t *testing.T) {
	filename := "fontRoboto.png"

	imager, err := chessImager.NewImagerFromPath("data/fontRoboto.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}
