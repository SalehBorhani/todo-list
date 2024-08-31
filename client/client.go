package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/salehborhani/todo-list/constance"
	"github.com/salehborhani/todo-list/server/requestparam"
	"net"
)

func main() {
	client, err := net.Dial(constance.Network, constance.Address)
	if err != nil {
		fmt.Println("Couldn't dial to the server: ", err)

	}

	defer func(client net.Conn) {
		err := client.Close()
		if err != nil {
			fmt.Println("Couldn't close the connection: ", err)
		}
	}(client)

	cmd := flag.String("command", "", "Does your work")
	data := flag.String("data", "", "Json data you send to the server")

	flag.Parse()

	request := requestparam.Request{Command: *cmd}

	switch request.Command {
	case "create-task":
		serializedData, mErr := json.Marshal(*data)

		if mErr != nil {
			fmt.Println("Couldn't marshal data: ", mErr)
		}

		_, wErr := client.Write(serializedData)

		if wErr != nil {
			fmt.Println("Couldn't write data to the server: ", wErr)
		}
	}

}
