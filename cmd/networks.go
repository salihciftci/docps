package cmd

import (
	"fmt"
	"strings"

	"github.com/salihciftci/liman/util"
)

//Network aasd
type Network struct {
	Name   string `json:"name,omitempty"`
	Driver string `json:"driver,omitempty"`
	Scope  string `json:"scope,omitempty"`
}

//ParseNetworks asd
func ParseNetworks() ([]Network, error) {
	cmdArgs := []string{
		"network",
		"ls",
		"--format",
		"{{.Name}}\t{{.Driver}}\t{{.Scope}}",
	}
	stdOut, err := util.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}
	var networks []Network
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		networks = append(networks,
			Network{
				Name:   s[0],
				Driver: s[1],
				Scope:  s[2],
			})
	}

	return networks, nil
}
