package chessImager

import (
	"os"
	"testing"

	"gopkg.in/freeeve/pgn.v1"
)

func TestSimpleExample(t *testing.T) {
	t.Parallel()

	filename := "simple.png"

	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	img, err := NewImager().Render(fen)
	if err != nil {
		t.Errorf("failed to render : %v", err)
	}
	if img == nil {
		t.Errorf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestMediumExample(t *testing.T) {
	t.Parallel()

	filename := "medium.png"

	// Create a new imager using embedded default.json settings
	imager := NewImager()

	// Create a new image context
	ctx := imager.NewContext("b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25")

	// Highlight square e7
	// Annotate square e7 with "!!"
	// Show move e1-e7
	ctx.AddHighlight("e7").AddAnnotation("e7", "!!").AddMove("e1", "e7")

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		t.Errorf("failed to render : %v", err)
	}
	if img == nil {
		t.Errorf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestAdvancedExample(t *testing.T) {
	t.Parallel()

	filename := "advanced.png"

	// Create a new imager using embedded default.json settings
	imager := NewImager()

	// Set the rendering order
	err := imager.SetOrder([]int{0, 1, 2, 3, 5, 4, 6})
	if err != nil {
		t.Errorf("failed to set rendering order : %v", err)
	}

	// Create a new image context
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	ctx := imager.NewContext(fen)

	// Create a highlight style, for the square e7
	hs, err := ctx.NewHighlightStyle(
		HighlightTypeFilledCircle, // Highlight type
		"#88E57C",                 // Highlight color
		4,                         // Highlight circle width
		0.9,                       // Highlight factor (not used for this Type)
	)
	if err != nil {
		t.Errorf("failed to create highlight style : %v", err)
	}

	// Create an annotation style, for the square e7
	as, err := ctx.NewAnnotationStyle(
		PositionTypeTopLeft, // Position
		25, 20, 1,           // Size, font size, border width
		"#E8E57C", "#000000", "#FFFFFF", // Background, font, border color
	)
	if err != nil {
		t.Errorf("failed to create annotation style : %v", err)
	}

	// Create a move style, for the move e1-e7
	ms, err := ctx.NewMoveStyle(
		MoveTypeDots, // Move type
		"#9D6B5EFF",  // Dot color
		"#9D6B5EFF",  // Dot color 2
		0.2,          // Dot size
		0,            // Padding
	)
	if err != nil {
		t.Errorf("failed to create move style : %v", err)
	}

	// Highlight the e7 square, annotate e7 as a brilliant move (!!) and
	// show move e1-e7.
	ctx.AddHighlightWithStyle("e7", hs).AddAnnotationWithStyle("e7", "!!", as).AddMoveWithStyle("e1", "e7", ms)

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		t.Errorf("failed to render : %v", err)
	}
	if img == nil {
		t.Errorf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestOtherExample(t *testing.T) {
	t.Parallel()

	filename := "other.png"

	// Create a new imager using your custom JSON file
	imager, err := NewImagerFromPath("test/data/other.json")
	if err != nil {
		t.Errorf("failed to create imager : %v", err)
	}

	// Create a new image context and highlight the e7 square, annotate e7 as a
	// brilliant move (!!) and show move e1-e7.
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	ctx := imager.NewContext(fen).
		AddHighlight("e7").
		AddAnnotation("e7", "!!").
		AddMove("e1", "e7")

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		t.Errorf("failed to render : %v", err)
	}
	if img == nil {
		t.Errorf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestCastlingExample(t *testing.T) {
	t.Parallel()

	filename := "castling.png"

	// Create a new imager using embedded default.json settings
	imager := NewImager()

	// Create a new image context, and add white king side castling,
	// and black queen side castling.
	ctx := imager.NewContext("2kr4/8/8/8/8/8/8/5RK1 b - - 1 25").AddMove("0-0", "").AddMove("", "0-0-0")

	// Render the image
	img, err := imager.RenderWithContext(ctx)
	if err != nil {
		t.Errorf("failed to render : %v", err)
	}
	if img == nil {
		t.Errorf("image is nil")
	}

	compareImages(t, filename, &img)
}

func TestPGNExample(t *testing.T) {
	t.Parallel()

	imager := NewImager()

	f, err := os.Open("./test/data/game.pgn")
	if err != nil {
		t.Errorf("failed to open PGN file : %v", err)
	}
	ps := pgn.NewPGNScanner(f)

	for ps.Next() {
		game, err := ps.Scan()
		if err != nil {
			t.Errorf("failed to scan PGN file : %v", err)
		}

		b := pgn.NewBoard()
		i := 1
		for _, move := range game.Moves {
			// Let's just test the first 10 moves
			if i > 10 {
				continue
			}
			_ = b.MakeMove(move)

			ctx := imager.NewContext(b.String()).
				AddMove(move.From.String(), move.To.String()).
				AddHighlight(move.From.String()).
				AddHighlight(move.To.String())

			img, err := imager.RenderWithContext(ctx)
			if err != nil {
				t.Errorf("failed to render image %d : %v", i, err)
			}
			if img == nil {
				t.Errorf("image is nil")
			}

			i++
		}
	}
}
