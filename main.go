package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

//PS docker ps -a
type PS struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	Size       string `json:"size"`
	RunningFor string `json:"runningFor"`
	Status     string `json:"status"`
}

//Images docker image ls
type Images struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Created    string `json:"created"`
	Size       string `json:"size"`
}

//Volumes docker volume ls
type Volumes struct {
	Driver string `json:"driver"`
	Name   string `json:"name"`
}

//Stats docker stats --no-stream
type Stats struct {
	Name     string `json:"name"`
	CPU      string `json:"cpu"`
	MemUsage string `json:"memUsage"`
	MemPerc  string `json:"memPerc"`
	NetIO    string `json:"netIO"`
	BlockIO  string `json:"blockIO"`
}

func getStats(cmdArgs []string) []Stats {
	var stats []Stats

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

			stats = append(stats,
				Stats{Name: s[0],
					CPU:      s[1],
					MemUsage: s[2],
					MemPerc:  s[3],
					NetIO:    s[4],
					BlockIO:  s[5],
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

	return stats
}

func getVolumes(cmdArgs []string) []Volumes {
	var volumes []Volumes

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

			volumes = append(volumes,
				Volumes{Driver: s[0],
					Name: s[1],
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

	return volumes
}

func getImages(cmdArgs []string) []Images {
	var images []Images

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

			images = append(images,
				Images{Repository: s[0],
					Tag:     s[1],
					Created: s[2],
					Size:    s[3],
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

	return images

}

func ps(cmdArgs []string) []PS {
	var container []PS

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
				PS{Name: s[0],
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

//IndexHandler writing all outPuts to http template
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}",
	}
	container := ps(cmdArgs)

	cmdArgs = []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	images := getImages(cmdArgs)

	cmdArgs = []string{
		"volume",
		"ls",
		"--format",
		"{{.Driver}}\t{{.Name}}",
	}
	volumes := getVolumes(cmdArgs)

	cmdArgs = []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}
	stats := getStats(cmdArgs)

	var out []interface{}
	out = append(out, container)
	out = append(out, images)
	out = append(out, volumes)
	out = append(out, stats)

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, out)
	if err != nil {
		log.Println(err)
	}

}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", IndexHandler)
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return r
}

func main() {
	log.Println("Listening:8080..")

	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}

}
