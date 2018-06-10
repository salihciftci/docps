package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

func generatePassword(l int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// dockerCmd runs docker commands and reads standart output line by line
func dockerCmd(cmdArgs []string) ([]string, error) {
	var stdOut []string

	cmd := exec.Command("docker", cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return stdOut, nil
}

// containers runs docker ps -a
func container() ([]PS, error) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}\t{{.Ports}}",
	}
	stdOut, err := dockerCmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var container []PS
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		container = append(container,
			PS{
				Name:       s[0],
				Image:      s[1],
				Size:       s[2],
				RunningFor: s[3],
				Status:     s[4][:1],
				Ports:      s[5],
			})
	}

	return container, nil
}

// images runs docker image ls
func images() ([]Images, error) {
	cmdArgs := []string{
		"image",
		"ls",
		"--format",
		"{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}",
	}
	stdOut, err := dockerCmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var images []Images
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		images = append(images,
			Images{
				Repository: s[0],
				Tag:        s[1],
				Created:    s[2],
				Size:       s[3],
			})
	}
	return images, nil
}

// volumes runs docker volume ls
func volumes() ([]Volumes, error) {
	cmdArgs := []string{
		"volume",
		"ls",
		"--format",
		"{{.Driver}}\t{{.Name}}",
	}
	stdOut, err := dockerCmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var volumes []Volumes
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		volumes = append(volumes,
			Volumes{
				Driver: s[0],
				Name:   s[1],
			})
	}

	return volumes, nil
}

// stats runs docker stats --no-stream
func stats() ([]Stats, error) {
	cmdArgs := []string{
		"stats",
		"--no-stream",
		"--format",
		"{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}",
	}
	stdOut, err := dockerCmd(cmdArgs)

	if err != nil {
		return nil, err
	}
	var stats []Stats
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		stats = append(stats,
			Stats{
				Name:     s[0],
				CPU:      s[1],
				MemUsage: s[2],
				MemPerc:  s[3],
				NetIO:    s[4],
				BlockIO:  s[5],
			})
	}

	return stats, nil
}

// logs runs docker logs <containerName>
func logs(container []PS) ([]Logs, error) {
	logs := []Logs{}
	for i := 0; i < len(container); i++ {
		cmdArgs := []string{
			"logs",
			container[i].Name,
		}

		cLog, err := dockerCmd(cmdArgs)
		if err != nil {
			return nil, fmt.Errorf("Docker daemon is not running")
		}

		if len(cLog) > 0 {
			x := []string{}
			for k := len(cLog) - 1; k != -1; k-- {
				x = append(x, cLog[k])
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

	return logs, nil
}

// networks runs docker network ls
func networks() ([]Networks, error) {
	cmdArgs := []string{
		"network",
		"ls",
		"--format",
		"{{.Name}}\t{{.Driver}}\t{{.Scope}}",
	}
	stdOut, err := dockerCmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}
	var networks []Networks
	for i := 0; i < len(stdOut); i++ {
		s := strings.Split(stdOut[i], "\t")
		networks = append(networks,
			Networks{
				Name:   s[0],
				Driver: s[1],
				Scope:  s[2],
			})
	}

	return networks, nil
}

// dashboard runs docker info and fetch docker infos for dashboard.
func dashboard() ([]interface{}, error) {
	cmdArgs := []string{
		"info",
		"--format",
		"{{.Containers}}\t{{.Name}}\t{{.ServerVersion}}\t{{.NCPU}}\t{{.MemTotal}}",
	}

	stdOut, err := dockerCmd(cmdArgs)

	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	var dashboard []interface{}

	s := strings.Split(stdOut[0], "\t")
	for i := 0; i < len(s); i++ {
		dashboard = append(dashboard, s[i])
	}

	images, err := images()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	volumes, err := volumes()
	if err != nil {
		return nil, fmt.Errorf("Docker daemon is not running")
	}

	networks, err := networks()
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

	notifiClear, notifiReverse := getNotification()

	dashboard = append(dashboard, notifiClear)
	dashboard = append(dashboard, notifiReverse)

	return dashboard, nil
}

func checkContainerStatus() ([]PS, error) {
	cmdArgs := []string{
		"ps",
		"-a",
		"--format",
		"{{.Names}}\t{{.Status}}",
	}

	stdOut, err := dockerCmd(cmdArgs)
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

	return container, nil
}

func getNotification() ([]notification, []notification) {
	var notifiReverse []notification
	for i := len(notifi) - 1; i >= 0; i-- {
		notifiReverse = append(notifiReverse, notifi[i])
	}

	var notifiClear []notification
	if len(notifiReverse) > 3 {
		notifiClear = notifiReverse[:3]
	} else {
		notifiClear = notifiReverse[:]
	}

	return notifiClear, notifiReverse
}
