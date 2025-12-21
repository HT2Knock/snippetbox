package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_home(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	home(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retured the wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from snippet"
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Errorf("handler returned the wrong body: got %v want %v", string(body), expected)
	}
}

func Test_snippetView(t *testing.T) {
	req, err := http.NewRequest("GET", "/snippet/view?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	snippetView(rr, req)

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

	snippetCreate(rr, req)

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
