package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestMoves(t *testing.T) {
	filename := "moves.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	style, err := ctx.NewMoveStyle(chessImager.MoveTypeDots, "#333333", 0.5)
	if err != nil {
		t.Fatalf("Failed to create a highlight style: %v", err)
	}
	ctx.AddMoveEx("a1", "a7", style)
	ctx.AddMoveEx("b1", "e4", style)
	ctx.AddMoveEx("f1", "g3", style)
	ctx.AddMoveEx("g5", "e6", style)

	// Render the image
	const fen = "8/r7/4n3/8/4q3/6n1/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}
