package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
	tpl    *template.Template
	secret = os.Getenv("secret")
	pass   = os.Getenv("pass")
)

//IndexHandler writing all outPuts to http template
func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("login")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "login",
			Value: "0",
			Path:  "/",
		}
	}

	if r.Method == "POST" {
		value := r.FormValue("inputPassword")
		if value == pass {
			cookie = &http.Cookie{
				Name:  "login",
				Value: secret,
				Path:  "/",
			}
		}
	}

	http.SetCookie(w, cookie)

	if cookie.Value != secret {
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
	cookie, err := r.Cookie("login")

	if cookie.Value == secret {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err = tpl.ExecuteTemplate(w, "login.tmpl", nil)
	if err != nil {
		log.Println(err.Error())
	}

}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func main() {
	letters := []rune(os.Getenv("secret"))
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	secret = string(b)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Listening :8080..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
