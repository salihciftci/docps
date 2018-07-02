package liman

import (
	"log"
	"net/http"
)

//Run Running Liman
func Run() error {
	//Checking Notifications every 5 second
	checkNotifications()

	log.Println("Listening http://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		return err
	}

	return nil
}

func handler() http.Handler {
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
