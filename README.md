# smolweb

Small, hackable, web server. 

Will attempt to serve files located in the current working directory of the server. 

Demonstrates redirection and HTTP header manipulation. 

Written in [Go](https://golang.org/). 

## Run

	go run webserver.go

Listens by default at <http://localhost:1337/>. 

## Usage

```
Usage of webserver:
  -cert string
        Certificate file for HTTPS (default "cert.pem")
  -key string
        Private key file for HTTPS (default "key.pem")
  -port string
        Listening port for HTTP server (default ":1337")
  -to string
        Path to redirect (default "foo.crm.dynamics.com/bar")
```

## Examples

Generate an certificate/key pair and serve HTTPS:

```
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
$ smolweb -port :443 -cert cert.pem -key key.pem
```
