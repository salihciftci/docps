package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/pkg/notification"

	"github.com/salihciftci/liman/pkg/docker"
)

//Volumes asd
type Volumes struct {
	Driver string `json:"driver,omitempty"`
	Name   string `json:"name,omitempty"`
}

func parseVolumes() ([]Volumes, error) {
	cmdArgs := []string{
		"volume",
		"ls",
		"--format",
		"{{.Driver}}\t{{.Name}}",
	}
	stdOut, err := docker.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var volumes []Volumes
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		volumes = append(volumes,
			Volumes{
				Driver: s[0],
				Name:   s[1],
			})
	}

	return volumes, nil
}

//VolumesHandler asd
func VolumesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	v, err := parseVolumes()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := notification.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, v)

	err = tpl.ExecuteTemplate(w, "volumes.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
