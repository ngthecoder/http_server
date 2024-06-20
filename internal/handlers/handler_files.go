package internal

import (
	"fmt"
	"net"
	"os"
	"strings"

	responce "github.com/ngthecoder/http_server/internal/responce"
)

func serveFile(conn net.Conn, path string, dir string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		responce.Respond(conn, 404, "Not Found")
		return
	}
	contents := make([]byte, 1024)
	n, err := file.Read(contents)
	if err != nil {
		responce.Respond(conn, 500, "Internal Server Error")
		return
	}
	responce.Respond(conn, 200, string(contents[:n]))
}

func saveFile(conn net.Conn, path string, dir string, body string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		responce.Respond(conn, 500, "Internal Server Error")
		return
	}
	n, err := file.Write([]byte(body))
	if err != nil {
		responce.Respond(conn, 500, "Internal Server Error")
		return
	}
	responce.Respond(conn, 201, fmt.Sprintf("File %s created with %d bytes", fileName, n))
}
