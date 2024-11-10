package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"gorm.io/gorm"
)

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
	fmt.Println("\nSHOW:\t", id)
	if err != nil || id < 1 {
		app.notFound(w)
		fmt.Println("SNIPPET NOT FOUND 1")
		return
	}
	fmt.Println("Id LEGAL")
	s, err := app.snippets.Get(id)
	// err := models.ErrorRecord
	// fmt.Println(err)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("out", s)
		// app.notFound(w)
		return
	}

	// fmt.Println("RETURN SNIPPET FOUND", s.ID, id, s.Title)

	fmt.Fprintf(w, "\t%s\n\t%s\n\t%s\n", s.Title, s.Content, s.Expires)

	// fmt.Fprintf(w, "\n\tDisplay a specific snippet with Id: %d...\n\t %v", id, s)

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "wwwwwwooooooooo"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\nKobayashi"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// app.snippets.Get(id)

	fmt.Println("Title: ", title, id)
	// app.snippets.Get(id)
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
