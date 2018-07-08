package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/util"
)

var (
	username = "root"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inputPass := r.FormValue("inputPassword")
		inputUser := r.FormValue("inputUser")
		match := util.CheckPass(inputPass, userPassword)
		if inputUser == username && match {
			cookie := &http.Cookie{
				Name:    "session",
				Value:   cookieValue,
				Path:    "/",
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)
		}
	}

	err := parseSessionCookie(w, r)
	if err != nil {
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
