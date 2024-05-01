// Copyright (c) 2019, Microsoft Corporation, Sean Hinchee
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

var (
	port   = flag.String("port", ":1337", "Listening port for HTTP server")
	target = flag.String("to", "foo.crm.dynamics.com/bar", "Path to redirect")
	cert   = flag.String("cert", "", "Certificate file for HTTPS (cert.pem)")
	key    = flag.String("key", "", "Private key file for HTTPS (key.pem)")
)

// Small utility web server intended to be easy to hack on top of
func main() {
	flag.Parse()

	http.HandleFunc("/", rootHandler)

	if len(*cert) > 0 && len(*key) > 0 {
		// HTTPS
		log.Println("Listening on https://localhost" + *port + " ...")
		log.Fatal(http.ListenAndServeTLS(*port, *cert, *key, nil))
	} else {
		// Plain HTTP
		log.Println("Listening on http://localhost" + *port + " ...")
		log.Fatal(http.ListenAndServe(*port, nil))
	}

	log.Fatal(http.ListenAndServe(*port, nil))
}

// Handle requests to /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(prettyRequest(r))

	// 404 handler
	notFound := func() {
		w.WriteHeader(404)
		w.Write([]byte("404 - page not found, sorry ☹"))
	}

	// Redirect with 307
	redir := func() {
		url := "https://" + *target
		http.Redirect(w, r, url, 307)
	}

	page := r.URL.EscapedPath()
	if len(page) > 1 {
		// Strip leading /
		page = page[1:]
	}

	// Lazily disallow parenting
	page = strings.ReplaceAll(page, "..", "")

	if page == "/" {
		// index.html alias
		page = "index.html"
	}

	extension := filepath.Ext(page)

	// Content types
	switch extension {
	case ".html":
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
	case ".swf":
		w.Header().Add("Content-Type", "application/x-shockwave-flash")
	default:

	}

	// Individual page handling, you can hardcode logic here
	switch page {
	case "redirect.php":
		redir()
	default:
		f, err := os.Open(page)
		if err != nil {
			notFound()
		}
		defer f.Close()

		writer := bufio.NewWriter(w)
		writer.ReadFrom(f)
		writer.Flush()
	}
}

// Pretty logging of HTTP requests
func prettyRequest(r *http.Request) string {
	path := r.URL.Path

	// Drop logging for favicon spam
	if strings.Contains(path, "favicon.ico") {
		return ""
	}

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return ""
	}

	return fmt.Sprint("\n+++ " + path + " Request\n\n" + string(dump))
}
