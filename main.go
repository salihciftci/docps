package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func parseSessionCookie(w http.ResponseWriter, r *http.Request) error {
	if userPassword == "" {
		http.Redirect(w, r, "/install", http.StatusFound)
		log.Println(r.Method, r.URL.Path, "Not Installed")
		return fmt.Errorf("100")
	}

	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "session",
			Value: "",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusFound)
		return fmt.Errorf("101")
	}

	if cookie.Value != cookieValue {
		http.Redirect(w, r, "/login", http.StatusFound)
		return fmt.Errorf("102")
	}

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inputPass := r.FormValue("inputPassword")
		inputUser := r.FormValue("inputUser")
		if inputUser == username && inputPass == userPassword {
			cookie := &http.Cookie{
				Name:    "session",
				Value:   cookieValue,
				Path:    "/",
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)
		}
	}

	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	d, err := parseDashboard()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, d)

	err = tpl.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

	log.Println(r.Method, r.URL.Path)
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}
	c, err := parseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, c)

	err = tpl.ExecuteTemplate(w, "containers.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	s, err := parseStats()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, s)

	err = tpl.ExecuteTemplate(w, "stats.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	i, err := parseImages()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, i)

	err = tpl.ExecuteTemplate(w, "images.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	v, err := parseVolumes()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, v)

	err = tpl.ExecuteTemplate(w, "volumes.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func networksHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	n, err := parseNetworks()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, n)

	err = tpl.ExecuteTemplate(w, "networks.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	params := r.URL.Query()
	key, ok := params["container"]

	if !ok || len(key) < 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	c, err := parseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	l, err := parseLogs(c)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	i := -1
	for p, v := range l {
		if v.Name == key[0] {
			i = p
		}
	}

	if i == -1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	bn, _ := getNotification()

	var data []interface{}

	data = append(data, bn)
	data = append(data, l[i].Name)
	data = append(data, l[i].Logs)

	err = tpl.ExecuteTemplate(w, "logs.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}
	bn, n := getNotification()

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

}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	if r.Method == "POST" {
		pass := r.FormValue("cpass")
		if pass != userPassword {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		nPass := r.FormValue("npass")
		cNPass := r.FormValue("cnpass")

		if nPass != cNPass {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		if nPass == userPassword {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		userPassword = nPass
		http.Redirect(w, r, "/logout", http.StatusFound)
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, version)
	data = append(data, apiKey)

	err = tpl.ExecuteTemplate(w, "settings.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
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
		Value: "",
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	log.Println(r.Method, r.URL.Path)
	http.Redirect(w, r, "/", http.StatusFound)
}

func installHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if userPassword == "" {
			inputPassword := r.FormValue("inputPassword")
			userPassword = inputPassword
			http.Redirect(w, r, "/", http.StatusFound)
			log.Println(r.Method, r.URL.Path, "Install complete.")
			return
		}
	}

	if userPassword != "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := tpl.ExecuteTemplate(w, "install.tmpl", nil)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func main() {
	// Generating keys for api and cookie
	apiKey = generatePassword(32)
	cookieValue = generatePassword(140)

	// Checking containers for sending notification
	savedContainers, err := parseContainerStatus()
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			parseContainers, err := parseContainerStatus()
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
						notifications = append(notifications, notification{
							Desc:   savedContainers[i].Name + " is stopped.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "E",
						})
					}

					if savedContainers[i].Status == "E" {
						notifications = append(notifications, notification{
							Desc:   savedContainers[i].Name + " is started.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "U",
						})
					}
				}
			}
			savedContainers = parseContainers
		}
	}()

	// HTTP handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/containers", containersHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/images", imagesHandler)
	http.HandleFunc("/volumes", volumesHandler)
	http.HandleFunc("/networks", networksHandler)
	http.HandleFunc("/logs", logsHandler)
	http.HandleFunc("/settings", settingsHandler)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/install", installHandler)

	http.HandleFunc("/notifications", notificationHandler)

	http.HandleFunc("/api/containers", apiContainer)
	http.HandleFunc("/api/images", apiImages)
	http.HandleFunc("/api/volumes", apiVolumes)
	http.HandleFunc("/api/networks", apiNetworks)
	http.HandleFunc("/api/stats", apiStats)
	http.HandleFunc("/api/logs", apiLogs)
	http.HandleFunc("/api/status", apiStatus)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Listening http://0.0.0.0:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
