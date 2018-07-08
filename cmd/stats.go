package cmd

import (
	"strings"

	"github.com/salihciftci/liman/util"
)

//Stat asd
type Stat struct {
	Name     string `json:"name,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	MemUsage string `json:"memUsage,omitempty"`
	MemPerc  string `json:"memPerc,omitempty"`
	NetIO    string `json:"netIO,omitempty"`
	BlockIO  string `json:"blockIO,omitempty"`
}

//ParseStats asd
func ParseStats() ([]Stat, error) {
	cmdArgs := []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}
	stdOut, err := util.Cmd(cmdArgs)

	if err != nil {
		return nil, err
	}
	var stats []Stat
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		stats = append(stats,
			Stat{
				Name:     s[0],
				CPU:      s[1],
				MemUsage: s[2],
				MemPerc:  s[3],
				NetIO:    s[4],
				BlockIO:  s[5],
			})
	}

	return stats, nil
}
