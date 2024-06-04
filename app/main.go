package main

import (
	"flag"
	"fmt"
	"net"
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
