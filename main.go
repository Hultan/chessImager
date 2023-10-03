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

	// Advanced call
	s, err := chessImager.LoadSettings("")
	if err != nil {
		panic(err)
	}

	// Highlight square
	hs := chessImager.NewHighlightStyle(0, "88008888", 0)
	s.AddHighlightEx("e7", hs)

	// Annotate square
	as := chessImager.NewAnnotationStyle(
		chessImager.PositionTopRight,
		20, 15, 1,
		"BBBBBBFF", "000000FF", "000000FF",
	)
	s.AddAnnotationEx("e7", "!!", as)

	// Add move
	ms := chessImager.NewMoveStyle(chessImager.MoveTypeDots, "80008080", 0.3)
	s.AddMoveEx("e1", "e7", ms)

	img2 := imager.GetImageEx(fen, s)
	f2, err := os.Create("/home/per/temp/img2.png")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	err = png.Encode(f2, img2)
	if err != nil {
		panic(err)
	}

	// Simple call
	//img := imager.GetImage(fen)
	//f, err := os.Create("/home/per/temp/example.png")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//err = png.Encode(f, img)
	//if err != nil {
	//	panic(err)
	//}
}
