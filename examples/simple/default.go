package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func main() {
	//runtime.GOMAXPROCS(8)
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:3001", nil)
}
