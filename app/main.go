package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

type httpRequest struct {
	method  string
	path    string
	headers map[string]string
	body    string
}

func main() {
	var dir string
	flag.StringVar(&dir, "directory", "", "Directory to serve")
	flag.Parse()
	if dir != "" {
		fmt.Println("Serving directory:", dir)
	}

	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("Listening on port 4221")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		go handleConn(conn, dir)
	}
}

func handleConn(conn net.Conn, dir string) {
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
		respondMethodNotAllowed(conn)
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

func handlePostRequest(httpRequest *httpRequest, conn net.Conn, dir string) {
	if strings.HasPrefix(httpRequest.path, "/files") {
		saveFile(conn, httpRequest.path, dir, httpRequest.body)
	} else {
		respondNotFound(conn)
	}
}

func serveFile(conn net.Conn, path string, dir string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		respondNotFound(conn)
		return
	}
	contents := make([]byte, 1024)
	n, err := file.Read(contents)
	if err != nil {
		respondServerError(conn)
		return
	}
	respondOK(conn, string(contents[:n]))
}

func saveFile(conn net.Conn, path string, dir string, body string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		respondServerError(conn)
		return
	}
	n, err := file.Write([]byte(body))
	if err != nil {
		respondServerError(conn)
		return
	}
	respondCreated(conn, fmt.Sprintf("File %s created with %d bytes", fileName, n))
}

func gzipString(message string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(message)); err != nil {
		return ""
	}
	if err := gz.Close(); err != nil {
		return ""
	}
	return b.String()
}

func respondOK(conn net.Conn, message string) {
	responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(message), message)
	conn.Write([]byte(responce))
}

func respondCreated(conn net.Conn, message string) {
	responce := fmt.Sprintf("HTTP/1.1 201 Created\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(message), message)
	conn.Write([]byte(responce))
}

func respondNotFound(conn net.Conn) {
	responce := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("404 Not Found"), "404 Not Found")
	conn.Write([]byte(responce))
}

func respondMethodNotAllowed(conn net.Conn) {
	responce := fmt.Sprintf("HTTP/1.1 405 Method Not Allowed\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("405 Method Not Allowed"), "405 Method Not Allowed")
	conn.Write([]byte(responce))
}

func respondServerError(conn net.Conn) {
	responce := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("500 Internal Server Error"), "500 Internal Server Error")
	conn.Write([]byte(responce))
}
