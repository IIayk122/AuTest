package main

import "text/template"

type set struct {
	Guess []int
	B, K  int
}

type game struct {
	secretNumber []int
	Sets         []set
	Win          bool
}

var tpl *template.Template
var data = map[string]game{}
