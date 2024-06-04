package main

import (
	"net"
	"strings"
)

func handlePostRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		saveFile(conn, httpRequest.path, dir, httpRequest.body)
	} else {
		respondNotFound(conn)
	}
}
