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

var cfgPath = flag.String("config", "config.json", "path to config file (in JSON format)")

func main() {
	cfg := loadConfig(*cfgPath)
	flag.Parse()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(cfg.PubDir))))
	go func() {
		err := http.ListenAndServeTLS(":"+cfg.HttpsPort, cfg.CertPem, cfg.KeyPem, nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}()
	err := http.ListenAndServe(":"+cfg.HttpPort, http.RedirectHandler("https://"+cfg.Domain+":"+cfg.HttpsPort, 301))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type config struct {
	HttpPort, HttpsPort,
	Domain, PubDir, CertPem, KeyPem string
}

func loadConfig(path string) (i config) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		log.Fatal(err)
	}
	return
}
