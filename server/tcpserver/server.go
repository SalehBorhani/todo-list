package main

import (
	"fmt"
	"github.com/salehborhani/todo-list/constance"
	"net"
)

func main() {
	listener, err := net.Listen(constance.Network, constance.Address)

	if err != nil {
		fmt.Println("Couldn't create connection: ", err)

		return
	}

	fmt.Println("The server is running on: ", listener.Addr())

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Couldn't close connection: ", err)
		}
	}(listener)

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Couldn't accept connection: ", err)

			continue
		}
		fmt.Printf("Connection Accepted. Client's IP: %s\n", connection.RemoteAddr())

		req := make([]byte, 1024)
		_, rErr := connection.Read(req)

		if rErr != nil {
			fmt.Println("Couldn't read the data: ", rErr)

			continue
		}

		fmt.Println("The client request: ", string(req))

	}

}
