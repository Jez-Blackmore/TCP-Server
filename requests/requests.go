package requests

import (
	"fmt"
	"week3Project-TCP/store"
)

func Put(key string, value string) string {

	objString := store.MainStoreMain.PutRequest(key, value)

	/* fmt.Print("Put") */

	return objString.Key
}

func Get(key string) string {

	objString := store.MainStoreMain.GetRequest(key)

	return objString.Key

}

func Delete(key string) string {

	objString := store.MainStoreMain.DeleteRequest(key)

	fmt.Print("Delete")
	return objString.Key
}

func Bye() string {

	fmt.Print("Bye")

	return "bye"
}
