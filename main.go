// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/db/sqlite"
	"github.com/salihciftci/liman/handlers"
)

func main() {
	if sqlite.IsInstalled() {
		handlers.IsInstalled = true
	}

	handlers.Version = "v0.7"
	handlers.BaseURL = os.Getenv("URL")

	err := cmd.CheckNotifications()
	if err != nil {
		log.Println(err)
	}

	log.Println("Listening http://0.0.0.0:8080")
	err = handlers.HTTPServer().ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
