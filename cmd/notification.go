package cmd

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/salihciftci/liman/util"
)

var (
	notifications []notification
)

type notification struct {
	Desc   string
	Time   string
	Status string
}

func parseContainerStatus() ([]ps, error) {
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

	var container []ps

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			ps{
				Name:   s[0],
				Status: s[1][:1],
			})
	}

	var reverse []ps

	for i := len(container) - 1; i != -1; i-- {
		reverse = append(reverse, container[i])
	}

	return reverse, nil
}

func getNotification() ([]notification, []notification) {
	var reverse []notification
	for i := len(notifications) - 1; i >= 0; i-- {
		reverse = append(reverse, notifications[i])
	}

	var basic []notification
	if len(reverse) > 3 {
		basic = reverse[:3]
	} else {
		basic = reverse[:]
	}

	return basic, reverse
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}
	bn, n := getNotification()

	if len(n) > 100 {
		n = n[:100]
	}

	var data []interface{}

	data = append(data, bn)
	data = append(data, n)

	err = tpl.ExecuteTemplate(w, "notifications.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)

}

func checkNotifications() {
	// Checking containers for sending notification
	sc, err := parseContainerStatus()
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			ps, err := parseContainerStatus()
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
						notifications = append(notifications, notification{
							Desc:   sc[i].Name + " has stopped.",
							Time:   time.Now().Format("02/01/2006 15:04"),
							Status: "E",
						})
					}

					if sc[i].Status == "E" {
						notifications = append(notifications, notification{
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
