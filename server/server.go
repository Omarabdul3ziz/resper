package main

import (
	"fmt"
	"net"
	"os"

	"github.com/omarabdul3ziz/resper/pkg/codec"
	"github.com/omarabdul3ziz/resper/pkg/ioutils"
)

const (
	port = 8888
)

func handleServerConnection(conn net.Conn) {
	defer conn.Close()

	for {
		msg := ioutils.Read(conn)
		fmt.Println("Request: ", msg)

		// Decode
		msg_ := codec.Decode(msg)
		fmt.Println("Decoded: ", msg_)

		// res := handlers.Handle(msg)
		// fmt.Println("Response: ", res)

		// ioutils.Write(conn, res)
		// fmt.Print(msg)
	}
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

		go handleServerConnection(conn)
	}

}
