package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr   string
	static string
	dsn    string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	cfg      config
	db       *sql.DB
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.static, "static", "./ui/static/", "Path to static assets")
	flag.StringVar(&cfg.dsn, "dsnc", "root:snippetbox@/snippetbox", "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			errorLog.Printf("error closing the database %v", err)
		}
	}()

	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
		cfg:      cfg,
		db:       db,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
