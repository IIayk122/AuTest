package main

import (
	"testing"
)

func TestNumberss_checkGuess(t *testing.T) {
	cases := []struct {
		guess    int
		secret   numbers
		expected set
	}{
		{1234, numbers{5, 6, 7, 8}, set{1234, 0, 0}},
		{1256, numbers{9, 7, 2, 1}, set{1256, 0, 0}},
		{1256, numbers{1, 2, 0, 0}, set{1256, 2, 0}},
		{0000, numbers{0, 0, 0, 0}, set{0000, 4, 0}},
		{1234, numbers{4, 3, 2, 1}, set{1234, 0, 4}},
		{5612, numbers{5, 1, 6, 2}, set{5612, 2, 2}},
		{0000, numbers{0, 1, 1, 1}, set{0000, 1, 3}},
	}

	for _, tc := range cases {
		v := checkGuess(tc.guess, tc.secret)
		if v != tc.expected {
			t.Error(
				"\n Для", tc.guess,
				"и", tc.secret,
				"\n хотел    B:", tc.expected.B, "K:", tc.expected.K,
				"\n получчил B:", v.B, "K:", v.K,
			)
		}
	}
}
