package ioutils

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func Write(conn net.Conn, msg string) {
	res := strings.TrimSpace(msg)

	response := []byte(res)
	_, err := conn.Write(response)
	if err != nil {
		fmt.Println("Error writing data:", err)
		os.Exit(1)
	}
}

func Read(conn net.Conn) string {
	buffer := make([]byte, 1024)

	len, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return ""
	}

	return strings.TrimSpace(string(buffer[:len]))
}
