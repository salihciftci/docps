package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIAuth(t *testing.T) {
	apiKey = "test"
	req, err := http.NewRequest("GET", "localhost:8080/api/containers?key=test", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APIContainer(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status ok; got %v", res.Status)
	}
}
func TestAPIContainer(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/containers", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APIContainer(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res.Status)
	}

}

func TestAPIImages(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/images", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APIImages(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res.Status)
	}

}

func TestAPIVolumes(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/volumes", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APIVolumes(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res.Status)
	}

}

func TestAPINetworks(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/networks", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APINetworks(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res.Status)
	}

}
func TestAPIStats(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/stats", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APIStats(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res.Status)
	}

}

func TestAPILogs(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/api/logs", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	APILogs(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res.Status)
	}

}
