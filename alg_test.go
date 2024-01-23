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
