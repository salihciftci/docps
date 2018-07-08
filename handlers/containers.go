package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func containersHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}
	c, err := cmd.ParseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, c)

	err = tpl.ExecuteTemplate(w, "containers.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
