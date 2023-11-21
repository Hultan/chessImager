package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestColors(t *testing.T) {
	// Create a new context
	ctx := chessImager.NewImager().NewContext()

	// Test color #RRGGBBAA
	_, err := ctx.NewMoveStyle(chessImager.MoveTypeDots, "#9D6B5EFF", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color #RRGGBBAA: %v", err)
	}

	// Test color #RRGGBB
	_, err = ctx.NewMoveStyle(chessImager.MoveTypeDots, "#9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color #RRGGBB: %v", err)
	}

	// Test color RRGGBBAA
	_, err = ctx.NewMoveStyle(chessImager.MoveTypeDots, "9D6B5EFF", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color RRGGBBAA: %v", err)
	}

	// Test color RRGGBB
	_, err = ctx.NewMoveStyle(chessImager.MoveTypeDots, "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color RRGGBB: %v", err)
	}
}
