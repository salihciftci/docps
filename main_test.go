package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	indexHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Errorf("expected status found; got %v", res.Status)
	}
}

func TestLoginHandler(t *testing.T) {
	req, err := http.NewRequest("Get", "localhost:8080/login", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()
	loginHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		t.Errorf("expected status found; got: %v", res.Status)
	}
}

func TestLogoutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/logout", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()
	logoutHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		t.Errorf("expected status found; got %v", res.Status)
	}

}

func TestSetCookie(t *testing.T) {
	rec := httptest.NewRecorder()

	http.SetCookie(rec, &http.Cookie{Name: "session", Value: "test"})

	req := &http.Request{Header: http.Header{"Cookie": rec.HeaderMap["Set-Cookie"]}}

	_, err := req.Cookie("session")
	if err != nil {
		t.Errorf("couldn't read Cookie")
	}
}
