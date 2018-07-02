package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/salihciftci/liman/pkg/tool"
)

var (
	username = "root"
)

func parseDashboard() ([]interface{}, error) {
	cmdArgs := []string{
		"info",
		"--format",
		"{{.Containers}}\t{{.Name}}\t{{.ServerVersion}}\t{{.NCPU}}\t{{.MemTotal}}",
	}

	stdOut, err := tool.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var dashboard []interface{}

	s := strings.Split(stdOut[0], "\t")
	for i := 0; i < len(s); i++ {
		dashboard = append(dashboard, s[i])
	}

	images, err := parseImages()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	volumes, err := parseVolumes()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	networks, err := parseNetworks()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	dashboard = append(dashboard, strconv.Itoa(len(images)))
	dashboard = append(dashboard, strconv.Itoa(len(volumes)))
	dashboard = append(dashboard, strconv.Itoa(len(networks)))

	intMemory, err := strconv.Atoi(dashboard[4].(string))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	floatMemory := float64(intMemory)
	GibMemory := ((floatMemory / 1024) / 1024) / 1024
	dashboard[4] = strconv.FormatFloat(GibMemory, 'f', 2, 64)

	dashboard[1] = strings.Title(dashboard[1].(string))

	basicNotification, notifications := getNotification()

	dashboard = append(dashboard, basicNotification)

	if len(notifications) > 3 {
		dashboard = append(dashboard, "0")
	} else {
		dashboard = append(dashboard, "")
	}

	return dashboard, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inputPass := r.FormValue("inputPassword")
		inputUser := r.FormValue("inputUser")
		if inputUser == username && inputPass == userPassword {
			cookie := &http.Cookie{
				Name:    "session",
				Value:   cookieValue,
				Path:    "/",
				Expires: time.Now().AddDate(2, 0, 0),
				MaxAge:  0,
			}
			http.SetCookie(w, cookie)
		}
	}

	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	d, err := parseDashboard()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	var data []interface{}
	data = append(data, d)

	err = tpl.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}

	log.Println(r.Method, r.URL.Path)
}
