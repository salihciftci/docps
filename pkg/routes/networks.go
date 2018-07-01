package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/pkg/notification"

	"github.com/salihciftci/liman/pkg/docker"
)

type networkss struct {
	Name   string `json:"name,omitempty"`
	Driver string `json:"driver,omitempty"`
	Scope  string `json:"scope,omitempty"`
}

func parseNetworks() ([]networkss, error) {
	cmdArgs := []string{
		"network",
		"ls",
		"--format",
		"{{.Name}}\t{{.Driver}}\t{{.Scope}}",
	}
	stdOut, err := docker.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}
	var networks []networkss
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		networks = append(networks,
			networkss{
				Name:   s[0],
				Driver: s[1],
				Scope:  s[2],
			})
	}

	return networks, nil
}

//NetworksHandler asd
func NetworksHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	n, err := parseNetworks()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := notification.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, n)

	err = tpl.ExecuteTemplate(w, "networks.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
}
