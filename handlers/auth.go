// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/salihciftci/liman/db/sqlite"
	"github.com/salihciftci/liman/util"
	"golang.org/x/crypto/bcrypt"
)

var (
	IsInstalled = false
	Version     = "0.6-develop"
	BaseURL     = ""

	secretKey  = util.GenerateSecretKey(120)
	sessionKey = ""

	pageError = ""

	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
)

func parseSessionCookie(w http.ResponseWriter, r *http.Request) error {
	if !IsInstalled {
		http.Redirect(w, r, BaseURL+"/install", http.StatusFound)
		log.Println("Installation started.")
		return fmt.Errorf("Not Installed")
	}

	if len(sessionKey) == 0 {
		http.Redirect(w, r, BaseURL+"/login", http.StatusFound)
		return fmt.Errorf("Session key not generated")
	}

	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "session",
			Value: "",
			Path:  BaseURL,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, BaseURL+"/login", http.StatusFound)
		return fmt.Errorf("Session not found")
	}

	if cookie.Value != sessionKey {
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, BaseURL+"/login", http.StatusFound)
		log.Println(r.Method, r.URL.Path, "Not logged in")
	}

	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if !IsInstalled {
		http.Redirect(w, r, BaseURL+"/install", http.StatusFound)
	}

	if r.Method == "POST" {
		inputPass := r.FormValue("inputPassword")
		inputUser := r.FormValue("inputUser")
		hash, err := sqlite.GetUserPassword(inputUser)
		if err != nil {
			pageError = "Invalid username or password"
			log.Println(r.Method, r.URL.Path, "User not found.")
		}

		match := bcrypt.CompareHashAndPassword([]byte(hash), []byte(inputPass))

		if match == nil {
			if len(sessionKey) == 0 {
				sessionKey = util.GenerateJWT(inputUser, secretKey)
				APIKey = util.GenerateJWT(inputUser, secretKey)
			}

			cookie := &http.Cookie{
				Name:    "session",
				Value:   sessionKey,
				Path:    BaseURL,
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, BaseURL, http.StatusFound)
			return
		}

		pageError = "Invalid username or password"
		log.Println(r.Method, r.URL.Path, "User not found.")
	}

	var data []interface{}
	data = append(data, pageError)
	pageError = ""
	err := tpl.ExecuteTemplate(w, "login.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

	log.Println(r.Method, r.URL.Path)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "",
		Path:  BaseURL,
	}

	http.SetCookie(w, cookie)
	log.Println(r.Method, r.URL.Path)
	http.Redirect(w, r, BaseURL+"/login", http.StatusFound)
}

func installHandler(w http.ResponseWriter, r *http.Request) {
	if IsInstalled {
		http.Redirect(w, r, BaseURL+"/", http.StatusFound)
		return
	}

	if r.Method == "POST" {
		if !IsInstalled {
			inputPassword := r.FormValue("inputPassword")
			inputUser := r.FormValue("inputUser")

			if len(inputPassword) == 0 || len(inputUser) == 0 {
				pageError = "Username and Password fields are required!"
				http.Redirect(w, r, BaseURL+"/install", http.StatusFound)
				return
			}

			hash, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
			if err != nil {
				log.Println(err)
				return
			}

			sessionKey := util.GenerateJWT(inputUser, secretKey)
			apiKey := util.GenerateJWT(inputUser, secretKey)

			err = sqlite.Install(inputUser, string(hash))
			if err != nil {
				log.Println(err)
				return
			}

			IsInstalled = true
			APIKey = apiKey

			cookie := &http.Cookie{
				Name:    "session",
				Value:   sessionKey,
				Path:    BaseURL,
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, BaseURL+"/", http.StatusFound)
			log.Println("Installation complete.")
			return
		}
	}

	var data []interface{}
	data = append(data, pageError)
	pageError = ""

	err := tpl.ExecuteTemplate(w, "install.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
