/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

//
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	cfg := loadConfig(*cfgPath)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(cfg.PubDir))))
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

var cfgPath = flag.String("config", "config.json", "path to config file (in JSON format)")

// config type contains the necessary server configuration strings.
type config struct {
	HTTPPort, HTTPSPort,
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
