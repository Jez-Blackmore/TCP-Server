package main

import (
	"fmt"
	"net"
	"week3Project-TCP/requests"
	"week3Project-TCP/store"
)

func handle(c net.Conn) {

	buf := make([]byte, 1000)

	_, err := c.Read(buf)

	myString := string(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	command := myString[0:3]
	/* 	fmt.Printf("Test2:   %v\n", myString) */

	/* 	runes := []rune(string(myString)) */

	/* 	fmt.Print("test1: ", runes) */
	/* myStringFirst3 := string(runes[1:3]) */

	/* 	fmt.Print("test1: ", myStringFirst3) */

	switch command {
	case "put":

		/* keyLength := string(myString)[3] */
		/* 	fmt.Printf("oh dear: " + string(keyLength) + "grr") */
		/* last22 := 3 + int(keyLength) */
		/* fmt.Printf("GGGGGGGG: %v", keyLength) */
		/* fmt.Printf("oh dear: " + string(last) + "grr") */
		/* valuseToUse := string(myString)[4:5] */

		/* 	fmt.Printf("my string 111111 = %v", myString[5:6]) */
		fmt.Printf("put in new value. with: %v\n", "cdsd")

		requests.Put("key", "stored value")

	case "del":

		requests.Delete("key")
		fmt.Println("delete something")
	case "get":
		requests.Get("put")
		fmt.Println("This is a get")
	case "bye":
		requests.Bye()
		fmt.Println("Say good bye")
	default:
		fmt.Println("something default")
	}

	// Send a response back to person contacting us.
	c.Write([]byte(myString))
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
		go handle(c)
	}
}
