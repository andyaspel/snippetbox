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

	for _, s := range s {
		if s.ID == 0 {
			break
		}
		fmt.Fprintf(w, "Last 10 records:\tID:%v\nTitle:\n\t\t%s\nContent:\n\t\t%s\nExpires in:\n\t\t%s\n", s.ID, s.Title, s.Content, s.Expires)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		fmt.Println("NOT A VALID ID")
		return
	}
	s, err := app.snippets.Get(id) // Get the record
	if err != nil {
		app.notFound(w)
		fmt.Println("RECORD NOT FOUND with ID - ", id)
		return
	}
	data := &templateData{Snippet: s}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/nav-bar.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "hmrn"
	content :=
		`func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	snippet := &models.Snippet{Title: title, Content: content, Expires: expires}
	s.DB.Create(&snippet) // pass a slice to insert multiple row\n\ts.DB.Save(&snippet)
	fmt.Printf("\nNew record created on %v\nID: %d\t\tTitle: %s\nContent: %s\nExpires: %s\n", snippet.CreatedAt, snippet.ID, title, content, expires)
	return int(snippet.ID), nil
}
`

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
