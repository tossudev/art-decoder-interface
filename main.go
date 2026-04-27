package main

import (
	"interface/decoder"
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Input string
	Output string
}

var pageData Page


func main() {
	http.HandleFunc("/decoder", FormHandler)
	http.HandleFunc("/", FrontendHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}


func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	input := r.FormValue("input")
	inputType := r.FormValue("action")
	output := ""

	switch inputType {
		case "Encode":
			output = decoder.Encode(input)
		case "Decode":
			output = decoder.Decode(input)
	}

	pageData = Page{
		Input: 	input,
		Output: output,
	}

	http.Redirect(w, r, "/", 302) // yea idk man
}


func FrontendHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ERR:", err)
	}

	tmpl.Execute(w, pageData)
}
