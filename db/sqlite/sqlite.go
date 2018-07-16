// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqlite

import (
	"database/sql"
	"log"
	"os"
	"time"

	//_ Sqlite3 Driver
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

//Connect Checks database if it exist Connect to db if not create a db folder and db then, connect it.
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

	return db, nil
}

// IsInstalled Checks Liman already installed or not.
func IsInstalled() bool {
	db, err := Connect()
	if err != nil {
		return false
	}

	var isInstalled bool
	err = db.QueryRow("SELECT key FROM config WHERE value = ?", "isInstalled").Scan(&isInstalled)
	if err != nil {
		return false
	}

	return isInstalled
}

//Install creates database, tables and insert configs
func Install(pass, sessionKey, apiKey, ver string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	s, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user TEXT,
			pass TEXT,
			sessionKey TEXT,
			permission TEXT,
			desc TEXT,
			created TEXT,
			updated TEXT)
	`)

	if err != nil {
		return err
	}

	s.Exec()

	stmt, err := db.Prepare(`INSERT INTO users
		(user, pass, sessionKey, permission, desc, created, updated)
		VALUES (?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	created := time.Now().Format("02/01/2006 15:04")
	updated := time.Now().Format("02/01/2006 15:04")

	stmt.Exec("root", pass, sessionKey, "root", "root user", created, updated)

	s, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS config (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			value TEXT,
			key TEXT)
	`)

	if err != nil {
		return err
	}

	s.Exec()

	stmt, err = db.Prepare(`INSERT INTO config
		(value,key) VALUES
		(?,?),
		(?,?),
		(?,?)
	`)

	if err != nil {
		return err
	}

	stmt.Exec("isInstalled", "true", "apiKey", apiKey, "version", ver)

	return nil
}

//ParseAPIKey parses api key from db
func ParseAPIKey() (string, error) {
	db, err := Connect()
	if err != nil {
		return "", err
	}

	var APIKey string
	err = db.QueryRow("SELECT key FROM config WHERE value = ?", "apiKey").Scan(&APIKey)

	if err != nil {
		return "", err
	}

	return APIKey, nil
}

// GetUserPasswordAndSessionKey parses user pass and sessionKey
func GetUserPasswordAndSessionKey(user string) (string, string, error) {
	db, err := Connect()
	if err != nil {
		return "", "", err
	}

	var hash string
	var key string
	err = db.QueryRow("SELECT pass, sessionKey FROM users WHERE user = ?", user).Scan(&hash, &key)
	if err != nil {
		return "", "", err
	}
	return hash, key, err
}

// GetUserFromSessionKey parses users from sessionKey
func GetUserFromSessionKey(key string) (string, error) {
	db, err := Connect()
	if err != nil {
		return "", err
	}

	var user string
	err = db.QueryRow("SELECT user FROM users WHERE sessionKey = ?", key).Scan(&user)

	if err != nil {
		return "", err
	}

	return user, nil
}

//ChangeUserPassword changes user password from db
func ChangeUserPassword(user, pass string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE users SET pass = ? WHERE user = ?")
	if err != nil {
		return err
	}

	stmt.Exec(pass, user)

	return nil
}

//ParseVersion parses version from db
func ParseVersion() (string, error) {
	db, err := Connect()
	if err != nil {
		return "", err
	}

	var version string

	err = db.QueryRow("SELECT key FROM config WHERE value = ?", "version").Scan(&version)

	if err != nil {
		return "", err
	}

	return version, nil
}
