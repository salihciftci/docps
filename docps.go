package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Docker struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	Size       string `json:"size"`
	RunningFor string `json:"runningFor"`
	Status     string `json:"status"`
}

//IndexHandler Execute the docker ps -a command and reading the stdout
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	cmdArgs := []string{"ps", "-a", "--format", "{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}"}
	container := ps(cmdArgs)

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, container)
	if err != nil {
		log.Println(err)
	}

}

func ps(cmdArgs []string) []Docker {
	var container []Docker

	cmd := exec.Command("docker", cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			outPut := scanner.Text()

			s := strings.Split(outPut, "\t")

			container = append(container,
				Docker{Name: s[0],
					Image:      s[1],
					Size:       s[2],
					RunningFor: s[3],
					Status:     s[4][:1],
				})
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return container

}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", IndexHandler)
	return r
}

func main() {
	log.Println("Listening:8080..")

	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}

}
