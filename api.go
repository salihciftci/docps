package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// APIStatus asd
func APIStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "METHOD_NOT_ALLOWED",
			})
		log.Println(r.Method, http.StatusOK, r.URL.Path)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":     "true",
		"result": nil,
	})

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

// APIAuth checks api authentication
func APIAuth(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "METHOD_NOT_ALLOWED",
			})
		return fmt.Errorf("METHOD_NOT_ALLOWED")
	}

	params := r.URL.Query()
	key, ok := params["key"]

	if !ok || len(key) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "API_KEY_NOT_FOUND",
			})
		return fmt.Errorf("API_KEY_NOT_FOUND")
	}

	if string(key[0]) != apiKey {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "API_KEY_NOT_INVALID",
			})
		return fmt.Errorf("API_KEY_NOT_INVALID")
	}

	return nil
}

// APIContainer response /api/containers requests
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": container,
		})
	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

// APIImages response /api/images requests
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": images,
		})

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

// APIVolumes response /api/volumes requests
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": volumes,
		})

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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": networks,
		})

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

// APIStats response /api/stats requests
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": stats,
		})

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

// APILogs response /api/logs requests
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": logs,
		})

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

// GenerateAPIPassword generates a random 32 length password for API
func GenerateAPIPassword() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
