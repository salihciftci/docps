package liman

import (
	"log"
	"net/http"
	"strings"

	"github.com/salihciftci/liman/pkg/tool"
)

type stat struct {
	Name     string `json:"name,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	MemUsage string `json:"memUsage,omitempty"`
	MemPerc  string `json:"memPerc,omitempty"`
	NetIO    string `json:"netIO,omitempty"`
	BlockIO  string `json:"blockIO,omitempty"`
}

func parseStats() ([]stat, error) {
	cmdArgs := []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}
	stdOut, err := tool.Cmd(cmdArgs)

	if err != nil {
		return nil, err
	}
	var stats []stat
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		stats = append(stats,
			stat{
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

func statsHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	s, err := parseStats()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := getNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, s)

	err = tpl.ExecuteTemplate(w, "stats.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}
