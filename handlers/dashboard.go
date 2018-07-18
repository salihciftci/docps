// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	perm, err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	access := false
	for _, v := range perm {
		if string(v) == "d" || string(v) == "R" {
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

	d, err := cmd.ParseDashboard()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, d)

	err = tpl.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

	log.Println(r.Method, r.URL.Path)
}
