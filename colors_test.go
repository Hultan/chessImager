package chessImager

import (
	"image/color"
	"reflect"
	"testing"
)

func Test_hexToRGBA(t *testing.T) {
	t.Parallel()

	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantCol color.RGBA
		wantErr bool
	}{
		{name: "#888888FF", args: args{"#888888FF"}, wantCol: color.RGBA{R: 136, G: 136, B: 136, A: 255},
			wantErr: false},
		{name: "#888888GG", args: args{"#888888GG"}, wantCol: color.RGBA{R: 136, G: 136, B: 136, A: 0}, wantErr: true},
		{name: "#88888GG", args: args{"#88888GG"}, wantCol: color.RGBA{R: 0, G: 0, B: 0, A: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCol, err := hexToRGBA(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToRGBA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCol, tt.wantCol) {
				t.Errorf("hexToRGBA() gotCol = %v, want %v", gotCol, tt.wantCol)
			}
		})
	}
}

func TestColors(t *testing.T) {
	t.Parallel()

	// Create a new image context
	const fen = "b2r3r/k3Rp1p/p2q1np1/Np1P4/3p1Q2/P4PPB/1PP4P/1K6 b - - 1 25"
	ctx := NewImager().NewContext(fen)

	// Test color #RRGGBBAA
	s, err := ctx.NewMoveStyle(MoveTypeDots, "#9D6B5EFF", "9D6B5E", 0.2, 0)
	if err != nil {
		t.Errorf("failed to create style with color #RRGGBBAA: %v", err)
	}
	if !validateColor(s) {
		t.Errorf("failed to validate style with color #RRGGBBAA: %v", err)
	}

	// Test color #RRGGBB
	s, err = ctx.NewMoveStyle(MoveTypeDots, "#9D6B5E", "9D6B5E", 0.2, 0)
	if err != nil {
		t.Errorf("failed to create style with color #RRGGBB: %v", err)
	}
	if !validateColor(s) {
		t.Errorf("failed to validate style with color #RRGGBB: %v", err)
	}

	// Test color RRGGBBAA
	s, err = ctx.NewMoveStyle(MoveTypeDots, "9D6B5EFF", "9D6B5E", 0.2, 0)
	if err != nil {
		t.Errorf("failed to create style with color RRGGBBAA: %v", err)
	}
	if !validateColor(s) {
		t.Errorf("failed to validate style with color RRGGBBAA: %v", err)
	}

	// Test color RRGGBB
	s, err = ctx.NewMoveStyle(MoveTypeDots, "9D6B5E", "9D6B5E", 0.2, 0)
	if err != nil {
		t.Errorf("failed to create style with color RRGGBB: %v", err)
	}
	if !validateColor(s) {
		t.Errorf("failed to validate style with color RRGGBB: %v", err)
	}
}

// Make sure that the correct color is applied to the style
func validateColor(s *MoveStyle) bool {
	if s.Color.R == 157 && s.Color.G == 107 && s.Color.B == 94 && s.Color.A == 255 {
		return true
	}
	return false
}

func TestColorRGBA_MarshalJSON(t *testing.T) {
	c := ColorRGBA{RGBA: color.RGBA{R: 16, G: 32, B: 64, A: 128}}

	b, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("failed to marshal color: %v", err)
	}

	valid := []byte{34, 35, 49, 48, 50, 48, 52, 48, 56, 48, 34}
	for i, v := range valid {
		if b[i] != v {
			t.Errorf("failed to marshal color: %v != %v", b[i], v)
		}
	}
}

func TestColorRGBA_UnMarshalJSON(t *testing.T) {
	c := ColorRGBA{RGBA: color.RGBA{R: 16, G: 32, B: 64, A: 128}}
	cBytes := []byte{34, 35, 49, 48, 50, 48, 52, 48, 56, 48, 34}

	err := c.UnmarshalJSON(cBytes)
	if err != nil {
		t.Errorf("failed to unmarshal color: %v", err)
	}

	if c.R != 16 {
		t.Errorf("failed to unmarshal color Red (%v!=16)", c.R)
	}
	if c.G != 32 {
		t.Errorf("failed to unmarshal color Green (%v!=32)", c.G)
	}
	if c.B != 64 {
		t.Errorf("failed to unmarshal color Blue (%v!=64)", c.B)
	}
	if c.A != 128 {
		t.Errorf("failed to unmarshal color Alpha (%v!=128)", c.A)
	}
}

func TestColorRGBA_UnMarshalInvalidJSON(t *testing.T) {
	c := ColorRGBA{}
	// Invalid byte array
	cBytes := []byte{34, 35, 49, 48, 50, 48, 52, 48, 56, 48, 34, 56, 32}

	err := c.UnmarshalJSON(cBytes)
	if err == nil {
		t.Errorf("succeeded to unmarshal color: %v", err)
	}
}

func TestColorRGBA_ToRGBA(t *testing.T) {
	c := ColorRGBA{RGBA: color.RGBA{R: 16, G: 32, B: 64, A: 128}}

	r, g, b, a := c.toRGBA()

	if r != 0.06274509803921569 {
		t.Errorf("failed to unmarshal color Red (%v!=0.06274509803921569)", r)
	}
	if g != 0.12549019607843137 {
		t.Errorf("failed to unmarshal color Green (%v!=0.12549019607843137)", g)
	}
	if b != 0.25098039215686274 {
		t.Errorf("failed to unmarshal color Blue (%v!=0.25098039215686274)", b)
	}
	if a != 0.5019607843137255 {
		t.Errorf("failed to unmarshal color Alpha (%v!=0.5019607843137255)", a)
	}
}
