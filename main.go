package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	tpl             *template.Template
	username        = "root"
	userPassword    = ""
	apiKey          = ""
	cookieValue     = ""
	savedContainers []PS
	notifi          []notification
)

type notification struct {
	Desc   string
	Time   string
	Status string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func parseSessionCookie(w http.ResponseWriter, r *http.Request) {
	if userPassword == "" {
		http.Redirect(w, r, "/install", http.StatusFound)
		return
	}

	cookie, err := r.Cookie("session")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "session",
			Value: "0",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		log.Println(r.Method, r.URL.Path, err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method == "POST" {
		inputPass := r.FormValue("inputPassword")
		inputUser := r.FormValue("inputUser")
		if inputUser == username && inputPass == userPassword {
			cookie = &http.Cookie{
				Name:  "session",
				Value: cookieValue,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}
	}

	if cookie.Value != cookieValue {
		cookie = &http.Cookie{
			Name:  "session",
			Value: "0",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	dashboard, err := parseDashboard()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, dashboard)

	err = tpl.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	containers, err := parseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	notifiClear, _ := getNotification()

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, containers)
	data = append(data, notifiClear)

	err = tpl.ExecuteTemplate(w, "containers.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	stats, err := parseStats()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	notifiClear, _ := getNotification()

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, stats)
	data = append(data, notifiClear)

	err = tpl.ExecuteTemplate(w, "stats.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	images, err := parseImages()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	notifiClear, _ := getNotification()

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, images)
	data = append(data, notifiClear)

	err = tpl.ExecuteTemplate(w, "images.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	volumes, err := parseVolumes()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	notifiClear, _ := getNotification()

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, volumes)
	data = append(data, notifiClear)

	err = tpl.ExecuteTemplate(w, "volumes.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func networksHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	networks, err := parseNetworks()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	notifiClear, _ := getNotification()

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, networks)
	data = append(data, notifiClear)

	err = tpl.ExecuteTemplate(w, "networks.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	parseSessionCookie(w, r)

	params := r.URL.Query()
	key, ok := params["container"]

	if !ok || len(key) < 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	containers, err := parseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	logs, err := parseLogs(containers)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	container := -1
	for p, v := range logs {
		if v.Name == key[0] {
			container = p
		}
	}

	if container == -1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	notifiClear, _ := getNotification()

	var data []interface{}

	data = append(data, apiKey)
	data = append(data, notifiClear)
	data = append(data, logs[container].Name)
	data = append(data, logs[container].Logs)

	err = tpl.ExecuteTemplate(w, "logs.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if userPassword == "" {
		http.Redirect(w, r, "/install", http.StatusFound)
	}
	if r.URL.Path != "/login" {
		log.Println(r.Method, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(r.Method, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	value := cookie.Value

	if value == cookieValue {
		log.Println(r.Method, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err = tpl.ExecuteTemplate(w, "login.tmpl", nil)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

	log.Println(r.Method, r.URL.Path)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		log.Println(r.Method, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cookie, _ := r.Cookie("session")
	cookie = &http.Cookie{
		Name:  "session",
		Value: "0",
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	log.Println(r.Method, r.URL.Path)
	http.Redirect(w, r, "/", http.StatusFound)
}

func installHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inputPassword := r.FormValue("inputPassword")
		os.Setenv("LimanPass", inputPassword)
		userPassword = inputPassword
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if userPassword != "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := tpl.ExecuteTemplate(w, "install.tmpl", "")
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
}

func main() {
	// Checking os environment variable for root password
	envPass := os.Getenv("LimanPass")
	if envPass != "" {
		userPassword = envPass
	}
	// Generating keys for api and cookie
	apiKey = generatePassword(32)
	cookieValue = generatePassword(140)

	// HTTP handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/containers", containersHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/images", imagesHandler)
	http.HandleFunc("/volumes", volumesHandler)
	http.HandleFunc("/networks", networksHandler)
	http.HandleFunc("/logs", logsHandler)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/install", installHandler)

	http.HandleFunc("/api/containers", APIContainer)
	http.HandleFunc("/api/images", APIImages)
	http.HandleFunc("/api/volumes", APIVolumes)
	http.HandleFunc("/api/networks", APINetworks)
	http.HandleFunc("/api/stats", APIStats)
	http.HandleFunc("/api/logs", APILogs)
	http.HandleFunc("/api/status", APIStatus)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Checking containers status
	go func() {
		savedContainers, err := checkContainerStatus()
		if err != nil {
			log.Println(err)
		}

		time.Sleep(10 * time.Second)
		for {
			checkContainers, err := checkContainerStatus()
			if err != nil {
				log.Println(err)
			}

			for i := 0; i < len(checkContainers); i++ {
				if savedContainers[i].Status != checkContainers[i].Status {
					if savedContainers[i].Status == "U" {
						log.Println(savedContainers[i].Name + "is stopped.")
						notifi = append(notifi, notification{
							Desc:   savedContainers[i].Name + " is stopped.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "E",
						})
					}

					if savedContainers[i].Status == "E" {
						log.Println(savedContainers[i].Name + " is started.")
						notifi = append(notifi, notification{
							Desc:   savedContainers[i].Name + " is started.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "U",
						})
					}
				}
			}

			savedContainers = checkContainers
			time.Sleep(time.Second * 60) // Waiting 60 second to checking container statuses again.
		}
	}()

	log.Println("Listening http://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
