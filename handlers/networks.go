package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func networksHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	n, err := cmd.ParseNetworks()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, n)

	err = tpl.ExecuteTemplate(w, "networks.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
}
