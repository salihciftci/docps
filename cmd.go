package main

import (
	"bufio"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

//PS docker ps -a
type PS struct {
	Name       string `json:"name,omitempty"`
	Image      string `json:"image,omitempty"`
	Size       string `json:"size,omitempty"`
	RunningFor string `json:"runningFor,omitempty"`
	Status     string `json:"status,omitempty"`
	Ports      string `json:"ports,omitempty"`
}

//Images docker image ls
type Images struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Created    string `json:"created,omitempty"`
	Size       string `json:"size,omitempty"`
}

//Volumes docker volume ls
type Volumes struct {
	Driver string `json:"driver,omitempty"`
	Name   string `json:"name,omitempty"`
}

//Stats docker stats --no-stream
type Stats struct {
	Name     string `json:"name,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	MemUsage string `json:"memUsage,omitempty"`
	MemPerc  string `json:"memPerc,omitempty"`
	NetIO    string `json:"netIO,omitempty"`
	BlockIO  string `json:"blockIO,omitempty"`
}

//Logs docker logs <name>
type Logs struct {
	Name string   `json:"name,omitempty"`
	Logs []string `json:"logs,omitempty"`
}

//Networks docker network ls
type Networks struct {
	Name   string `json:"name,omitempty"`
	Driver string `json:"driver,omitempty"`
	Scope  string `json:"scope,omitempty"`
}

func dockerCmd(cmdArgs []string) []string {
	var stdOut []string

	cmd := exec.Command("docker", cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return nil
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		return nil
	}

	return stdOut
}

func getDocker() []interface{} {
	var data []interface{}

	var container []PS
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}\t{{.Ports}}",
	}
	stdOut := dockerCmd(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			PS{Name: s[0],
				Image:      s[1],
				Size:       s[2],
				RunningFor: s[3],
				Status:     s[4][:1],
				Ports:      s[5],
			})
	}
	data = append(data, container)

	var images []Images
	cmdArgs = []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	stdOut = dockerCmd(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		images = append(images,
			Images{Repository: s[0],
				Tag:     s[1],
				Created: s[2],
				Size:    s[3],
			})
	}
	data = append(data, images)

	var volumes []Volumes
	cmdArgs = []string{
		"volume",
		"ls",
		"--format",
		"{{.Driver}}\t{{.Name}}",
	}
	stdOut = dockerCmd(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		volumes = append(volumes,
			Volumes{Driver: s[0],
				Name: s[1],
			})
	}
	data = append(data, volumes)

	var stats []Stats
	cmdArgs = []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}
	stdOut = dockerCmd(cmdArgs)

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
	data = append(data, stats)

	logs := []Logs{}
	for i := 0; i < len(container); i++ {
		cmdArgs = []string{
			"logs",
			container[i].Name,
		}

		log := dockerCmd(cmdArgs)
		if len(log) > 0 {
			x := []string{}
			for k := 0; k < len(log); k++ {
				x = append(x, log[k])
			}
			logs = append(logs, Logs{
				Name: container[i].Name,
				Logs: x,
			})

		} else {
			logs = append(logs, Logs{
				Name: container[i].Name,
				Logs: []string{},
			})
		}
	}
	data = append(data, logs)

	var networks []Networks
	cmdArgs = []string{
		"network",
		"ls",
		"--format",
		"{{.Name}}\t{{.Driver}}\t{{.Scope}}",
	}
	stdOut = dockerCmd(cmdArgs)

	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		networks = append(networks,
			Networks{
				Name:   s[0],
				Driver: s[1],
				Scope:  s[2],
			})
	}
	data = append(data, networks)

	cmdArgs = []string{
		"info",
		"--format",
		"{{.ContainersRunning}}\t{{.ContainersPaused}}\t{{.ContainersStopped}}\t{{.Name}}\t{{.ServerVersion}}\t{{.NCPU}}\t{{.MemTotal}}",
	}

	stdOut = dockerCmd(cmdArgs)

	dashboard := []string{}

	s := strings.Split(stdOut[0], "\t")
	for i := 0; i < len(s); i++ {
		dashboard = append(dashboard, s[i])
	}

	dashboard = append(dashboard, strconv.Itoa(len(images)))
	dashboard = append(dashboard, strconv.Itoa(len(volumes)))
	dashboard = append(dashboard, strconv.Itoa(len(networks)))

	intMemory, err := strconv.Atoi(dashboard[6])
	if err != nil {
		log.Println(err)
	}

	floatMemory := float64(intMemory)
	GibMemory := ((floatMemory / 1024) / 1024) / 1024
	dashboard[6] = strconv.FormatFloat(GibMemory, 'f', 2, 64)

	data = append(data, dashboard)

	return data
}
