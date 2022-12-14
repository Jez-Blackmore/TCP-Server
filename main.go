package main

import (
	"errors"
	"fmt"
	"net"
	"week3Project-TCP/requestObject"
	"week3Project-TCP/requests"
	"week3Project-TCP/store"
)

func handleFetch(r requestObject.GlobalTCPObj) (string, error) {
	switch string(r.Command) {
	case "put":
		valuseString := requests.Put(r.Key, r.Value)

		if valuseString == "" {
			return "", errors.New("Not found")
		}

		fmt.Println("Value added: ", valuseString)
		return "ack", nil

	case "del":
		valuseString := requests.Delete(r.Key)

		if valuseString == "" {
			return "", errors.New("Not found")
		}

		fmt.Println("Value deleted: ", valuseString)
		return "ack", nil

	case "get":
		valuseString := requests.Get(r.Key)

		if valuseString == "" {
			return "", errors.New("Not found")
		}

		fmt.Println("Get value: ", valuseString)
		return "ack", nil

	case "bye":
		requests.Bye()
		fmt.Println("Say good bye")
		return "ack", nil
	default:
		fmt.Println("something default")
		return "ack", nil
	}
}

func handler(c net.Conn) {
	fetchObj := requestObject.NewHandlerObj(c)

	confirm, err := handleFetch(fetchObj)

	if err != nil {
		c.Write([]byte("nil"))
	} else {
		// Send a response back to person contacting us.
		c.Write([]byte(confirm))
	}

	// Close the connection when you're done with it.
	c.Close()
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

		go handler(c)
	}
}
