package notification

import (
	"strings"

	"github.com/salihciftci/liman/pkg/docker"
)

var (
	//Notifications asd
	Notifications []Notification
)

//Notification asd
type Notification struct {
	Desc   string
	Time   string
	Status string
}

//PS asd
type PS struct {
	Name       string `json:"name,omitempty"`
	Image      string `json:"image,omitempty"`
	Size       string `json:"size,omitempty"`
	RunningFor string `json:"runningFor,omitempty"`
	Status     string `json:"status,omitempty"`
	Ports      string `json:"ports,omitempty"`
}

//ParseContainerStatus asd
func ParseContainerStatus() ([]PS, error) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Status}}",
	}

	stdOut, err := docker.Cmd(cmdArgs)
	if err != nil {
		return nil, err
	}

	var container []PS

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			PS{
				Name:   s[0],
				Status: s[1][:1],
			})
	}

	var reverse []PS

	for i := len(container) - 1; i != -1; i-- {
		reverse = append(reverse, container[i])
	}

	return reverse, nil
}

//GetNotification asd
func GetNotification() ([]Notification, []Notification) {
	var reverse []Notification
	for i := len(Notifications) - 1; i >= 0; i-- {
		reverse = append(reverse, Notifications[i])
	}

	var basic []Notification
	if len(reverse) > 3 {
		basic = reverse[:3]
	} else {
		basic = reverse[:]
	}

	return basic, reverse
}
