package main

import (
	"image/png"
	"os"

	"chessImager/chessImager"
)

const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"

func main() {
	imager := chessImager.NewImager()

	// Simple call
	imgSimple, err := imager.Render(fen)
	if err != nil {
		panic(err)
	}
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

	err = ctx.SetOrder([]int{0, 1, 2, 3, 5, 4, 6})
	if err != nil {
		panic(err)
	}

	hs, err := ctx.NewHighlightStyle(0, "#E8E57C", 0)
	if err != nil {
		panic(err)
	}
	ctx.AddHighlightEx("e7", hs)

	as, _ := ctx.NewAnnotationStyle(
		chessImager.PositionTopRight,
		17, 14, 1,
		"E8E57C", "000000", "E8E57C",
	)
	ctx.AddAnnotationEx("e7", "!!", as)

	ms, _ := ctx.NewMoveStyle(chessImager.MoveTypeDots, "#6D6B5EFF", 0.3)
	ctx.AddMoveEx("e1", "e7", ms)

	imgAdv, err := imager.RenderEx(fen, ctx)
	if err != nil {
		panic(err)
	}

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

func generateSample() {
	imager := chessImager.NewImager()
	ctx, _ := chessImager.NewContext()

	// Highlight yellow square
	hs, _ := ctx.NewHighlightStyle(0, "#E8E57CFF", 0)
	ctx.AddHighlightEx("e7", hs)

	// Annotate square e7 with "!!"
	as, _ := ctx.NewAnnotationStyle(
		chessImager.PositionTopRight,
		17, 14, 1,
		"#E8E57CFF", "#000000FF", "#E8E57CFF",
	)
	ctx.AddAnnotationEx("e7", "!!", as)

	// Show move e1-e7
	ms, _ := ctx.NewMoveStyle(chessImager.MoveTypeDots, "#6D6B5EFF", 0.3)
	ctx.AddMoveEx("e1", "e7", ms)

	// Render the image and save it to a file
	img, err := imager.RenderEx(fen, ctx)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create("/home/per/temp/img2.png")
	defer file.Close()
	_ = png.Encode(file, img)
}
