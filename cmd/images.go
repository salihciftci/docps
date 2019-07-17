// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/salihciftci/liman/util"
)

//Image stores values for handlers
type Image struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Created    string `json:"created,omitempty"`
	Size       string `json:"size,omitempty"`
}

//ParseImages is parses and splits image values from output
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
