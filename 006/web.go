package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

//Temps is template structure
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

//Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

//setup template function
func setupTemp() *Temps {
	temps := new(Temps)

	temps.notemp = notemp()

	//set index template
	indx, er := template.ParseFiles("templates/index.html")
	if er != nil {
		indx = temps.notemp
	}
	temps.indx = indx

	//set hello template.
	helo, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo

	return temps
}

//index handler
func index(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

//hello handler.
func hello(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {

	var flg bool = false

	rand.Seed(time.Now().Unix())

	switch rand.Intn(2) {
	case 1:
		flg = true
	case 0:
		flg = false
	}

	ses, _ := cs.Get(rq, "hello-session")

	nm := ""
	pw := ""
	if rq.Method == "POST" {
		ses.Values["login"] = nil
		ses.Values["name"] = nil
		nm := rq.PostFormValue("name")
		pw := rq.PostFormValue("pass")
		if nm == pw {
			ses.Values["login"] = true
			ses.Values["name"] = nm
		}
		ses.Save(rq, w)
	}

	isLogin, _ := ses.Values["login"].(bool)
	lname, _ := ses.Values["name"].(string)

	item := struct {
		Flg      bool
		Title    string
		Message  string
		JMessage string
		Items    []string
		ParamId  string
		PostName string
		PostPass string
		IsLogin  bool
		Lname    string
	}{
		Flg:      flg,
		Title:    "Send values",
		Message:  "Content01",
		JMessage: "Content02",
		Items:    []string{"One", "Two", "Three"},
		ParamId:  rq.FormValue("id"),
		PostName: nm,
		PostPass: pw,
		IsLogin:  isLogin,
		Lname:    lname,
	}

	er := tmp.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {

	temps := setupTemp()

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq, temps.indx)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq, temps.helo)
	})

	http.ListenAndServe("127.0.0.1:8050", nil)
}
