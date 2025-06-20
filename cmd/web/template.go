package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/andyaspel/snippetbox/pkg/models"
)

type logEntry struct {
	Time      string
	Method    string
	URL       string
	Status    int
	Duration  string
	Remote    string
	UserAgent string
}

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
	List     *models.List
	Lists    []*models.List
	Logs     []logEntry
}

func humanDate(t time.Time) string {
	// Format the time as "2 Jan 2006 at 15:04"
	return t.Format("2 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

// study function  use print function to log output
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// print ts
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		// print ts
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		// print ts

		cache[name] = ts
		// print cache
	}

	return cache, nil
}
