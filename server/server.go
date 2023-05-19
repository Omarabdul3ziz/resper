package main

import (
	"fmt"
	"net"
	"os"
)

const (
	port = 8888
)

func readConnection(conn net.Conn) {
	defer conn.Close()
	for {
		msg := make([]byte, 1024)
		_, err := conn.Read(msg)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return
		}
		handleRequest(string(msg))
	}
}

func handleRequest(msg string) {
	fmt.Print("< ", msg)
}

func main() {
	fmt.Println("Opened server on port: ", port)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go readConnection(conn)
	}

}
