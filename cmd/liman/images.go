package liman

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/pkg/tool"
)

type image struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Created    string `json:"created,omitempty"`
	Size       string `json:"size,omitempty"`
}

func parseImages() ([]image, error) {
	cmdArgs := []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	stdOut, err := tool.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var images []image
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		images = append(images,
			image{
				Repository: s[0],
				Tag:        s[1],
				Created:    s[2],
				Size:       s[3],
			})
	}
	return images, nil
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	i, err := parseImages()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, i)

	err = tpl.ExecuteTemplate(w, "images.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
