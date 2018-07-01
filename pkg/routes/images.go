package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/pkg/docker"
	"github.com/salihciftci/liman/pkg/notification"
)

type imagess struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Created    string `json:"created,omitempty"`
	Size       string `json:"size,omitempty"`
}

func parseImages() ([]imagess, error) {
	cmdArgs := []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	stdOut, err := docker.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var images []imagess
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		images = append(images,
			imagess{
				Repository: s[0],
				Tag:        s[1],
				Created:    s[2],
				Size:       s[3],
			})
	}
	return images, nil
}

//ImagesHandler asd
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	i, err := parseImages()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := notification.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, i)

	err = tpl.ExecuteTemplate(w, "images.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
