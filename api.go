package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func apiGET(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	params := r.URL.Query()
	key, ok := params["key"]

	if !ok || len(key) < 1 {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "API_KEY_NOT_FOUND"})
		log.Println(r.Method, http.StatusOK, r.URL.Path, "API_KEY_NOT_FOUND")
		return
	}

	if string(key[0]) != apiKey {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "API_KEY_INVALID"})
		log.Println(r.Method, http.StatusOK, r.URL.Path, "API_KEY_INVALID")
		return
	}

	switch r.URL.Path {
	case "/api/containers":
		container, err := container()
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}
		json.NewEncoder(w).Encode(container)
		log.Println(r.Method, http.StatusOK, r.URL.Path)
	case "/api/images":
		images, err := images()
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}
		json.NewEncoder(w).Encode(images)
		log.Println(r.Method, http.StatusOK, r.URL.Path)
	case "/api/volumes":
		volumes, err := volumes()
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}
		json.NewEncoder(w).Encode(volumes)
		log.Println(r.Method, http.StatusOK, r.URL.Path)
	case "/api/stats":
		stats, err := stats()
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}
		json.NewEncoder(w).Encode(stats)
		log.Println(r.Method, http.StatusOK, r.URL.Path)
	case "/api/networks":
		networks, err := networks()
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}
		json.NewEncoder(w).Encode(networks)
		log.Println(r.Method, http.StatusOK, r.URL.Path)
	case "/api/logs":
		container, err := container()
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}

		logs, err := logs(container)
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
			return
		}
		json.NewEncoder(w).Encode(logs)
		log.Println(r.Method, http.StatusOK, r.URL.Path)
	}
}

//generateAPIPassword generates a random 32 length password for API
func generateAPIPassword() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
