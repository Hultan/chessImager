package main

import (
	"image/png"
	"os"

	"chessImager/chessImager"
)

const (
	fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
)

func main() {
	imager := chessImager.NewImager()

	// Simple call
	imgSimple := imager.Render(fen)
	fileSimple, err := os.Create("/home/per/temp/img.png")
	if err != nil {
		panic(err)
	}
	defer fileSimple.Close()
	err = png.Encode(fileSimple, imgSimple)
	if err != nil {
		panic(err)
	}

	// Advanced call
	ctx, err := chessImager.NewContext()
	if err != nil {
		panic(err)
	}

	err = ctx.SetOrder([]int{0, 1, 2, 5, 4, 3, 7})
	if err != nil {
		panic(err)
	}

	hs, err := ctx.NewHighlightStyle(0, "#88008888", 0)
	if err != nil {
		panic(err)
	}
	ctx.AddHighlightEx("e7", hs)

	as, _ := ctx.NewAnnotationStyle(
		chessImager.PositionTopRight,
		18, 15, 1,
		"EEEEEEFF", "000000FF", "000000FF",
	)
	ctx.AddAnnotationEx("e7", "!!", as)

	ms, _ := ctx.NewMoveStyle(chessImager.MoveTypeDots, "80008080", 0.3)
	ctx.AddMoveEx("e1", "e7", ms)

	imgAdv := imager.RenderEx(fen, ctx)
	fileAdv, err := os.Create("/home/per/temp/img2.png")
	if err != nil {
		panic(err)
	}
	defer fileAdv.Close()
	err = png.Encode(fileAdv, imgAdv)
	if err != nil {
		panic(err)
	}
}
