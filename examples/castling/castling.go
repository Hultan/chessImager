package main

import (
	"image/png"
	"os"

	"github.com/Hultan/chessImager"
)

func main() {
	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Create a new image context, and add white king side castling,
	// and black queen side castling.
	ctx := imager.NewContext("2kr4/8/8/8/8/8/8/5RK1 b - - 1 25").AddMove("0-0", "").AddMove("", "0-0-0")

	// Render the image
	img, _ := imager.RenderWithContext(ctx)

	// Save the image to a file
	file, _ := os.Create("examples/castling/castling.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
