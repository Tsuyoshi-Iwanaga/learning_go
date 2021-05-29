package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("127.0.0.1:8050", http.FileServer(http.Dir(".")))
}
