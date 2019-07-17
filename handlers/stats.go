// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func statsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	s, err := cmd.ParseStats()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, s)

	err = tpl.ExecuteTemplate(w, "stats.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
