// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/salihciftci/liman/util"
)

//Logs stores values for handlers
type Logs struct {
	Name string   `json:"name,omitempty"`
	Logs []string `json:"logs,omitempty"`
}

//ParseLogs is parses and splits logs from output
func ParseLogs(container []PS) ([]Logs, error) {
	l := []Logs{}
	for i := 0; i < len(container); i++ {
		cmdArgs := []string{
			"logs",
			"--tail",
			"100",
			container[i].Name,
		}

		cLog, err := util.Cmd(cmdArgs)
		if err != nil {
			return nil, fmt.Errorf("Docker daemon is not running")
		}

		if len(cLog) > 0 {
			x := []string{}
			for k := len(cLog) - 1; k != -1; k-- {
				x = append(x, cLog[k])
			}
			l = append(l, Logs{
				Name: container[i].Name,
				Logs: x,
			})
		} else {
			l = append(l, Logs{
				Name: container[i].Name,
				Logs: []string{},
			})
		}
	}

	return l, nil
}
