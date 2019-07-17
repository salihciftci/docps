// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/salihciftci/liman/util"
)

//PS stores values for handlers
type PS struct {
	Name       string `json:"name,omitempty"`
	Image      string `json:"image,omitempty"`
	Size       string `json:"size,omitempty"`
	RunningFor string `json:"runningFor,omitempty"`
	Status     string `json:"status,omitempty"`
	Ports      string `json:"ports,omitempty"`
}

//ParseContainers is parses and splits containers process state values from output
func ParseContainers() ([]PS, error) {
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

	var container []PS
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			PS{
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
