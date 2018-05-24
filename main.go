package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tpl         *template.Template
	pass        = os.Getenv("pass")
	apiKey      = ""
	cookieValue = ""
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func cookieCheck(w http.ResponseWriter, r *http.Request) {
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
		input := r.FormValue("inputPassword")
		if input == pass {
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
	cookieCheck(w, r)

	dashboard, err := dashboard()
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
	cookieCheck(w, r)

	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))

	containers, err := container()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	logs, err := logs(containers)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, containers)
	data = append(data, logs)

	err = tpl.ExecuteTemplate(w, "containers.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	cookieCheck(w, r)

	stats, err := stats()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, stats)

	err = tpl.ExecuteTemplate(w, "stats.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	cookieCheck(w, r)

	images, err := images()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, images)

	err = tpl.ExecuteTemplate(w, "images.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	cookieCheck(w, r)

	volumes, err := volumes()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, volumes)

	err = tpl.ExecuteTemplate(w, "volumes.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}

func networksHandler(w http.ResponseWriter, r *http.Request) {
	cookieCheck(w, r)

	networks, err := networks()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, apiKey)
	data = append(data, networks)

	err = tpl.ExecuteTemplate(w, "networks.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	apiKey = generatePassword(32)
	cookieValue = generatePassword(140)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/containers", containersHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/images", imagesHandler)
	http.HandleFunc("/volumes", volumesHandler)
	http.HandleFunc("/networks", networksHandler)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.HandleFunc("/api/containers", APIContainer)
	http.HandleFunc("/api/images", APIImages)
	http.HandleFunc("/api/volumes", APIVolumes)
	http.HandleFunc("/api/networks", APINetworks)
	http.HandleFunc("/api/stats", APIStats)
	http.HandleFunc("/api/logs", APILogs)
	http.HandleFunc("/api/status", APIStatus)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Listening http://0.0.0.0:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
