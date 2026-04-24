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


func main() {
	http.HandleFunc("/", FormHandler)
	http.ListenAndServe(":8080", nil)
}


func FormHandler(w http.ResponseWriter, r *http.Request) {
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

	data := Page{
		Input: 	input,
		Output: output,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ERR:", err)
	}

	tmpl.Execute(w, data)
}
