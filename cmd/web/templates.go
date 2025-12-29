package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/T2Knock/snippetbox/internal/models"
	"github.com/T2Knock/snippetbox/ui"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, error := fs.Glob(ui.Files, "html/pages/*.html")
	if error != nil {
		return nil, error
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"html/base.html",
			"html/partials/nav.html",
			page,
		}

		tmpl, err := template.ParseFS(ui.Files, files...)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
