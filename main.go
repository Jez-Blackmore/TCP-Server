package main

import (
	"errors"
	"fmt"
	"net"
	"week3Project-TCP/requestObject"
	"week3Project-TCP/requests"
	"week3Project-TCP/store"
)

func HandleFetch(r requestObject.GlobalTCPObj) (string, error) {
	switch string(r.Command) {
	case "put":
		valuseString := requests.Put(r)

		if valuseString == "" {
			return "", errors.New("Not found")
		}

		fmt.Println("Value added: ", valuseString)
		return "ack", nil

	case "del":
		valuseString := requests.Delete(r.Key)

		if valuseString == "" {
			fmt.Println("Value is not valid: ")
			return "ack", nil
		}

		fmt.Println("Value deleted: ", valuseString)
		return "ack", nil

	case "get":
		valuseString := requests.Get(r.Key, r)

		if valuseString == "" {
			return "", errors.New("Not found")
		}

		fmt.Println("Get value: ", valuseString)
		return valuseString, nil

	case "bye":
		/* requests.Bye() */
		fmt.Println("Say good bye")
		return "bye", nil
	default:
		fmt.Println("something default")
		return "ack", nil
	}
}

func Handler(c net.Conn) {

	defer c.Close()

	for {
		fetchObj, err := requestObject.NewHandlerObj(c)

		if fetchObj.Command == "EOF" {
			break
		} else if err != nil {
			c.Write([]byte("err"))
		} else {
			confirm, err := HandleFetch(fetchObj)

			if err != nil {
				c.Write([]byte("nil")) // nil
			} else if confirm == "bye" {
				c.Close()
				break
			} else {
				// Send a response back to person contacting us.
				c.Write([]byte(confirm)) // ack
			}
		}
	}

	// Close the connection when you're done with it.

}

func main() {

	store.MainStoreMain = store.NewStoreMain()

	go store.MainStoreMain.Monitor()

	listener, err := net.Listen("tcp4", ":1234")
	if err != nil {
		panic(err)
	}
	defer func() { _ = listener.Close() }()
	for {
		c, err := listener.Accept()
		if err != nil {
			break
		}

		go Handler(c)
	}
}
