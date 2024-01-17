package chessImager

import (
	"testing"
)

func TestBoardDefault(t *testing.T) {
	filename := "boardDefault.png"

	imager, err := NewImagerFromPath("test/data/boardDefault.json")
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

func TestBoardImage(t *testing.T) {
	filename := "boardImage.png"

	imager, err := NewImagerFromPath("test/data/boardImage.json")
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

func TestBorder(t *testing.T) {
	filename := "border.png"

	imager, err := NewImagerFromPath("test/data/border.json")
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

func TestRankAndFileNone(t *testing.T) {
	filename := "rankAndFileNone.png"

	imager, err := NewImagerFromPath("test/data/rankAndFileNone.json")
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

func TestRankAndFileBorder(t *testing.T) {
	filename := "rankAndFileBorder.png"

	imager, err := NewImagerFromPath("test/data/rankAndFileBorder.json")
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

func TestRankAndFileSquare(t *testing.T) {
	filename := "rankAndFileSquare.png"

	imager, err := NewImagerFromPath("test/data/rankAndFileSquare.json")
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

func TestPiecesDefault(t *testing.T) {
	filename := "piecesDefault.png"

	imager, err := NewImagerFromPath("test/data/piecesDefault.json")
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

func TestPiecesImages(t *testing.T) {
	filename := "piecesImages.png"

	imager, err := NewImagerFromPath("test/data/piecesImages.json")
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

func TestPiecesImageMap(t *testing.T) {
	filename := "piecesImageMap.png"

	imager, err := NewImagerFromPath("test/data/piecesImageMap.json")
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

func TestSetOrderDefault(t *testing.T) {
	filename := "setOrderDefault.png"

	// Create a new imager using embedded default.json settings
	imager := NewImager()

	ctx := imager.NewContext()
	ctx.AddHighlightEx("e7", nil)

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	_ = ctx.SetFEN(fen)

	// Render the image
	img, err := imager.RenderEx(ctx)
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
	imager := NewImager()

	// Set the rendering order
	err := imager.SetOrder([]int{0, 1, 2, 4, 3, 5, 6})
	if err != nil {
		t.Fatalf("failed to set rendering order : %v", err)
	}

	ctx := imager.NewContext()
	ctx.AddHighlightEx("e7", nil)

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	_ = ctx.SetFEN(fen)

	// Render the image
	img, err := imager.RenderEx(ctx)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}
