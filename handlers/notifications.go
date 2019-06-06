// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	bn, n := cmd.GetNotification()

	if len(n) > 100 {
		n = n[:100]
	}

	var data []interface{}

	data = append(data, bn)
	data = append(data, n)

	err = tpl.ExecuteTemplate(w, "notifications.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)

}
