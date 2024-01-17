package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/Hultan/chessImager"
	"gopkg.in/freeeve/pgn.v1"
)

func main() {
	imager := chessImager.NewImager()

	f, err := os.Open("./game.pgn")
	if err != nil {
		log.Fatal(err)
	}
	ps := pgn.NewPGNScanner(f)

	for ps.Next() {
		game, err := ps.Scan()
		if err != nil {
			log.Fatal(err)
		}

		b := pgn.NewBoard()
		i := 1
		for _, move := range game.Moves {
			_ = b.MakeMove(move)

			ctx := imager.NewContext()
			ctx.AddMove(move.From.String(), move.To.String()).AddHighlight(move.From.String()).AddHighlight(move.To.String())
			_ = ctx.SetFEN(b.String())

			img, _ := imager.RenderEx(ctx)

			file, _ := os.Create(fmt.Sprintf("%d.png", i))
			_ = png.Encode(file, img)
			_ = file.Close()
			i++
		}
	}
}
