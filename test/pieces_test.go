package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestPiecesDefault(t *testing.T) {
	filename := "piecesDefault.png"

	imager, err := chessImager.NewImagerFromPath("data/piecesDefault.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	saveImage(filename, img)
	compareImages(t, filename, &img)
}

func TestPiecesImages(t *testing.T) {
	filename := "piecesImages.png"

	imager, err := chessImager.NewImagerFromPath("data/piecesImages.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	saveImage(filename, img)
	compareImages(t, filename, &img)
}

func TestPiecesImageMap(t *testing.T) {
	filename := "piecesImageMap.png"

	imager, err := chessImager.NewImagerFromPath("data/piecesImageMap.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}
	ctx := imager.NewContext()

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	saveImage(filename, img)
	compareImages(t, filename, &img)
}
