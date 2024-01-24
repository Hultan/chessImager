package chessImager

import (
	"reflect"
	"testing"
)

func Test_newAlg(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    alg
		wantErr bool
	}{
		// Valid moves
		{"0-0", args{"0-0"}, alg{pos: "0-0", status: moveStatusKingSideCastling}, false},
		{"O-O", args{"O-O"}, alg{pos: "o-o", status: moveStatusKingSideCastling}, false},
		{"o-o", args{"o-o"}, alg{pos: "o-o", status: moveStatusKingSideCastling}, false},
		{"0-0-0", args{"0-0-0"}, alg{pos: "0-0-0", status: moveStatusQueenSideCastling}, false},
		{"O-O-O", args{"O-O-O"}, alg{pos: "o-o-o", status: moveStatusQueenSideCastling}, false},
		{"o-o-o", args{"o-o-o"}, alg{pos: "o-o-o", status: moveStatusQueenSideCastling}, false},
		{"", args{""}, alg{pos: "", status: moveStatusEmpty}, false},
		{"A1", args{"A1"}, alg{pos: "a1", x: 0, y: 0, status: moveStatusNormal}, false},
		{"H8", args{"H8"}, alg{pos: "h8", x: 7, y: 7, status: moveStatusNormal}, false},
		{"A8", args{"A8"}, alg{pos: "a8", x: 0, y: 7, status: moveStatusNormal}, false},
		{"H1", args{"H1"}, alg{pos: "h1", x: 7, y: 0, status: moveStatusNormal}, false},
		{"a1", args{"a1"}, alg{pos: "a1", x: 0, y: 0, status: moveStatusNormal}, false},
		{"h8", args{"h8"}, alg{pos: "h8", x: 7, y: 7, status: moveStatusNormal}, false},
		{"a8", args{"a8"}, alg{pos: "a8", x: 0, y: 7, status: moveStatusNormal}, false},
		{"h1", args{"h1"}, alg{pos: "h1", x: 7, y: 0, status: moveStatusNormal}, false},

		// Invalid moves
		{"h11", args{"h11"}, alg{pos: "h11", status: moveStatusIllegal}, true},
		{"a9", args{"a9"}, alg{pos: "a9", status: moveStatusIllegal}, true},
		{"a0", args{"a0"}, alg{pos: "a0", status: moveStatusIllegal}, true},
		{"i1", args{"i1"}, alg{pos: "i1", status: moveStatusIllegal}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newAlg(tt.args.s, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("newAlg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAlg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAlgCoord(t *testing.T) {
	t.Parallel()

	type args struct {
		s        string
		inverted bool
	}
	tests := []struct {
		name         string
		args         args
		want1, want2 int
	}{
		{"A1", args{"A1", false}, 0, 0},
		{"H8", args{"H8", false}, 7, 7},
		{"A8", args{"A8", false}, 0, 7},
		{"H1", args{"H1", false}, 7, 0},
		{"a1", args{"a1", false}, 0, 0},
		{"h8", args{"h8", false}, 7, 7},
		{"a8", args{"a8", false}, 0, 7},
		{"h1", args{"h1", false}, 7, 0},

		{"A1", args{"A1", true}, 7, 7},
		{"H8", args{"H8", true}, 0, 0},
		{"A8", args{"A8", true}, 7, 0},
		{"H1", args{"H1", true}, 0, 7},
		{"a1", args{"a1", true}, 7, 7},
		{"h8", args{"h8", true}, 0, 0},
		{"a8", args{"a8", true}, 7, 0},
		{"h1", args{"h1", true}, 0, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := newAlg(tt.args.s, tt.args.inverted)
			got1, got2 := got.coords()
			if got1 != tt.want1 {
				t.Errorf("got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("got1 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func Test_newAlgCoordPanic(t *testing.T) {
	t.Parallel()

	type args struct {
		s        string
		inverted bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{"", false}},
		{"0-0", args{"0-0", false}},
		{"0-0-0", args{"0-0-0", false}},
		{"o-o", args{"o-o", false}},
		{"o-o-o", args{"0-0-0", false}},
		{"i9", args{"i9", false}},
		{"a", args{"a", false}},
		{"", args{"", true}},
		{"0-0", args{"0-0", true}},
		{"0-0-0", args{"0-0-0", true}},
		{"o-o", args{"o-o", true}},
		{"o-o-o", args{"0-0-0", true}},
		{"i9", args{"i9", true}},
		{"a", args{"a", true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			got, _ := newAlg(tt.args.s, tt.args.inverted)
			_, _ = got.coords()
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	a, _ := newAlg("e4", false)
	got := a.String()
	if got != "move: e4" {
		t.Errorf("String(e4) failed: %v", got)
	}
}
