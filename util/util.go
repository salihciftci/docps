// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//GeneratePassword for apiKey and cookieValue
func GeneratePassword(l int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

//Version getting git commit id
func Version() (string, error) {
	var version string

	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			version = scanner.Text()
		}
	}()

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return version, nil
}

//ShortPorts Parsing containers ports and shoring them
func ShortPorts(p string) string {
	if len(p) < 9 {
		return p
	}

	ports := strings.Split(p, ", ")
	for i := range ports {
		if len(ports[i]) > 9 {
			ports[i] = ports[i][8:]
		}
	}

	p = strings.Join(ports, ", ")

	return p
}

//HashPasswordAndSave Hashing root password and storing them.
func HashPasswordAndSave(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	path, _ := os.Getwd()
	if err != nil {
		return "", err
	}

	err = os.Mkdir(path+"/data", 0777)
	if err != nil {
		log.Println("/data folder already exist. Skipping.")
	}

	err = ioutil.WriteFile(path+"/data/pass", b, 0644)
	if err != nil {
		return "", err
	}

	return string(b), nil

}

//ReadPassword read hashed password from file
func ReadPassword() string {
	path, _ := os.Getwd()
	h, err := ioutil.ReadFile(path + "/data/pass")
	if err != nil {
		return ""
	}

	return string(h)
}

//CheckPass matchs hash value with pass
func CheckPass(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
