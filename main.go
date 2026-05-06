package main

import (
	"interface/decoder"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"html/template"
)

type Page struct {
	Input string
	Output string
	HttpResponseText string
	ErrorMessage string
	InfoMessage string
}

const (
	StatusPrefix string = "Status code: "
	Port string = "6969"
)

var pageData Page


func main() {
	http.HandleFunc("POST /decoder", FormHandler)
	http.HandleFunc("/", FrontendHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	fmt.Println(fmt.Sprintf("Started server at: localhost:%s", Port))
	http.ListenAndServe(":" + Port, nil)
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
	timeStart := time.Now()

	switch inputType {
		case "Encode":
			var shortenedPercentage float32
			
			output, shortenedPercentage = decoder.Encode(input)
			
			timeElapsed := time.Since(timeStart)
			fmt.Println("Encoding took:", timeElapsed)

			if shortenedPercentage > 0.0 {
				infoMessage = fmt.Sprintf("Encoded string is %.1f%% shorter!", shortenedPercentage)
			}
			responseCode = 202

		case "Decode":
			var err error
			err, output = decoder.Decode(input)
			
			timeElapsed := time.Since(timeStart)
			fmt.Println("Decoding took:", timeElapsed)
			
			responseCode = 202

			if err != nil {
				output = ""
				responseCode = 400
				errorMessage = "Malformed input!"
			}
	}

	pageData = Page{
		Input: 				input,
		Output: 			output,
		HttpResponseText:	StatusPrefix + strconv.Itoa(responseCode),
		ErrorMessage:		errorMessage,
		InfoMessage:		infoMessage,
	}

	w.WriteHeader(responseCode)
	
	// force update site, this is horseshit but I couldn't figure out how to do it without redirect
	FrontendHandler(w, r)
}


func FrontendHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path != "/" && path != "/decoder" {
		w.WriteHeader(404)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ERR:", err)
	}

	tmpl.Execute(w, pageData)
}

