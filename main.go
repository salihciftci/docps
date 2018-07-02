package main

import (
	"log"

	"github.com/salihciftci/liman/cmd/liman"
)

func main() {
	err := liman.Run()
	if err != nil {
		log.Println(err)
	}
}
