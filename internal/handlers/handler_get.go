package internal

import (
	"net"
	"strings"

	middleware "github.com/ngthecoder/http_server/internal/middleware"
	responces "github.com/ngthecoder/http_server/internal/responces"
)

func handleGetRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		serveFile(conn, httpRequest.path, dir)
	} else if strings.HasPrefix(httpRequest.path, "/echo") {
		if httpRequest.headers["Accept-Encoding"] == "gzip" {
			message := strings.TrimPrefix(httpRequest.path, "/echo/")
			encodedMessage := middleware.GzipString(message)
			responces.RespondOK(conn, encodedMessage)
		} else {
			message := strings.TrimPrefix(httpRequest.path, "/echo/")
			responces.RespondOK(conn, message)
		}
	} else if httpRequest.path == "/user-agent" {
		responces.RespondOK(conn, httpRequest.headers["User-Agent"])
	} else if httpRequest.path == "/" {
		responces.RespondOK(conn, "Hello, World!")
	} else {
		responces.RespondNotFound(conn)
	}
}
