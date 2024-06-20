package responces

import (
	"fmt"
	"net"
)

func RespondOK(conn net.Conn, message string) {
	responce := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(message), message)
	conn.Write([]byte(responce))
}

func RespondCreated(conn net.Conn, message string) {
	responce := fmt.Sprintf("HTTP/1.1 201 Created\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len(message), message)
	conn.Write([]byte(responce))
}

func RespondNotFound(conn net.Conn) {
	responce := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("404 Not Found"), "404 Not Found")
	conn.Write([]byte(responce))
}

func RespondMethodNotAllowed(conn net.Conn) {
	responce := fmt.Sprintf("HTTP/1.1 405 Method Not Allowed\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("405 Method Not Allowed"), "405 Method Not Allowed")
	conn.Write([]byte(responce))
}

func RespondServerError(conn net.Conn) {
	responce := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", len("500 Internal Server Error"), "500 Internal Server Error")
	conn.Write([]byte(responce))
}
