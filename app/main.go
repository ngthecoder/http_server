package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Start a TCP server that listens on port 4221
	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("Listening on port 4221")

	// Accept a connection and handle it
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// Read from the connection
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	req := string(buff[:n])
	lines := strings.Split(req, "\r\n")
	reqLine := strings.Split(lines[0], " ")
	method := reqLine[0]
	path := reqLine[1]
	userAgent := "User-Agent not found"
	for _, line := range lines {
		if strings.HasPrefix(line, "User-Agent:") {
			userAgent = strings.TrimSpace(strings.TrimPrefix(line, "User-Agent:"))
			break
		}
	}

	// Write to the connection
	if method == "GET" {
		if strings.HasPrefix(path, "/echo") {
			message := strings.TrimPrefix(path, "/echo/")
			responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(message), message)
			conn.Write([]byte(responce))
		} else if path == "/user-agent" {
			responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(userAgent), userAgent)
			conn.Write([]byte(responce))
		} else if path == "/" {
			responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("Hello World"), "Hello World")
			conn.Write([]byte(responce))
		} else {
			responce := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("404 Not Found"), "404 Not Found")
			conn.Write([]byte(responce))
		}
	} else if method == "POST" {
		fmt.Println("POST request")
	}

}
