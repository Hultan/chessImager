package main

import (
	"image/png"
	"os"

	"chessImager/chessImager"
)

const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"

func main() {
	renderSimple("/home/per/temp/simple.png")
	renderKasparov("/home/per/temp/kasparov.png")
	renderAdvanced("/home/per/temp/advanced.png")
}

func renderAdvanced(fileName string) {
	imager := chessImager.NewImager()

	// Advanced call
	ctx, err := imager.NewContext()
	if err != nil {
		panic(err)
	}

	err = ctx.SetOrder([]int{0, 1, 2, 3, 5, 4, 6})
	if err != nil {
		panic(err)
	}

	hs, err := ctx.NewHighlightStyle(chessImager.HighlightX, "#88E57C", 5, 0.95)
	if err != nil {
		panic(err)
	}

	as, _ := ctx.NewAnnotationStyle(
		chessImager.PositionTopLeft,
		25, 20, 1,
		"E8E57C", "000000", "FFFFFF",
	)

	ms, _ := ctx.NewMoveStyle(chessImager.MoveTypeDots, "#9D6B5EFF", 0.3)

	ctx.AddHighlightEx("e7", hs).AddAnnotationEx("e7", "!!", as).AddMoveEx("e1", "e7", ms)

	imgAdv, err := imager.RenderEx(fen, ctx)
	if err != nil {
		panic(err)
	}

	fileAdv, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fileAdv.Close()
	err = png.Encode(fileAdv, imgAdv)
	if err != nil {
		panic(err)
	}

}

func renderKasparov(fileName string) {
	imager := chessImager.NewImager()
	ctx, _ := imager.NewContext()

	// Highlight yellow square and
	// Annotate square e7 with "!!" and
	// Show move e1-e7
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("e1", "e7")

	// Render the image and save it to a file
	img, err := imager.RenderEx(fen, ctx)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create(fileName)
	defer file.Close()
	_ = png.Encode(file, img)
}

func renderSimple(fileName string) {
	imager := chessImager.NewImager()

	// Simple call
	imgSimple, err := imager.Render(fen)
	if err != nil {
		panic(err)
	}
	fileSimple, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fileSimple.Close()
	err = png.Encode(fileSimple, imgSimple)
	if err != nil {
		panic(err)
	}
}
