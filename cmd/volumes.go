// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/salihciftci/liman/util"
)

//Volume stores values for handlers
type Volume struct {
	Driver string `json:"driver,omitempty"`
	Name   string `json:"name,omitempty"`
}

//ParseVolumes is parses and splits volume values from output
func ParseVolumes() ([]Volume, error) {
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

	var volumes []Volume
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		volumes = append(volumes,
			Volume{
				Driver: s[0],
				Name:   s[1],
			})
	}

	return volumes, nil
}
