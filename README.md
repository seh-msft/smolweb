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
  -p string
        Listening port for HTTP server (default ":1337")
  -to string
        Path to redirect (default "foo.crm.dynamics.com/bar")
```
