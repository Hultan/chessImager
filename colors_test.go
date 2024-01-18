package chessImager

import (
	"testing"
)

func TestColors(t *testing.T) {
	// Create a new image context
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	ctx := NewImager().NewContext(fen)

	// Test color #RRGGBBAA
	s, err := ctx.NewMoveStyle(MoveTypeDots, "#9D6B5EFF", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color #RRGGBBAA: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color #RRGGBBAA: %v", err)
	}

	// Test color #RRGGBB
	s, err = ctx.NewMoveStyle(MoveTypeDots, "#9D6B5E", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color #RRGGBB: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color #RRGGBB: %v", err)
	}

	// Test color RRGGBBAA
	s, err = ctx.NewMoveStyle(MoveTypeDots, "9D6B5EFF", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color RRGGBBAA: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color RRGGBBAA: %v", err)
	}

	// Test color RRGGBB
	s, err = ctx.NewMoveStyle(MoveTypeDots, "9D6B5E", "9D6B5E", 0.2)
	if err != nil {
		t.Fatalf("failed to create style with color RRGGBB: %v", err)
	}
	if !validateColor(s) {
		t.Fatalf("failed to validate style with color RRGGBB: %v", err)
	}
}

// Make sure that the correct color is applied to the style
func validateColor(s *MoveStyle) bool {
	if s.Color.R == 157 && s.Color.G == 107 && s.Color.B == 94 && s.Color.A == 255 {
		return true
	}
	return false
}
