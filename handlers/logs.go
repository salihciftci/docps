// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func logsHandler(w http.ResponseWriter, r *http.Request) {
	perm, err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	access := false
	for _, v := range perm {
		if string(v) == "l" || string(v) == "R" {
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
	params := r.URL.Query()
	key, ok := params["container"]

	if !ok || len(key) < 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	c, err := cmd.ParseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	l, err := cmd.ParseLogs(c)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	i := -1
	for p, v := range l {
		if v.Name == key[0] {
			i = p
		}
	}

	if i == -1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}

	data = append(data, bn)
	data = append(data, l[i].Name)
	data = append(data, l[i].Logs)

	err = tpl.ExecuteTemplate(w, "logs.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
