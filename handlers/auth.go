package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/salihciftci/liman/util"
)

var (
	tpl          = template.Must(template.ParseGlob("templates/*.tmpl"))
	cookieValue  = util.GeneratePassword(140)
	userPassword = util.ReadPassword()
)

func parseSessionCookie(w http.ResponseWriter, r *http.Request) error {
	if userPassword == "" {
		http.Redirect(w, r, "/install", http.StatusFound)
		log.Println("Installation started.")
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

	cookie := &http.Cookie{
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
			hash, err := util.HashPasswordAndSave(inputPassword)
			if err != nil {
				log.Println(err)
				return
			}
			userPassword = hash
			http.Redirect(w, r, "/", http.StatusFound)
			log.Println("Installation complete.")
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
