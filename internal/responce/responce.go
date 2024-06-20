package responce

import (
	"fmt"
	"net"
)

func Respond(conn net.Conn, statusCode int, message string) {
	responce := fmt.Sprintf("HTTP/1.1 %d\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s", statusCode, len(message), message)
	conn.Write([]byte(responce))
}
