package main

import (
	"interface/decoder"
	"fmt"
	"strconv"
	"net/http"
	"html/template"
)

type Page struct {
	Input string
	Output string
	HttpResponse int
	HttpResponseText string
	ErrorMessage string
	InfoMessage string
}

const (
	StatusPrefix string = "Status code: "
)

var pageData Page


func main() {
	http.HandleFunc("POST /decoder", FormHandler)
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
	errorMessage := ""
	infoMessage := ""

	responseCode := 302

	switch inputType {
		case "Encode":
			var shortenedPercentage float32
			output, shortenedPercentage = decoder.Encode(input)
			infoMessage = fmt.Sprintf("Encoded string is %.1f%% shorter!", shortenedPercentage)
			responseCode = 202

		case "Decode":
			var err error
			err, output = decoder.Decode(input)
			
			if err != nil {
				output = ""
				responseCode = 400
				errorMessage = "Malformed input!"
			}
	}

	pageData = Page{
		Input: 				input,
		Output: 			output,
		HttpResponse:		responseCode,
		HttpResponseText:	StatusPrefix + strconv.Itoa(responseCode),
		ErrorMessage:		errorMessage,
		InfoMessage:		infoMessage,
	}

	w.WriteHeader(responseCode)
	
	// force update site, this is horseshit but I couldn't figure out how to do it without redirect
	tmpl, err2 := template.ParseFiles("index.html")
	if err2 != nil {
		fmt.Println("ERR:", err2)
	}

	tmpl.Execute(w, pageData)
}


func FrontendHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ERR:", err)
	}

	tmpl.Execute(w, pageData)
}
