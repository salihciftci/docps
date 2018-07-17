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
	//IsInstalled boolen for Liman already Installed or not
	IsInstalled = false

	//Version of Liman
	Version = "0.6-develop"

	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
)

func parseSessionCookie(w http.ResponseWriter, r *http.Request) error {
	if !IsInstalled {
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

	user, err := sqlite.GetUserFromSessionKey(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}

	if user == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return fmt.Errorf("102")
	}

	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if !IsInstalled {
		http.Redirect(w, r, "/install", http.StatusFound)
	}

	if r.Method == "POST" {
		inputPass := r.FormValue("inputPassword")
		inputUser := r.FormValue("inputUser")
		log.Println(inputUser)
		hash, key, err := sqlite.GetUserPasswordAndSessionKey(inputUser)
		if err != nil {
			log.Println(r.Method, r.URL.Path, "User not found.")
		}

		match := bcrypt.CompareHashAndPassword([]byte(hash), []byte(inputPass))

		if match == nil {
			cookie := &http.Cookie{
				Name:    "session",
				Value:   key,
				Path:    "/",
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}

	err := tpl.ExecuteTemplate(w, "login.tmpl", nil)
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
	if IsInstalled {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if r.Method == "POST" {
		if !IsInstalled {
			inputPassword := r.FormValue("inputPassword")
			inputUser := r.FormValue("inputUser")

			hash, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
			if err != nil {
				log.Println(err)
				return
			}

			sessionKey := util.GenerateKey(140)
			apiKey := util.GenerateKey(40)

			err = sqlite.Install(inputUser, string(hash), sessionKey, apiKey, Version)
			if err != nil {
				log.Println(err)
				return
			}

			IsInstalled = true
			APIKey = apiKey

			cookie := &http.Cookie{
				Name:    "session",
				Value:   sessionKey,
				Path:    "/",
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
			log.Println("Installation complete.")
			return
		}
	}

	err := tpl.ExecuteTemplate(w, "install.tmpl", nil)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
