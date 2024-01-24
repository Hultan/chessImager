package chessImager

import (
	"testing"
)

func TestBoardDefault(t *testing.T) {
	t.Parallel()

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

func TestBoardInvalidPath(t *testing.T) {
	t.Parallel()

	_, err := NewImagerFromPath("test/data/boardInvalid.json")
	if err == nil {
		t.Fatalf("Invalid path returned no error")
	}
}

func TestBoardImage(t *testing.T) {
	t.Parallel()

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

func TestBoardImageInvalid(t *testing.T) {
	t.Parallel()

	imager, err := NewImagerFromPath("test/data/boardInvalidImagePath.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	_, err = imager.Render(fen)
	if err == nil {
		t.Fatalf("boardInvalidImagePath did not fail")
	}
}

func TestInvalidSetOrder(t *testing.T) {
	t.Parallel()

	imager, err := NewImagerFromPath("test/data/boardImage.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	err = imager.SetOrder([]int{0, 1})
	if err == nil {
		t.Fatalf("SetOrder did not fail")
	}
}

func TestInvalidSetOrderReset(t *testing.T) {
	t.Parallel()

	imager, err := NewImagerFromPath("test/data/boardImage.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	err = imager.SetOrder([]int{})
	if err != nil {
		t.Fatalf("SetOrder reset failed")
	}
}

func TestInvalidSetOrderIndex(t *testing.T) {
	t.Parallel()

	imager, err := NewImagerFromPath("test/data/boardImage.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	err = imager.SetOrder([]int{0, 1, 2, 3, 4, 5, 7})
	if err == nil {
		t.Fatalf("SetOrder with invalid index did not return error")
	}
}

func TestInvalidSetOrderDuplicateIndex(t *testing.T) {
	t.Parallel()

	imager, err := NewImagerFromPath("test/data/boardImage.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	err = imager.SetOrder([]int{0, 1, 2, 3, 3, 5, 6})
	if err == nil {
		t.Fatalf("SetOrder with duplicate index did not return error")
	}
}

func TestInvalidSetOrderJson(t *testing.T) {
	t.Parallel()

	imager, err := NewImagerFromPath("test/data/boardInvalidOrder.json")
	if err != nil {
		t.Fatalf("Failed to load JSON file: %v", err)
	}

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	_, err = imager.Render(fen)
	if err == nil {
		t.Fatalf("boardInvalidOrder did not fail")
	}
}

func TestBorder(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	filename := "setOrderDefault.png"

	// Create a new imager using embedded default.json settings
	imager := NewImager()

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	ctx := imager.NewContext(fen).AddHighlightWithStyle("e7", nil)

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestSetOrderVariant(t *testing.T) {
	t.Parallel()

	filename := "setOrderVariant.png"

	// Create a new imager using embedded default.json settings
	imager := NewImager()

	// Set the rendering order
	err := imager.SetOrder([]int{0, 1, 2, 4, 3, 5, 6})
	if err != nil {
		t.Fatalf("failed to set rendering order : %v", err)
	}

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	ctx := imager.NewContext(fen).AddHighlightWithStyle("e7", nil)

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		t.Fatalf("failed to render : %v", err)
	}
	if img == nil {
		t.Fatalf("image is nil")
	}

	compareImages(t, filename, &img)
}
