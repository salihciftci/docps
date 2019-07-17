// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"net/http"
	"time"
)

//HTTPServer returns *http.Server
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

	r.HandleFunc(BaseURL+"/", indexHandler)
	r.HandleFunc(BaseURL+"/containers", containersHandler)
	r.HandleFunc(BaseURL+"/stats", statsHandler)
	r.HandleFunc(BaseURL+"/images", imagesHandler)
	r.HandleFunc(BaseURL+"/volumes", volumesHandler)
	r.HandleFunc(BaseURL+"/networks", networksHandler)
	r.HandleFunc(BaseURL+"/logs", logsHandler)
	r.HandleFunc(BaseURL+"/settings", settingsHandler)

	r.HandleFunc(BaseURL+"/login", loginHandler)
	r.HandleFunc(BaseURL+"/logout", logoutHandler)
	r.HandleFunc(BaseURL+"/install", installHandler)

	r.HandleFunc(BaseURL+"/notifications", notificationHandler)

	r.HandleFunc(BaseURL+"/api/containers", apiContainer)
	r.HandleFunc(BaseURL+"/api/images", apiImages)
	r.HandleFunc(BaseURL+"/api/volumes", apiVolumes)
	r.HandleFunc(BaseURL+"/api/networks", apiNetworks)
	r.HandleFunc(BaseURL+"/api/stats", apiStats)
	r.HandleFunc(BaseURL+"/api/logs", apiLogs)
	r.HandleFunc(BaseURL+"/api/status", apiStatus)

	r.Handle(BaseURL+"/css/", http.StripPrefix(BaseURL+"/css/", http.FileServer(http.Dir("public/css"))))
	r.Handle(BaseURL+"/img/", http.StripPrefix(BaseURL+"/img/", http.FileServer(http.Dir("public/img"))))
	r.Handle(BaseURL+"/scripts/", http.StripPrefix(BaseURL+"/scripts/", http.FileServer(http.Dir("public/scripts"))))
	return r
}
