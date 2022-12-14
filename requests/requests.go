package requests

import (
	"fmt"
	"week3Project-TCP/store"
)

func Put(key string, value string) string {

	objString := store.MainStoreMain.PutRequest(key, value)

	/* fmt.Print("Put") */

	return objString.Value
}

func Get(key string) string {

	valueString := store.MainStoreMain.GetRequest(key)

	fmt.Print("Get: ", valueString)
	return valueString

}

func Delete(key string) string {

	objString, err := store.MainStoreMain.DeleteRequest(key)

	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}

	fmt.Print("Delete: ", objString)
	return objString.Value
}

func Bye() string {

	fmt.Print("Bye")

	return "bye"
}
