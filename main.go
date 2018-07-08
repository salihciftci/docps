package main

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/handlers"
)

func main() {
	//Checking Notifications every 5 second
	cmd.CheckNotifications()

	log.Println("Listening http://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", handlers.Handler())
	if err != nil {
		panic(err)
	}
}
