package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tf, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		tf, _ = template.New("index").Parse("<html><body><h1>NO TEMPLATE.</h1></body></html>")
	}

	hh := func(w http.ResponseWriter, rq *http.Request) {
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}

	http.HandleFunc("/hello", hh)

	http.ListenAndServe("127.0.0.1:8050", nil)
}
