package main

import (
	"testing"
)

func TestNumberss_checkGuess(t *testing.T) {
	cases := []struct {
		left     int
		right    numbers
		expected set
	}{
		{1234, numbers{1, 2, 3, 4}, set{1234, 4, 0}},
		{1234, numbers{5, 6, 7, 8}, set{1234, 0, 0}},
		{5612, numbers{5, 1, 6, 2}, set{5612, 2, 2}},
		{1256, numbers{9, 7, 2, 1}, set{1256, 0, 2}},
	}

	for _, tc := range cases {
		v := checkGuess(tc.left, tc.right)
		if v != tc.expected {
			t.Error(
				"For", tc.left,
				"expected", tc.right,
				"got", v,
			)
		}
	}
}
