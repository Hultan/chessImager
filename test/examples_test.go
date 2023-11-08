package test

import (
	"image"
	"testing"

	"github.com/Hultan/chessImager"
)

func TestSimpleExample(t *testing.T) {
	filename := "simple.png"

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := chessImager.NewImager().Render(fen)
	if err != nil {
		t.Error(err)
	}

	compareImages(t, filename, &img)
}

func TestMediumExample(t *testing.T) {
	filename := "medium.png"

	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Create a new context
	ctx := imager.NewContext()

	// Highlight square e7
	// Annotate square e7 with "!!"
	// Show move e1-e7
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("e1", "e7")

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := imager.RenderEx(fen, ctx)
	if err != nil {
		t.Errorf("failed to render : %v", err)
	}

	compareImages(t, filename, &img)
}

func compareImages(t *testing.T, filename string, img *image.Image) {
	ok, err := compareFiles(img, "valid/"+filename)
	if err != nil {
		t.Errorf("error during compare : %v", err)
	}
	if !ok {
		t.Errorf("failed to compare, images differ!")
	}
}
