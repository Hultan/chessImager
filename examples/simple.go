package main

import (
	"image/png"
	"os"

	"chessImager"
)

func main() {
	// Render simple image
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, _ := chessImager.NewImager().Render(fen)

	// Save image
	file, _ := os.Create("examples/simple.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
