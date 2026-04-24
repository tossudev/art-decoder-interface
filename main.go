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
	
	parsedInput := decoder.Decode(input)
	// do somthing with the input here

	data := Page{
		Input: 	input,
		Output: parsedInput,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ERR:", err)
	}

	tmpl.Execute(w, data)
}
