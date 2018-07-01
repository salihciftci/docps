package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/salihciftci/liman/pkg/notification"

	"github.com/salihciftci/liman/pkg/docker"
)

//Logs asd
type Logs struct {
	Name string   `json:"name,omitempty"`
	Logs []string `json:"logs,omitempty"`
}

func parseLogs(container []ps) ([]Logs, error) {
	logs := []Logs{}
	for i := 0; i < len(container); i++ {
		cmdArgs := []string{
			"logs",
			"--tail",
			"100",
			container[i].Name,
		}

		cLog, err := docker.Cmd(cmdArgs)
		if err != nil {
			return nil, fmt.Errorf("Docker daemon is not running")
		}

		if len(cLog) > 0 {
			x := []string{}
			for k := len(cLog) - 1; k != -1; k-- {
				x = append(x, cLog[k])
			}
			logs = append(logs, Logs{
				Name: container[i].Name,
				Logs: x,
			})
		} else {
			logs = append(logs, Logs{
				Name: container[i].Name,
				Logs: []string{},
			})
		}
	}

	return logs, nil
}

//LogsHandler asd
func LogsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	params := r.URL.Query()
	key, ok := params["container"]

	if !ok || len(key) < 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	c, err := parseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	l, err := parseLogs(c)
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

	bn, _ := notification.GetNotification()

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
