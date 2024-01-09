package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go readMessages(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: \n")
		msg, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, msg)
	}
}

func readMessages(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected from server")
			os.Exit(0)
		}
		fmt.Println("New message: " + message)
	}
}
