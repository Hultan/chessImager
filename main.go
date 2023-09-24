package main

import (
	"image/color"
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
	s, err := chessImager.GetSettings()
	if err != nil {
		panic(err)
	}
	s.AddHighlight("g3")
	styleHigh := &chessImager.HighlightedSquareStyle{
		Color: chessImager.ColorRGBA{RGBA: color.RGBA{R: 0, G: 255, A: 80}},
		Type:  0,
		Width: 3,
	}
	s.AddHighlightEx("F6", styleHigh)
	s.AddAnnotation("c4", "#")
	styleAnn := &chessImager.AnnotationStyle{
		Position:        3,
		Size:            16,
		FontSize:        12,
		BackgroundColor: chessImager.ColorRGBA{RGBA: color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		ForegroundColor: chessImager.ColorRGBA{RGBA: color.RGBA{R: 0, G: 0, B: 0, A: 255}},
		BorderColor:     chessImager.ColorRGBA{RGBA: color.RGBA{R: 0, G: 0, B: 0, A: 255}},
		BorderWidth:     1,
	}
	s.AddAnnotationEx("a1", "!!", styleAnn)
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
