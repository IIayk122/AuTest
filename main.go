package main

import (
	"net/http"
	"text/template"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/play/", play)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
