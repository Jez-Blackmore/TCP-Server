package requests

import (
	"fmt"
	"week3Project-TCP/requestObject"
	"week3Project-TCP/store"
)

func Put(r requestObject.GlobalTCPObj) string {

	objString := store.MainStoreMain.PutRequest(r)

	return objString.Value

}

func Get(key string, r requestObject.GlobalTCPObj) string {

	valueString := store.MainStoreMain.GetRequest(key, r)

	return valueString

}

func Delete(key string) string {

	objString, err := store.MainStoreMain.DeleteRequest(key)

	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return ""
	}

	return objString.Value

}
