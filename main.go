package main

import (
	"chessImager/chessImager"
)

func main() {
	imager, err := chessImager.NewImager()
	if err != nil {
		panic(err)
	}
	settings := chessImager.ImageSettings{
		Highlight: []chessImager.HighlightedSquare{
			{
				Square: "c4",
				Color:  "#00606480",
				Width:  3,
				Type:   1,
			},
			{
				Square: "b1",
				Color:  "#40f0ff80",
				Width:  0,
				Type:   0,
			},
			{
				Square: "b7",
				Color:  "#55FF4580",
				Width:  0,
				Type:   0,
			},
		},
	}
	_ = imager.GetImage(settings)
}
