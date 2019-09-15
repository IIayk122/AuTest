package main

import "text/template"

type numbers [4]int

type set struct {
	Guess, B, K int
}

type game struct {
	secretNumber numbers
	Sets         []set
	Win          bool
}

var tpl *template.Template
var data = map[string]game{}
