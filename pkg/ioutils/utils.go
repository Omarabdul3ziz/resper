package ioutils

import (
	"fmt"
	"net"
	"os"
	"strings"
	"unicode"
)

func trimSpacesOnly(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) && r != ' '
	})
}

func Write(conn net.Conn, msg string) {
	res := trimSpacesOnly(msg)

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

	return trimSpacesOnly(string(buffer[:len]))
}
