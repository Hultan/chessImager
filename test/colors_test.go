package test

import (
	"testing"

	"github.com/Hultan/chessImager"
)

func TestColors(t *testing.T) {
	// Create a new context
	ctx := chessImager.NewImager().NewContext()

	// Test color #RRGGBBAA
	s, err := ctx.NewMoveStyle(chessImager.MoveTypeDots, "#9D6B5EFF", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color #RRGGBBAA: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color #RRGGBBAA: %v", err)
	}

	// Test color #RRGGBB
	s, err = ctx.NewMoveStyle(chessImager.MoveTypeDots, "#9D6B5E", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color #RRGGBB: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color #RRGGBB: %v", err)
	}

	// Test color RRGGBBAA
	s, err = ctx.NewMoveStyle(chessImager.MoveTypeDots, "9D6B5EFF", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color RRGGBBAA: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color RRGGBBAA: %v", err)
	}

	// Test color RRGGBB
	s, err = ctx.NewMoveStyle(chessImager.MoveTypeDots, "9D6B5E", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color RRGGBB: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color RRGGBB: %v", err)
	}
}

// Make sure that the correct color is applied to the style
func validateColor(s *chessImager.MoveStyle) bool {
	if s.Color.R == 157 && s.Color.G == 107 && s.Color.B == 94 && s.Color.A == 255 {
		return true
	}
	return false
}
