package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestMovesDots(t *testing.T) {
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

func TestMovesKnight(t *testing.T) {
	filename := "movesKnight.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("f6", "g8").AddMove("f6", "h7")
	ctx.AddMove("f6", "h5").AddMove("f6", "g4")
	ctx.AddMove("f6", "e4").AddMove("f6", "d5")
	ctx.AddMove("f6", "d7").AddMove("f6", "e8")

	// Render the image
	const fen = "8/8/5n2/8/8/8/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesBishop(t *testing.T) {
	filename := "movesBishop.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("a1", "f6").AddMove("d8", "f6")
	ctx.AddMove("h8", "f6").AddMove("h4", "f6")

	// Render the image
	const fen = "8/8/5B2/8/8/8/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesRook(t *testing.T) {
	filename := "movesRook.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("f1", "f6").AddMove("f8", "f6")
	ctx.AddMove("a6", "f6").AddMove("h6", "f6")

	// Render the image
	const fen = "8/8/5R2/8/8/8/8/8 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}

func TestMovesCastling(t *testing.T) {
	filename := "movesCastling.png"

	imager := chessImager.NewImager()
	ctx := imager.NewContext()

	ctx.AddMove("h1", "f1").AddMove("e1", "g1")

	// Render the image
	const fen = "8/8/8/8/8/8/8/5rk1 b - - 1 25"
	img, _ := imager.RenderEx(fen, ctx)

	compareImages(t, filename, &img)
}
