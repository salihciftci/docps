package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/salihciftci/liman/pkg/notification"

	"github.com/salihciftci/liman/pkg/routes"
)

var (
	tpl          *template.Template
	username     = "root"
	userPassword = ""
	apiKey       = ""
	cookieValue  = ""
	version      = "v0.6"
)

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*.tmpl"))
}

/*func notificationHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}
	bn, n := notification.GetNotification()

	if len(n) > 100 {
		n = n[:100]
	}

	var data []interface{}

	data = append(data, bn)
	data = append(data, n)

	err = tpl.ExecuteTemplate(w, "notifications.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)

}*/

func main() {
	// Generating keys for api and cookie
	apiKey = generatePassword(32)
	cookieValue = generatePassword(140)

	// Checking containers for sending notification
	savedContainers, err := notification.ParseContainerStatus()
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			parseContainers, err := notification.ParseContainerStatus()
			if err != nil {
				log.Println(err)
			}

			if len(parseContainers) != len(savedContainers) {
				savedContainers = parseContainers
				continue
			}

			for i, v := range savedContainers {
				if v.Status != parseContainers[i].Status {
					if savedContainers[i].Status == "U" {
						notification.Notifications = append(notification.Notifications, notification.Notification{
							Desc:   savedContainers[i].Name + " has stopped.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "E",
						})
					}

					if savedContainers[i].Status == "E" {
						notification.Notifications = append(notification.Notifications, notification.Notification{
							Desc:   savedContainers[i].Name + " has live.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "U",
						})
					}
				}
			}
			savedContainers = parseContainers
			time.Sleep(5 * time.Second)
		}
	}()

	// HTTP handlers
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/containers", routes.ContainersHandler)
	http.HandleFunc("/stats", routes.StatsHandler)
	http.HandleFunc("/images", routes.ImagesHandler)
	http.HandleFunc("/volumes", routes.VolumesHandler)
	http.HandleFunc("/networks", routes.NetworksHandler)
	http.HandleFunc("/logs", routes.LogsHandler)
	http.HandleFunc("/settings", routes.SettingsHandler)

	http.HandleFunc("/login", routes.LoginHandler)
	http.HandleFunc("/logout", routes.LogoutHandler)
	http.HandleFunc("/install", routes.InstallHandler)

	//http.HandleFunc("/notifications", notificationHandler)

	/*http.HandleFunc("/api/containers", apiContainer)
	http.HandleFunc("/api/images", apiImages)
	http.HandleFunc("/api/volumes", apiVolumes)
	http.HandleFunc("/api/networks", apiNetworks)
	http.HandleFunc("/api/stats", apiStats)
	http.HandleFunc("/api/logs", apiLogs)
	http.HandleFunc("/api/status", apiStatus)*/

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("../public"))))

	log.Println("Listening http://0.0.0.0:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func generatePassword(l int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
