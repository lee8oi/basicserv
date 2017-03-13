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
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*pubdir))))
	go func() {
		err := http.ListenAndServeTLS(":"+*httpsPort, *certPem, *keyPem, nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}()
	err := http.ListenAndServe(":"+*httpPort, http.RedirectHandler("https://"+*domain+":"+*httpsPort, 301))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
