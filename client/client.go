package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/omarabdul3ziz/resper/pkg/ioutils"
)

const (
	port = 8888
)

func handleClientConnection(conn net.Conn) {
	defer conn.Close()

	// create buffer from the std input
	scanner := bufio.NewScanner(os.Stdin)

	// keep scanning until exit
	for {
		// as prompt
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		// write to the connection
		ioutils.Write(conn, input)

		// read from the connection
		res := ioutils.Read(conn)

		// print the response
		fmt.Println("<", res)
	}

}

func main() {
	// connect to the open socket
	fmt.Println("Connected to the server on port: ", port)

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		fmt.Println("Error Dialling: ", err.Error())
	}

	handleClientConnection(conn)
}
