package main

import (
	"log"

	"github.com/salihciftci/liman/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
