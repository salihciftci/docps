// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"strings"
	"time"

	"github.com/salihciftci/liman/util"
)

var (
	notifications []Notification
)

//Notification stores values for notifications
type Notification struct {
	Desc   string
	Time   string
	Status string
}

//ParseContainerStatus is parses and splits containers process state from output for containers status changes
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

//GetNotification is revers notifications
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

//CheckNotifications is checking containers status for notifications
func CheckNotifications() error {
	sc, err := ParseContainerStatus()
	if err != nil {
		return err
	}

	go func() error {
		for {
			ps, err := ParseContainerStatus()
			if err != nil {
				return err
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

	return nil
}
