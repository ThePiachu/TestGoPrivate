package main

import (
	appengine "google.golang.org/appengine/v2"
	"html/template"
	"net/http"

	"github.com/ThePiachu/TestGoPrivate/Work"
	"github.com/ThePiachu/Go/Log"
)

var MainTemplate *template.Template

func init() {
	http.HandleFunc("/", hello)

	//Admin functions
	http.HandleFunc("/update", update)

	MainTemplate, _ = template.ParseFiles("html/main.html")
}

func main() {
	appengine.Main()
}

func hello(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	Work.Work(c)
	var err error
	MainTemplate, err = template.ParseFiles("html/main.html")
	if (err!=nil) {
		Log.Errorf(c, "%v", err)
	}
	MainTemplate.Execute(w, nil)
}

func update(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	Work.Work(c)
	MainTemplate.Execute(w, nil)
}
