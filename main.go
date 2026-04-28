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
	HttpResponse string
}

var pageData Page


func main() {
	http.HandleFunc("/decoder", FormHandler)
	http.HandleFunc("/", FrontendHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":6969", nil)
}


func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	input := r.FormValue("input")
	inputType := r.FormValue("action")
	output := ""
	var err error

	switch inputType {
		case "Encode":
			output = decoder.Encode(input)
		case "Decode":
			err, output = decoder.Decode(input)
	}

	if err != nil {
		output = ""
		w.WriteHeader(400)
	}

	pageData = Page{
		Input: 			input,
		Output: 		output,
		HttpResponse:	"Response status",
		//HttpResponse:	r.Response.Status,
	}

	res, err := http.Get("/")

	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
}


func FrontendHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ERR:", err)
	}

	tmpl.Execute(w, pageData)
}
