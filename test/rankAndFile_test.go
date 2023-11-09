package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestRankAndFileNone(t *testing.T) {
	filename := "rankAndFileNone.png"

	imager, err := chessImager.NewImagerFromPath("data/rankAndFileNone.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestRankAndFileBorder(t *testing.T) {
	filename := "rankAndFileBorder.png"

	imager, err := chessImager.NewImagerFromPath("data/rankAndFileBorder.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestRankAndFileSquare(t *testing.T) {
	filename := "rankAndFileSquare.png"

	imager, err := chessImager.NewImagerFromPath("data/rankAndFileSquare.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}
