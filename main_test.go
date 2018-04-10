package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStats(t *testing.T) {
	cmdArgs := []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}

	stats := getStats(cmdArgs)

	if stats == nil {
		log.Println("stats returned nil")
	}
}

func TestPS(t *testing.T) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}",
	}

	ps := ps(cmdArgs)

	if ps == nil {
		log.Println("process state returned nil")
	}

}

func TestGetImages(t *testing.T) {
	cmdArgs := []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}

	images := getImages(cmdArgs)

	if images == nil {
		log.Println("images returned nil")
	}

}

func TestGetVolumes(t *testing.T) {
	cmdArgs := []string{
		"volume",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}

	volume := getVolumes(cmdArgs)

	if volume == nil {
		log.Println("volumes returned nil")
	}

}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	IndexHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("exğected status OK; got %v", res.Status)
	}

}
