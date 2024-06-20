package internal

import (
	"fmt"
	"net"
	"strings"

	responce "github.com/ngthecoder/http_server/internal/responce"
)

type httpRequest struct {
	method  string
	path    string
	headers map[string]string
	body    string
}

func HandleConn(conn net.Conn, dir string) {
	defer conn.Close()

	httpRequest, err := readRequest(conn)
	if err != nil {
		fmt.Println("Error reading request:", err.Error())
		return
	}

	switch httpRequest.method {
	case "GET":
		handleGetRequest(httpRequest, conn, dir)
	case "POST":
		handlePostRequest(httpRequest, conn, dir)
	default:
		responce.Respond(conn, 405, "Method Not Allowed")
	}
}

func readRequest(conn net.Conn) (*httpRequest, error) {
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		return nil, err
	}

	req := string(buff[:n])
	lines := strings.Split(req, "\r\n")
	reqLine := strings.Split(lines[0], " ")
	headers := parseHeaders(lines[1:])
	body := lines[len(lines)-1]

	return &httpRequest{
		method:  reqLine[0],
		path:    reqLine[1],
		headers: headers,
		body:    body,
	}, nil
}

func parseHeaders(headerLines []string) map[string]string {
	headers := make(map[string]string)
	for _, line := range headerLines {
		if parts := strings.SplitN(line, ": ", 2); len(parts) == 2 {
			headers[parts[0]] = parts[1]
		}
	}
	return headers
}
