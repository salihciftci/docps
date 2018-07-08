// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/util"
)

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	if r.Method == "POST" {
		pass := r.FormValue("cpass")

		match := util.CheckPass(pass, userPassword)
		if !match {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		nPass := r.FormValue("npass")
		cNPass := r.FormValue("cnpass")

		if nPass != cNPass {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		match = util.CheckPass(nPass, userPassword)

		if match {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		bNPass, err := util.HashPasswordAndSave(nPass)
		if err != nil {
			log.Println(err)
			return
		}

		userPassword = string(bNPass)
		http.Redirect(w, r, "/logout", http.StatusFound)
	}

	bn, _ := cmd.GetNotification()

	version, err := util.Version()
	if err != nil {
		log.Println(err)
	}

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
