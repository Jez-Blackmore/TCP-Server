package main

import (
	"fmt"
	"net"
	"week3Project-TCP/requestObject"
	"week3Project-TCP/requests"
	"week3Project-TCP/store"
)

func handleFetch(r requestObject.GlobalTCPObj) {
	switch string(r.Command) {
	case "put":
		requests.Put(r.Key, r.Value)
		fmt.Println("Put something")
	case "del":
		requests.Delete(r.Key)
		fmt.Println("Delete something")
	case "get":
		requests.Get(r.Key)
		fmt.Println("Get something")
	case "bye":
		requests.Bye()
		fmt.Println("Say good bye")
	default:
		fmt.Println("something default")
	}
}

func handler(c net.Conn) {
	fetchObj := requestObject.NewHandlerObj(c)

	handleFetch(fetchObj)

	// Send a response back to person contacting us.
	c.Write([]byte("ack"))
	// Close the connection when you're done with it.
	c.Close()
}

func main() {

	store.MainStoreMain = store.NewStoreMain()

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
