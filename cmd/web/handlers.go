package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/andyaspel/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest()
	if err == models.ErrorRecords {
		app.serverError(w, err)
		return
	}
	for _, snippet := range s {
		if snippet.ID == 0 {
			break
		}
		// fmt.Fprintf(w, "%v\n", *snippet)
		fmt.Fprintf(w, "Last 10 records:\tID:%v\nTitle:\n\t\t%s\nContent:\n\t\t%s\nExpires in:\n\t\t%s\n", snippet.ID, snippet.Title, snippet.Content, snippet.Expires)

	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		fmt.Println("NOT A VALID ID")
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil && s.Title == "" {
		app.notFound(w)
		fmt.Println("RECORD NOT FOUND with ID - ", id)
		return
	}
	fmt.Fprintf(w, "Found Record:\t%v\nTitle:\n\t\t%s\nContent:\n\t\t%s\nExpires in:\n\t\t%s\n", s.ID, s.Title, s.Content, s.Expires)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "hmrn"
	content := "O snail\n\tClimb Mount Fuji,\n\tBut slowly, slowly!\n\tKobayashi"
	expires := "8"
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

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
