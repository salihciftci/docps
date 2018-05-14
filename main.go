package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
)

var (
	tpl  *template.Template
	pass = os.Getenv("pass")
	s    = securecookie.New([]byte(securecookie.GenerateRandomKey(64)), []byte(securecookie.GenerateRandomKey(32)))
)

func encode(value bool) string {
	valuemap := map[string]bool{
		"liman": value,
	}

	encode, _ := s.Encode("session", valuemap)
	return encode
}

func decode(cookie *http.Cookie) bool {
	value := map[string]bool{
		"liman": false,
	}

	s.Decode("session", cookie.Value, &value)
	return value["liman"]
}

//IndexHandler writing all outPuts to http template
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println(r.Method, http.StatusNotFound, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cookie, err := r.Cookie("session")

	login := encode(false)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "session",
			Value: login,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		log.Println(r.Method, http.StatusFound, r.URL.Path)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	login = encode(true)
	if r.Method == "POST" {
		input := r.FormValue("inputPassword")
		if input == pass {
			cookie = &http.Cookie{
				Name:  "session",
				Value: login,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}
	}

	value := decode(cookie)
	if value != true {
		log.Println(r.Method, http.StatusFound, r.URL.Path)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var data []interface{}

	container, err := container()
	if err != nil {
		return
	}
	data = append(data, container)

	images, err := images()
	if err != nil {
		return
	}
	data = append(data, images)

	volumes, err := volumes()
	if err != nil {
		return
	}
	data = append(data, volumes)

	stats, err := stats()
	if err != nil {
		return
	}
	data = append(data, stats)

	logs, err := logs(container)
	if err != nil {
		return
	}
	data = append(data, logs)

	networks, err := networks()
	if err != nil {
		return
	}
	data = append(data, networks)

	dashboard, err := dashboard(images, volumes, networks)
	if err != nil {
		return
	}
	data = append(data, dashboard)

	err = tpl.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		log.Println(r.Method, http.StatusNotFound, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(r.Method, http.StatusFound, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	value := decode(cookie)

	if value == true {
		log.Println(r.Method, http.StatusFound, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err = tpl.ExecuteTemplate(w, "login.tmpl", nil)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(r.Method, http.StatusOK, r.URL.Path)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		log.Println(r.Method, http.StatusFound, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	login := encode(false)

	cookie, _ := r.Cookie("session")
	cookie = &http.Cookie{
		Name:  "session",
		Value: login,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	log.Println(r.Method, http.StatusOK, r.URL.Path)
	http.Redirect(w, r, "/", http.StatusFound)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.HandleFunc("/api/containers", apiGET)
	http.HandleFunc("/api/images", apiGET)
	http.HandleFunc("/api/volumes", apiGET)
	http.HandleFunc("/api/stats", apiGET)
	http.HandleFunc("/api/logs", apiGET)
	http.HandleFunc("/api/networks", apiGET)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Listening http://0.0.0.0:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
