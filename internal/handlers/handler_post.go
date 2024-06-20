package internal

import (
	"net"
	"strings"

	"github.com/ngthecoder/http_server/internal/responces"
)

func handlePostRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		saveFile(conn, httpRequest.path, dir, httpRequest.body)
	} else {
		responces.RespondNotFound(conn)
	}
}
