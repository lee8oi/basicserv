/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

//
package main

import (
	"crypto/md5"
	"crypto/subtle"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var cfgPath = flag.String("config", "config.json", "path to config file (in JSON format)")

func main() {
	flag.Parse()
	cfg := loadConfig(*cfgPath)
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(cfg.PubDir))))
	http.HandleFunc("/", authHandler(fileServer(cfg.PubDir),
		hasher(cfg.User), hasher(cfg.Pass), "Please enter your username and password"))
	go func() {
		err := http.ListenAndServeTLS(":"+cfg.HTTPSPort, cfg.CertPem, cfg.KeyPem, nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}()
	err := http.ListenAndServe(":"+cfg.HTTPPort, http.RedirectHandler("https://"+cfg.Domain+":"+cfg.HTTPSPort, 301))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// hasher hashes the given string and returns the sum as a slice of bytes.
func hasher(s string) []byte {
	val := md5.Sum([]byte(s))
	return val[:]
}

// config type contains the necessary server configuration strings.
type config struct {
	HTTPPort, HTTPSPort, User, Pass,
	Domain, PubDir, CertPem, KeyPem string
}

// loadConfig loads configuration values from file.
func loadConfig(path string) (c config) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// fileServer returns a root ("/") FileServer handler function for the given directory.
func fileServer(dir string) http.HandlerFunc {
	h := http.FileServer(http.Dir(dir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
		h.ServeHTTP(w, r)
	})
}

// authHandler wraps a handler function to provide http basic authentication.
func authHandler(handler http.HandlerFunc, username, password []byte, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		userByt, passByt := hasher(user), hasher(pass)
		if !ok || subtle.ConstantTimeCompare(userByt,
			username) != 1 || subtle.ConstantTimeCompare(passByt, password) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

// // authHandler wraps a handler function to provide http basic authentication.
// func authHandler(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		u, p, ok := r.BasicAuth()
// 		if !ok || u != username || p != password {
// 			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("401 Unauthorized\n"))
// 			return
// 		}
// 		handler(w, r)
// 	}
// }
