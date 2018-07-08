package main

import (
	"log"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/handlers"
)

func main() {
	cmd.CheckNotifications()

	log.Println("Listening http://0.0.0.0:8080")
	err := handlers.HTTPServer().ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
