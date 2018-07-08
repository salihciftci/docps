package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/salihciftci/liman/util"
)

var (
	notifications []Notification
)

//Notification asd
type Notification struct {
	Desc   string
	Time   string
	Status string
}

//ParseContainerStatus asd
func ParseContainerStatus() ([]PS, error) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Status}}",
	}

	stdOut, err := util.Cmd(cmdArgs)
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
	for i := len(notifications) - 1; i >= 0; i-- {
		reverse = append(reverse, notifications[i])
	}

	var basic []Notification
	if len(reverse) > 3 {
		basic = reverse[:3]
	} else {
		basic = reverse[:]
	}

	return basic, reverse
}

//CheckNotifications asd
func CheckNotifications() {
	// Checking containers for sending notification
	sc, err := ParseContainerStatus()
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			ps, err := ParseContainerStatus()
			if err != nil {
				log.Println(err)
			}

			if len(ps) != len(sc) {
				sc = ps
				continue
			}

			for i, v := range sc {
				if v.Status != ps[i].Status {
					if sc[i].Status == "U" {
						notifications = append(notifications, Notification{
							Desc:   sc[i].Name + " has stopped.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "E",
						})
					}

					if sc[i].Status == "E" {
						notifications = append(notifications, Notification{
							Desc:   sc[i].Name + " has started.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "U",
						})
					}
				}
			}
			sc = ps
			time.Sleep(5 * time.Second)
		}
	}()
}
