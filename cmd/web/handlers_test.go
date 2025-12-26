package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var app = application{
	infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
	errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
}

func Test_home(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.home(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retured the wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<title>Home</title>"
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), expected) {
		t.Errorf("handler returned the wrong body: got %v want %v", string(body), expected)
	}
}

func Test_snippetView(t *testing.T) {
	req, err := http.NewRequest("GET", "/snippet/view?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.snippetView(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retured the wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Display a specific snippet for id 1..."
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Errorf("handler returned the wrong body: got %v want %v", string(body), expected)
	}
}

func Test_snippetCreate(t *testing.T) {
	req, err := http.NewRequest("POST", "/snippet/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.snippetCreate(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retured the wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Create a new snippet..."
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Errorf("handler returned the wrong body: got %v want %v", string(body), expected)
	}
}
