package main

import (
	"net"
	"strings"
)

func handleGetRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		serveFile(conn, httpRequest.path, dir)
	} else if strings.HasPrefix(httpRequest.path, "/echo") {
		if httpRequest.headers["Accept-Encoding"] == "gzip" {
			message := strings.TrimPrefix(httpRequest.path, "/echo/")
			encodedMessage := gzipString(message)
			respondOK(conn, encodedMessage)
		} else {
			message := strings.TrimPrefix(httpRequest.path, "/echo/")
			respondOK(conn, message)
		}
	} else if httpRequest.path == "/user-agent" {
		respondOK(conn, httpRequest.headers["User-Agent"])
	} else if httpRequest.path == "/" {
		respondOK(conn, "Hello, World!")
	} else {
		respondNotFound(conn)
	}
}
