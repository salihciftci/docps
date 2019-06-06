// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/db/sqlite"
	"github.com/salihciftci/liman/util"
	"golang.org/x/crypto/bcrypt"
)

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	if r.Method == "POST" {
		pass := r.FormValue("cpass")
		cookie, err := r.Cookie("session")
		user := util.ParseJWT(cookie.Value, secretKey)
		strUser := fmt.Sprintf("%v", user)

		hash, err := sqlite.GetUserPassword(strUser)
		if err != nil {
			log.Println(err)
			return
		}

		match := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
		if match != nil {
			log.Println(pass)
			log.Println(match)
			log.Println(hash)
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		nPass := r.FormValue("npass")
		cNPass := r.FormValue("cnpass")

		if nPass != cNPass {
			log.Println("Passwords are not equal")
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		match = bcrypt.CompareHashAndPassword([]byte(hash), []byte(nPass))

		if match == nil {
			log.Println("wut")
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		nHash, err := bcrypt.GenerateFromPassword([]byte(nPass), 14)
		if err != nil {
			log.Println("wuuuuut")
			log.Println(err)
			return
		}

		err = sqlite.ChangeUserPassword(strUser, string(nHash))
		if err != nil {
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/logout", http.StatusFound)
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, Version)
	data = append(data, APIKey)

	err = tpl.ExecuteTemplate(w, "settings.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
