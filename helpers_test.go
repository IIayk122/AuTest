package main

import (
	"reflect"
	"testing"
)

func Test_checkGuess(t *testing.T) {
	type args struct {
		num1   []int
		secret []int
	}
	tests := []struct {
		name string
		args args
		want set
	}{

		{"ololo",
			args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 0, 0, 0}},
			set{[]int{9, 0, 0, 0}, 4, 0}},
		{"ololo", args{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}}, set{[]int{0, 0, 0, 0}, 4, 0}},
		{"ololo", args{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}}, set{[]int{0, 0, 0, 0}, 4, 0}},
	}
	for _, tt := range tests {
		if got := checkGuess(tt.args.num1, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. checkGuess() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_generateSecret(t *testing.T) {
	tests := []struct {
		name          string
		wantGenerated []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if gotGenerated := generateSecret(); !reflect.DeepEqual(gotGenerated, tt.wantGenerated) {
			t.Errorf("%q. generateSecret() = %v, want %v", tt.name, gotGenerated, tt.wantGenerated)
		}
	}
}
