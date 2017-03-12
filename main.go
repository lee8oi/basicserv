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
var pubdir = flag.String("pubdir", "public", "path to public directory")

func main() {
	flag.Parse()
	http.HandleFunc("/", indexServer)
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir(*pubdir))))
	go log.Fatal(http.ListenAndServeTLS(":"+*httpsPort, *certPem, *keyPem, nil))
	log.Fatal(http.ListenAndServe(":"+*httpPort, nil))
}

// indexServer serves the index.html.
func indexServer(w http.ResponseWriter, r *http.Request) {
	redirect(w, r)
	http.ServeFile(w, r, *pubdir+"/index.html")
}

// redirect will reroute to https url if TLS not in use.
func redirect(w http.ResponseWriter, r *http.Request) {
	if r.TLS == nil {
		http.Redirect(w, r, func() string {
			if *httpsPort == "443" {
				return "https://" + *domain
			}
			return "https://" + *domain + ":" + *httpsPort
		}(), 301)
	}
}
