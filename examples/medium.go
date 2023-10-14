package main

import (
	"image/png"
	"os"

	"chessImager"
)

func main() {
	imager := chessImager.NewImager()
	ctx, _ := imager.NewContext() // New context using default.json settings

	// Highlight square e7
	// Annotate square e7 with "!!"
	// Show move e1-e7
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("e1", "e7")

	// Render the image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	// Save the image to a file
	file, _ := os.Create("examples/medium.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
