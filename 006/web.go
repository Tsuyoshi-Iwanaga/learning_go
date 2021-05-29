package main

import (
	"net/http"
)

func main() {
	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Hello, This is GO-server!!"))
	}
	http.HandleFunc("/hello", hh)
	http.ListenAndServe("127.0.0.1:8050", nil)
	// http.ListenAndServe("127.0.0.1:8050", http.FileServer(http.Dir(".")))
}
