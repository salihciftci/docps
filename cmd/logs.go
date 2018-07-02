package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/salihciftci/liman/pkg/tool"
)

type logs struct {
	Name string   `json:"name,omitempty"`
	Logs []string `json:"logs,omitempty"`
}

func parseLogs(container []ps) ([]logs, error) {
	l := []logs{}
	for i := 0; i < len(container); i++ {
		cmdArgs := []string{
			"logs",
			"--tail",
			"100",
			container[i].Name,
		}

		cLog, err := tool.Cmd(cmdArgs)
		if err != nil {
			return nil, fmt.Errorf("Docker daemon is not running")
		}

		if len(cLog) > 0 {
			x := []string{}
			for k := len(cLog) - 1; k != -1; k-- {
				x = append(x, cLog[k])
			}
			l = append(l, logs{
				Name: container[i].Name,
				Logs: x,
			})
		} else {
			l = append(l, logs{
				Name: container[i].Name,
				Logs: []string{},
			})
		}
	}

	return l, nil
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
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

	bn, _ := getNotification()

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
