package main

import (
	"image/png"
	"os"

	"github.com/Hultan/chessImager"
)

func main() {
	// Create a new imager using your custom JSON file
	imager, _ := chessImager.NewImagerFromPath("examples/other/other.json")

	// Highlight the e7 square, annotate e7 as a brilliant move (!!) and
	// show move e1-e7.
	ctx := imager.NewContext("b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25").
		AddHighlight("e7").
		AddAnnotation("e7", "!!").
		AddMove("e1", "e7")

	// Render the image
	img, _ := imager.RenderWithContext(ctx)

	// Save the image
	file, _ := os.Create("examples/other/other.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
