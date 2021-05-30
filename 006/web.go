package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	html := `<html>
		<body>
		<h1>Hello</h1>
		<p>This is sample message.</p>
		</body></html>`
	tf, er := template.New("index").Parse(html)
	if er != nil {
		log.Fatal(er)
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
