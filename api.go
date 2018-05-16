package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//APIAuth checks api authentication
func APIAuth(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "METHOD_NOT_ALLOWED"})
		return fmt.Errorf("METHOD_NOT_ALLOWED")
	}

	params := r.URL.Query()
	key, ok := params["key"]

	if !ok || len(key) < 1 {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "API_KEY_NOT_FOUND"})
		return fmt.Errorf("API_KEY_NOT_FOUND")
	}

	if string(key[0]) != apiKey {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "API_KEY_NOT_INVALID"})
		return fmt.Errorf("API_KEY_NOT_INVALID")
	}

	return nil
}

//APIContainer response /api/containers requests
func APIContainer(w http.ResponseWriter, r *http.Request) {
	err := APIAuth(w, r)
	if err != nil {
		log.Println(r.Method, http.StatusOK, r.URL.Path, err)
		return
	}

	container, err := container()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}
	json.NewEncoder(w).Encode(container)
	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

//APIImages response /api/images requests
func APIImages(w http.ResponseWriter, r *http.Request) {
	err := APIAuth(w, r)
	if err != nil {
		log.Println(r.Method, http.StatusOK, r.URL.Path, err)
		return
	}

	images, err := images()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}
	json.NewEncoder(w).Encode(images)
	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

//APIVolumes response /api/volumes requests
func APIVolumes(w http.ResponseWriter, r *http.Request) {
	err := APIAuth(w, r)
	if err != nil {
		log.Println(r.Method, http.StatusOK, r.URL.Path, err)
		return
	}

	volumes, err := volumes()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}
	json.NewEncoder(w).Encode(volumes)
	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

//APINetworks response /api/networks requests
func APINetworks(w http.ResponseWriter, r *http.Request) {
	err := APIAuth(w, r)
	if err != nil {
		log.Println(r.Method, http.StatusOK, r.URL.Path, err)
		return
	}

	networks, err := networks()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}
	json.NewEncoder(w).Encode(networks)
	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

//APIStats response /api/stats requests
func APIStats(w http.ResponseWriter, r *http.Request) {
	err := APIAuth(w, r)
	if err != nil {
		log.Println(r.Method, http.StatusOK, r.URL.Path, err)
		return
	}

	stats, err := stats()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}
	json.NewEncoder(w).Encode(stats)
	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

//APILogs response /api/logs requests
func APILogs(w http.ResponseWriter, r *http.Request) {
	err := APIAuth(w, r)
	if err != nil {
		log.Println(r.Method, http.StatusOK, r.URL.Path, err)
		return
	}

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
