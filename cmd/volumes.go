package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/util"
)

type volume struct {
	Driver string `json:"driver,omitempty"`
	Name   string `json:"name,omitempty"`
}

func parseVolumes() ([]volume, error) {
	cmdArgs := []string{
		"volume",
		"ls",
		"--format",
		"{{.Driver}}\t{{.Name}}",
	}
	stdOut, err := util.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var volumes []volume
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		volumes = append(volumes,
			volume{
				Driver: s[0],
				Name:   s[1],
			})
	}

	return volumes, nil
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	v, err := parseVolumes()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, v)

	err = tpl.ExecuteTemplate(w, "volumes.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
