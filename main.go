package main

import (
	"image/png"
	"os"

	"chessImager/chessImager"
)

func main() {
	imager, err := chessImager.NewImager()
	if err != nil {
		panic(err)
	}

	// Simple call
	img := imager.GetImage("6Qb/K7/P1P5/5R2/2P1N3/P1pp3n/1P5P/1B2k3 w - - 0 1")
	f, err := os.Create("/home/per/temp/img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
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
		Color:  "#8844ffff",
		Type:   0,
	})
	img2 := imager.GetImageEx("6Qb/K7/P1P5/5R2/2P1N3/P1pp3n/1P5P/1B2k3 w - - 0 1", settings)
	f2, err := os.Create("/home/per/temp/img2.png")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	err = png.Encode(f2, img2)
	if err != nil {
		panic(err)
	}
}
