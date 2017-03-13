# BasicServ
A simple web server written in Go. Intended as a basic starting point for building a fast & secure web host.
* Demonstrates the basic use of Go's http package including TLS.
* Automatically redirects to https to secure web traffic.
* Uses http FileServer handler to serve files.
* Simple server configuration via json file.

# Usage

##### Generate SSL certificates.
On Linux/Unix systems the simplest way to obtain your SSL certificate & key file is by using Certbot
[https://certbot.eff.org/](https://certbot.eff.org/). You'll need the fullchain.pem and the privkey.pem files.

##### Build the server binary.
```Bash
$ go build
```

##### Edit server configuration.
Rename `config.json.example` file to `config.json` and edit the configuration to suit your needs.

##### Start the server.
```Bash
$ basicserv -config "path/to/config.json"
```

##### Test the server.
Visit your web address in a web browser to test the server and verify your site is using a secure connection.

##### Modify the code.
Once you have successfully tested the server you are ready to edit the code and build a website to suit your needs. Good luck and have fun!