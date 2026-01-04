package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/T2Knock/snippetbox/internal/models"
	"github.com/T2Knock/snippetbox/ui"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, error := fs.Glob(ui.Files, "html/pages/*.html")
	if error != nil {
		return nil, error
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.New(name).Funcs(functions).ParseFS(ui.Files, "html/base.html", "html/partials/*.html", page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
