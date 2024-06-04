package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

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
