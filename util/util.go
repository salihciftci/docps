// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateSecretKey generates a secret token for jwt
func GenerateSecretKey(l int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}

	return string(b)
}

//GenerateJWT generates a jwt key for authentication
func GenerateJWT(user string, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}

	return tokenString
}

// ParseJWT parses JWT key to pair info
func ParseJWT(jwtString, secret string) interface{} {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Something badly happend while parsing JWT")
		}

		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user"]
	}

	return err
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
