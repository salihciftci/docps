// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"database/sql"
	"log"
	"os"

	//_ Sqlite3 Driver
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

//Connect connect to sqlite db
func Connect() (*sql.DB, error) {
	if _, err := os.Stat("data/liman.db"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			log.Println("Data folder already exist. Skipping.")
		}
	}

	db, err := sql.Open("sqlite3", "data/liman.db")
	if err != nil {
		return nil, err
	}

	s, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			user TEXT,
			pass TEXT,
			sessionKey TEXT,
			permission TEXT,
			desc TEXT,
			created TEXT,
			updated TEXT)
	`)

	s.Exec()

	return db, nil
}
