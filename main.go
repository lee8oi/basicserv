/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

//
package main

import (
	"flag"
	"log"
	"net/http"
)

var httpPort = flag.String("http", "80", "http port")
var httpsPort = flag.String("https", "443", "https port")
var certPem = flag.String("cert", "certs/fullchain.pem", "path to cert pem file")
var keyPem = flag.String("key", "certs/privkey.pem", "path key pem file")
var domain = flag.String("domain", "localhost", "web domain")
var public = flag.String("pubdir", "public", "path to public directory")

func main() {
	flag.Parse()
	http.HandleFunc("/", indexServer)
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir(*public))))
	go func() {
		err := http.ListenAndServeTLS(":"+*httpsPort, *certPem, *keyPem, nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}()
	err := http.ListenAndServe(":"+*httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// indexServer serves the index.html.
func indexServer(w http.ResponseWriter, r *http.Request) {
	if !checkTLS(r) {
		http.Redirect(w, r, getHttpsUrl(), 301)
	}
	http.ServeFile(w, r, *public+"/index.html")
}

// getHttpsUrl returns the https url used for redirecting from http.
func getHttpsUrl() string {
	if *httpsPort == "443" {
		return "https://" + *domain
	}
	return "https://" + *domain + ":" + *httpsPort
}

// checkTLS returns true if TLS handshake is complete or false if not.
func checkTLS(r *http.Request) bool {
	if r.TLS != nil && r.TLS.HandshakeComplete {
		return true
	}
	return false
}
