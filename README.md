# basicserv
Basic web server written in Go (includes TLS).

# Usage
Build the server binary.
```Go
go build
```
Start the server.
```Go
basicserv -domain "mydomain.com" -cert "certdir/fullchain.pem" -key "keydir/privkey.pem"
```
Visit your web address in your web browser to test. Then hack the code to suit your website needs.