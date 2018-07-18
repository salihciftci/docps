// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/db/sqlite"
	"github.com/salihciftci/liman/util"
	"golang.org/x/crypto/bcrypt"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
	perm, err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	access := false
	for _, v := range perm {
		if string(v) == "R" {
			access = true
		}
	}

	if !access {
		err = tpl.ExecuteTemplate(w, "permission.tmpl", nil)
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
		}
		return
	}
	bn, _ := cmd.GetNotification()

	users, err := sqlite.ListUsers()
	if err != nil {
		log.Println(err)
	}

	var data []interface{}
	data = append(data, bn)
	data = append(data, "")
	data = append(data, users)

	if r.Method == "POST" {
		var perm string
		user := r.FormValue("username")
		password := r.FormValue("password")
		cpassword := r.FormValue("cpassword")
		desc := r.FormValue("desc")
		key := util.GenerateKey(140)

		userIsExist, err := sqlite.CheckUserExist(user)
		if err != nil {
			log.Println(err)
			return
		}

		if userIsExist {
			data[1] = "exist"
			err = tpl.ExecuteTemplate(w, "users.tmpl", data)
			if err != nil {
				log.Println(r.Method, r.URL.Path, err)
			}
			return
		}

		if password != cpassword {
			data[1] = "match"
			err = tpl.ExecuteTemplate(w, "users.tmpl", data)
			if err != nil {
				log.Println(r.Method, r.URL.Path, err)
			}
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			log.Println(err)
			return
		}

		admin := r.FormValue("admin")

		if admin != "" {
			perm = "R"
			err = sqlite.CreateUser(user, string(hash), key, perm, desc)
			if err != nil {
				log.Println(err)
			}
			data[1] = "success"
			err = tpl.ExecuteTemplate(w, "users.tmpl", data)
			if err != nil {
				log.Println(r.Method, r.URL.Path, err)
			}
			return

		}

		perm = r.FormValue("dashboard") + r.FormValue("containers") + r.FormValue("stats") + r.FormValue("images")
		perm = perm + r.FormValue("volumes") + r.FormValue("networks") + r.FormValue("logs") + r.FormValue("notifications")

		err = sqlite.CreateUser(user, string(hash), key, perm, desc)
		if err != nil {
			log.Println(err)
		}
		data[1] = "success"
		err = tpl.ExecuteTemplate(w, "users.tmpl", data)
		if err != nil {
			log.Println(r.Method, r.URL.Path, err)
		}
		return

	}
	err = tpl.ExecuteTemplate(w, "users.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

	log.Println(r.Method, r.URL.Path)
}
