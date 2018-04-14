package main

import (
	"bufio"
	"log"
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

func read(cmdArgs []string) []string {
	var stdOut []string

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

			stdOut = append(stdOut, outPut)
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

	return stdOut
}

func getDocker() []interface{} {
	var container []PS
	var images []Images
	var volumes []Volumes
	var stats []Stats

	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}",
	}
	stdOut := read(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			PS{Name: s[0],
				Image:      s[1],
				Size:       s[2],
				RunningFor: s[3],
				Status:     s[4][:1],
			})
	}

	cmdArgs = []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	stdOut = read(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		images = append(images,
			Images{Repository: s[0],
				Tag:     s[1],
				Created: s[2],
				Size:    s[3],
			})
	}

	cmdArgs = []string{
		"volume",
		"ls",
		"--format",
		"{{.Driver}}\t{{.Name}}",
	}
	stdOut = read(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		volumes = append(volumes,
			Volumes{Driver: s[0],
				Name: s[1],
			})
	}

	cmdArgs = []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}
	stdOut = read(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		stats = append(stats,
			Stats{Name: s[0],
				CPU:      s[1],
				MemUsage: s[2],
				MemPerc:  s[3],
				NetIO:    s[4],
				BlockIO:  s[5],
			})
	}

	var out []interface{}
	out = append(out, container)
	out = append(out, images)
	out = append(out, volumes)
	out = append(out, stats)

	return out
}
