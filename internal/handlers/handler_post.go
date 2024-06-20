package internal

import (
	"net"
	"strings"

	responce "github.com/ngthecoder/http_server/internal/responce"
)

func handlePostRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		saveFile(conn, httpRequest.path, dir, httpRequest.body)
	} else {
		responce.Respond(conn, 404, "Not Found")
	}
}
