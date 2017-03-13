# BasicServ
A simple web server written in Go. Intended as a basic starting point for building a fast & secure web host.
* Demonstrates the basic use of Go's http package including TLS.
* Automatically redirects to https to secure web traffic.
* Uses http FileServer handler to host website files.

# Usage

##### Generate SSL certificates.
On Linux/Unix systems the simplest way to obtain your SSL certificate & key file is by using Certbot
[https://certbot.eff.org/](https://certbot.eff.org/). You'll need the fullchain.pem and the privkey.pem files.

##### Build the server binary.
```
go build
```

##### Start the server.
(skip the -http & -https switches to use standard ports 80 & 443)
```
basicserv -domain "mydomain.com" -cert "certdir/fullchain.pem" -key "keydir/privkey.pem" -http 8080 -https 8081 -pubdir "path/to/pub"
```

##### Test the server.
Visit your web address in a web browser to test the server and verify your site is using a secure connection.

##### Modify the code.
Once you have successfully tested the server you are ready to edit the code and build a website to suit your needs. Good luck and have fun!