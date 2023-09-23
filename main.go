package main

import (
	"image/png"
	"os"

	"chessImager/chessImager"
)

const (
	fen = "3Rr3/8/b1kp1p1p/1q5p/1P2P3/P4K2/6Qb/6N1 w - - 0 1"
)

func main() {
	imager, err := chessImager.NewImager()
	if err != nil {
		panic(err)
	}

	// Advanced call
	settings, err := chessImager.GetSettings()
	if err != nil {
		panic(err)
	}
	settings.Highlight = append(settings.Highlight, chessImager.HighlightedSquare{
		Square: "F3",
		Color:  "#8844ff80",
		Type:   0,
	})
	img2 := imager.GetImageEx(fen, settings)
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
	img := imager.GetImage(fen)
	f, err := os.Create("/home/per/temp/img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
