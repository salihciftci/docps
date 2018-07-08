package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	i, err := cmd.ParseImages()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, i)

	err = tpl.ExecuteTemplate(w, "images.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
