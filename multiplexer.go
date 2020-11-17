package main

import (
	"io"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

func wel(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to home page!")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/", wel)

	http.ListenAndServe(":5050", mux)
}
