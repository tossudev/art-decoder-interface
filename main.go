package main

import (
	"net/http"
	"fmt"
	//"io"
)


func main() {
	/*
	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "fuck you\n")
	}
	*/

	http.HandleFunc("/input", FormHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

}


func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	input := r.FormValue("input")

	fmt.Println(input)
}
