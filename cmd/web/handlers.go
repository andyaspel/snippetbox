package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andyaspel/snippetbox/pkg/models"
)

// PAGES
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	l, err := app.lists.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	lg, err := app.logs.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", templateData{
		Snippets: s,
		Lists:    l,
		Logs:     lg,
	})
}
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}
	app.render(w, r, "about.page.tmpl", templateData{})
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact" {
		http.NotFound(w, r)
		return
	}
	app.render(w, r, "contact.page.tmpl", templateData{})

	// files := []string{
	// 	"./ui/html/contact.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// 	"./ui/html/nav-bar.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
}

// REPORTS
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
	app.render(w, r, "show.page.tmpl", templateData{
		Snippet: s,
	})
}
func (app *application) showList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		fmt.Println("NOT A VALID ID")
		return
	}
	l, err := app.lists.Get(id) // Get the record
	if err != nil {
		app.notFound(w)
		fmt.Println("RECORD NOT FOUND with ID - ", id)
		return
	}
	app.render(w, r, "show.page.tmpl", templateData{
		List: l,
	})
}

// FORMS
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "hmrn"
	content := `
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
func (app *application) createList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "todo"
	content := `finish an application to manage snippets and lists`
	done := false
	id, err := app.lists.Insert(title, content, done)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the list.
	http.Redirect(w, r, fmt.Sprintf("/list?id=%d", id), http.StatusSeeOther)

}

// LOGS
func (app *application) showLogs(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		fmt.Println("NOT A VALID ID")
		return
	}
	l, err := app.logs.Get(id) // Get the record
	if err != nil {
		app.notFound(w)
		fmt.Println("RECORD NOT FOUND with ID - ", id)
		return
	}
	app.render(w, r, "logs.page.tmpl", templateData{Logs: []*models.Log{l}})
}
