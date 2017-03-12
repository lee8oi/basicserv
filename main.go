/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

//
package main

import (
	"flag"
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

var httpPort = flag.String("http", "80", "http port")
var httpsPort = flag.String("https", "443", "https port")
var certPem = flag.String("cert", "certs/fullchain.pem", "path to cert pem file")
var keyPem = flag.String("key", "certs/privkey.pem", "path key pem file")
var domain = flag.String("domain", "localhost", "web domain")

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(homeServer))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	go func() {
		// Certs generated by letsencrypt software
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

var homeTemplate = template.Must(template.ParseFiles("index.html"))

// homeServer serves the home page to the requesting browser.
func homeServer(w http.ResponseWriter, r *http.Request) {
	if !checkTLS(r) {
		http.Redirect(w, r, redirectUrl(), 301)
	}
	type data struct {
		Status string
	}
	d := data{
		Status: func() string {
			if checkTLS(r) {
				return "SECURED"
			} else {
				return "UNSECURED"
			}
		}(),
	}
	homeTemplate.Execute(w, d)
}

// redirectUrl assembles the https url for redirect to TLS.
func redirectUrl() string {
	rUrl := "https://"
	if *httpsPort == "443" {
		rUrl = rUrl + *domain
	} else {
		rUrl = rUrl + *domain + ":" + *httpsPort
	}
	return rUrl
}

// checkTLS returns true if TLS handshake is complete or false if not.
func checkTLS(r *http.Request) bool {
	if r.TLS != nil && r.TLS.HandshakeComplete {
		return true
	}
	return false
}
