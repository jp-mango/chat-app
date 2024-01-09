package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Chat server started...")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	clients := make(map[net.Conn]bool)
	messages := make(chan string)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}

			clients[conn] = true
			go handleConnection(conn, messages)
		}
	}()

	for msg := range messages {
		for client := range clients {
			_, err := client.Write([]byte(msg))
			if err != nil {
				fmt.Println(err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func handleConnection(conn net.Conn, messages chan<- string) {
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		messages <- message
	}
}
