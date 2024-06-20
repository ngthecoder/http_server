package internal

import (
	"net"
	"strings"

	middleware "github.com/ngthecoder/http_server/internal/middleware"
	responce "github.com/ngthecoder/http_server/internal/responce"
)

func handleGetRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		serveFile(conn, httpRequest.path, dir)
	} else if strings.HasPrefix(httpRequest.path, "/echo") {
		if httpRequest.headers["Accept-Encoding"] == "gzip" {
			message := strings.TrimPrefix(httpRequest.path, "/echo/")
			encodedMessage := middleware.GzipString(message)
			responce.Respond(conn, 200, encodedMessage)
		} else {
			message := strings.TrimPrefix(httpRequest.path, "/echo/")
			responce.Respond(conn, 200, message)
		}
	} else if httpRequest.path == "/user-agent" {
		responce.Respond(conn, 200, httpRequest.headers["User-Agent"])
	} else if httpRequest.path == "/" {
		responce.Respond(conn, 200, "OK")
	} else {
		responce.Respond(conn, 404, "Not Found")
	}
}
