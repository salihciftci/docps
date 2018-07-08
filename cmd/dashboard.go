package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/salihciftci/liman/util"
)

//ParseDashboard asd
func ParseDashboard() ([]interface{}, error) {
	cmdArgs := []string{
		"info",
		"--format",
		"{{.Containers}}\t{{.Name}}\t{{.ServerVersion}}\t{{.NCPU}}\t{{.MemTotal}}",
	}

	stdOut, err := util.Cmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var dashboard []interface{}

	s := strings.Split(stdOut[0], "\t")
	for i := 0; i < len(s); i++ {
		dashboard = append(dashboard, s[i])
	}

	images, err := ParseImages()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	volumes, err := ParseVolumes()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	networks, err := ParseNetworks()
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

	basicNotification, notifications := GetNotification()

	dashboard = append(dashboard, basicNotification)

	if len(notifications) > 3 {
		dashboard = append(dashboard, "0")
	} else {
		dashboard = append(dashboard, "")
	}

	return dashboard, nil
}
