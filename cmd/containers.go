package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/util"
)

type ps struct {
	Name       string `json:"name,omitempty"`
	Image      string `json:"image,omitempty"`
	Size       string `json:"size,omitempty"`
	RunningFor string `json:"runningFor,omitempty"`
	Status     string `json:"status,omitempty"`
	Ports      string `json:"ports,omitempty"`
}

func parseContainers() ([]ps, error) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}\t{{.Ports}}",
	}
	stdOut, err := util.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var container []ps
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			ps{
				Name:       s[0],
				Image:      s[1],
				Size:       s[2],
				RunningFor: s[3],
				Status:     s[4][:1],
				Ports:      s[5],
			})
		container[i].Ports = util.ShortPorts(container[i].Ports)
	}

	return container, nil
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}
	c, err := parseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, c)

	err = tpl.ExecuteTemplate(w, "containers.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
