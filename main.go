package main

import (
	"net/http"
	//"io"
)


func main() {
	/*
	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "fuck you\n")
	}
	*/

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

}
