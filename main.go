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
	cookie, err := r.Cookie("session")

	login := encode(false)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "session",
			Value: login,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
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
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	out := getDocker()
	err = tpl.ExecuteTemplate(w, "index.tmpl", out)
	if err != nil {
		log.Println(err.Error())
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	value := decode(cookie)

	if value == true {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err = tpl.ExecuteTemplate(w, "login.tmpl", nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	login := encode(false)

	cookie, _ := r.Cookie("session")
	cookie = &http.Cookie{
		Name:  "session",
		Value: login,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Listening :8080..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
