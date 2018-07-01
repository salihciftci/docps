package routes

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/pkg/notification"
)

var (
	version = "test"
	apiKey  = "123"
)

//SettingsHandler aa
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	if r.Method == "POST" {
		pass := r.FormValue("cpass")
		if pass != userPassword {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		nPass := r.FormValue("npass")
		cNPass := r.FormValue("cnpass")

		if nPass != cNPass {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		if nPass == userPassword {
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}

		userPassword = nPass
		http.Redirect(w, r, "/logout", http.StatusFound)
	}

	bn, _ := notification.GetNotification()

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
