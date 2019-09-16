package main

import (
	"math/rand"
	"time"
)

//
func checkGuess(num1 []int, secret []int) set {
	guess := num1
	B := 0
	K := 0
	in := num1

	if in[0] == secret[0] {
		B++
	} else if in[0] == secret[1] || in[0] == secret[2] || in[0] == secret[3] {
		K++
	}
	if in[1] == secret[1] {
		B++
	} else if in[1] == secret[0] || in[1] == secret[2] || in[1] == secret[3] {
		K++
	}
	if in[2] == secret[2] {
		B++
	} else if in[2] == secret[1] || in[2] == secret[0] || in[2] == secret[3] {
		K++
	}
	if in[3] == secret[3] {
		B++
	} else if in[3] == secret[1] || in[3] == secret[2] || in[3] == secret[0] {
		K++
	}
	return set{guess, B, K}
}

func generateSecret() (generated []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	generated = []int{
		rand.Intn(10),
		rand.Intn(10),
		rand.Intn(10),
		rand.Intn(10),
	}
	return
}
