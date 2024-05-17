package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Parse command line arguments
	var dir string
	flag.StringVar(&dir, "directory", "", "Directory to serve")
	flag.Parse()
	if dir != "" {
		fmt.Println("Serving directory:", dir)
	}

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
		go handleConn(conn, dir)
	}
}

func handleConn(conn net.Conn, dir string) {
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
	body := lines[len(lines)-1]

	// Write to the connection
	if method == "GET" {
		if strings.HasPrefix(path, "/files") {
			serveFile(conn, path, dir)
		} else if strings.HasPrefix(path, "/echo") {
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
		saveFile(conn, path, dir, body)
	}

}

func serveFile(conn net.Conn, path string, dir string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		responce := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("404 Not Found"), "404 Not Found")
		conn.Write([]byte(responce))
		return
	}
	contents := make([]byte, 1024)
	n, err := file.Read(contents)
	if err != nil {
		responce := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("500 Internal Server Error"), "500 Internal Server Error")
		conn.Write([]byte(responce))
		return
	}
	responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", n, string(contents[:n]))
	conn.Write([]byte(responce))
}

func saveFile(conn net.Conn, path string, dir string, body string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		responce := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("500 Internal Server Error"), "500 Internal Server Error")
		conn.Write([]byte(responce))
		return
	}
	n, err := file.Write([]byte(body))
	if err != nil {
		responce := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("500 Internal Server Error"), "500 Internal Server Error")
		conn.Write([]byte(responce))
		return
	}
	responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", n, "File saved")
	conn.Write([]byte(responce))
}
