package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIFunc(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/containers", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	req, err = http.NewRequest("GET", "localhost:8080/api/images", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	req, err = http.NewRequest("GET", "localhost:8080/api/volumes", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	req, err = http.NewRequest("GET", "localhost:8080/api/stats", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	req, err = http.NewRequest("GET", "localhost:8080/api/logs", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	req, err = http.NewRequest("GET", "localhost:8080/api/networks", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	apiGET(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status found; got %v", res.Status)
	}
}
