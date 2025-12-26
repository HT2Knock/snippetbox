package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.cfg.static))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /snippet/view", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)

	return mux
}
