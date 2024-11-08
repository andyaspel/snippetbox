package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/nav-bar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
	}
	fmt.Fprintf(w, "\n\tDisplay a specific snippet with Id: %d...\n", id)

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	w.Write([]byte("\n\tCreate a new snippet...\n"))

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/about.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/nav-bar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/contact.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/nav-bar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
