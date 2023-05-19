package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	port = 8888
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		fmt.Println("Error Dialling: ", err.Error())
	}

	defer conn.Close()

	// create buffer from the std input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// as prompt
		fmt.Print("> ")

		// scan
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		// write to the connection
		fmt.Fprintln(conn, input)
	}

}
