# BasicServ
A basic starter web server written in Go. Designed to automatically redirect to https for better security.

# Usage

#### Generate SSL certificates.
On Linux/Unix systems the simplest way to obtain your SSL certificate & key file is by using Certbot
[https://certbot.eff.org/](https://certbot.eff.org/). You'll need the fullchain.pem and the privkey.pem files.

#### Build the server binary.
```
go build
```

#### Start the server.
```
basicserv -domain "mydomain.com" -cert "certdir/fullchain.pem" -key "keydir/privkey.pem"
```

#### Test the server.
Visit your web address in a web browser to test the server and verify your site is using a secure connection.

#### Modify the code.
Once you have successfully tested the server you are ready to edit the code and build a website to suit your needs. Good luck and have fun!