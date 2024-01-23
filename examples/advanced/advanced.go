package main

import (
	"image/png"
	"os"

	"github.com/Hultan/chessImager"
)

func main() {
	// Create a new imager using embedded default.json settings
	imager := chessImager.NewImager()

	// Set the rendering order
	_ = imager.SetOrder([]int{0, 1, 2, 3, 5, 4, 6})

	// Create a new image context
	ctx := imager.NewContext("b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25")

	// Create a highlight style, for the square e7
	hs, _ := ctx.NewHighlightStyle(
		chessImager.HighlightTypeFilledCircle, // Highlight type
		"#88E57C",                             // Highlight color
		4,                                     // Highlight circle width
		0.9,                                   // Highlight factor (not used for this Type)
	)

	// Create an annotation style, for the square e7
	as, _ := ctx.NewAnnotationStyle(
		chessImager.PositionTypeTopLeft, // Position
		25, 20, 1,                       // Size, font size, border width
		"#E8E57C", "#000000", "#FFFFFF", // Background, font, border color
	)

	// Create a move style, for the move e1-e7
	ms, _ := ctx.NewMoveStyle(
		chessImager.MoveTypeDots, // Move type
		"#9D6B5EFF",              // Dot color
		"#9D6B5EFF",              // Dot color 2
		0.2,                      // Dot size
		0,                        // Padding
	)

	// Highlight the e7 square, annotate e7 as a brilliant move (!!) and
	// show move e1-e7.
	ctx.AddHighlightWithStyle("e7", hs).AddAnnotationWithStyle("e7", "!!", as).AddMoveWithStyle("e1", "e7", ms)

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		panic(err)
	}

	// Save the image
	file, _ := os.Create("examples/advanced/advanced.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
