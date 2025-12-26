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
	cfg      config
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
		cfg:      cfg,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
