// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"net/http"
	"time"
)

//HTTPServer asd
func HTTPServer() *http.Server {
	return &http.Server{
		Addr:         ":8080",
		Handler:      handlers(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
}

func handlers() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/containers", containersHandler)
	r.HandleFunc("/stats", statsHandler)
	r.HandleFunc("/images", imagesHandler)
	r.HandleFunc("/volumes", volumesHandler)
	r.HandleFunc("/networks", networksHandler)
	r.HandleFunc("/logs", logsHandler)
	r.HandleFunc("/settings", settingsHandler)

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/install", installHandler)

	r.HandleFunc("/notifications", notificationHandler)

	r.HandleFunc("/api/containers", apiContainer)
	r.HandleFunc("/api/images", apiImages)
	r.HandleFunc("/api/volumes", apiVolumes)
	r.HandleFunc("/api/networks", apiNetworks)
	r.HandleFunc("/api/stats", apiStats)
	r.HandleFunc("/api/logs", apiLogs)
	r.HandleFunc("/api/status", apiStatus)

	r.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	return r
}
