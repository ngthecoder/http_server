package internal

import (
	"fmt"
	"net"
	"os"
	"strings"

	responces "github.com/ngthecoder/http_server/internal/responces"
)

func serveFile(conn net.Conn, path string, dir string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		responces.RespondNotFound(conn)
		return
	}
	contents := make([]byte, 1024)
	n, err := file.Read(contents)
	if err != nil {
		responces.RespondServerError(conn)
		return
	}
	responces.RespondOK(conn, string(contents[:n]))
}

func saveFile(conn net.Conn, path string, dir string, body string) {
	fileName := strings.TrimPrefix(path, "/files")
	filePath := fmt.Sprint(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		responces.RespondServerError(conn)
		return
	}
	n, err := file.Write([]byte(body))
	if err != nil {
		responces.RespondServerError(conn)
		return
	}
	responces.RespondCreated(conn, fmt.Sprintf("File %s created with %d bytes", fileName, n))
}
