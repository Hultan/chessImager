package main

import (
	"image/png"
	"os"

	"github.com/Hultan/chessImager"
)

func main() {
	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Create a new context
	ctx := imager.NewContext()

	// Highlight square e7
	// Annotate square e7 with "!!"
	// Show move e1-e7
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("f6", "h5")

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	// Save the image to a file
	file, _ := os.Create("examples/medium/medium.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
