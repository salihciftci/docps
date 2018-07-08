package cmd

import (
	"fmt"
	"strings"

	"github.com/salihciftci/liman/util"
)

//Image asd
type Image struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Created    string `json:"created,omitempty"`
	Size       string `json:"size,omitempty"`
}

//ParseImages asd
func ParseImages() ([]Image, error) {
	cmdArgs := []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	stdOut, err := util.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var images []Image
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		images = append(images,
			Image{
				Repository: s[0],
				Tag:        s[1],
				Created:    s[2],
				Size:       s[3],
			})
	}
	return images, nil
}
