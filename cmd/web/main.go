package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr   string
	static string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.static, "static", "./ui/static/", "Path to static assets")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.static))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /snippet/view", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  mux,
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
