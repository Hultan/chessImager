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

	// Add white king side castling, and black queen side castling
	ctx.AddMove("0-0", "").AddMove("", "0-0-0")

	// Render the image
	const fen = "2kr4/8/8/8/8/8/8/5RK1 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	// Save the image to a file
	file, _ := os.Create("examples/castling/castling.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
