package requestObject

import (
	"fmt"
	"net"
	"strconv"
)

type GlobalTCPObj struct {
	Command       string
	keyBytes      int
	keyByteSize   int
	Key           string
	valueBytes    int
	valueByteSize int
	Value         string
}

var (
	populatedGlobalTCPObj GlobalTCPObj
)

func handleBytes(num int, c net.Conn) []byte {

	buf := make([]byte, num)

	_, err := c.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	return buf
}

func NewHandlerObj(c net.Conn) GlobalTCPObj {

	command := handleBytes(3, c)

	/* 	fmt.Printf("1. command = %s\n", command) */

	if string(command) == "bye" {
		populatedGlobalTCPObj := GlobalTCPObj{
			Command: string(command),
		}

		return populatedGlobalTCPObj
	}
	keyBytes := handleBytes(1, c)
	keyBytesAsInt, err := strconv.Atoi(string(keyBytes[0]))

	if err != nil {
		fmt.Println("some kind of error")
	}
	/* 	fmt.Printf("2. key as int = %v\n", keyBytesAsInt) */

	keyByteSizeStringOfDigits := handleBytes(keyBytesAsInt, c)
	byteSizeStringOfDigitsAsInt, err := strconv.Atoi(string(keyByteSizeStringOfDigits[0]))

	if err != nil {
		fmt.Println("some kind of error")
	}

	/* fmt.Printf("3. key as int = %v\n", byteSizeStringOfDigitsAsInt) */
	AcutalKeyValue := handleBytes(byteSizeStringOfDigitsAsInt, c)

	AcutalKeyValueAsString := string(AcutalKeyValue)
	/* fmt.Print("test: ", AcutalKeyValueAsString) */

	var valueByteSizeStringOfDigitsAsInt int
	var valueBytesAsInt int
	var AcutalValueAsString string

	if string(command) == "put" {

		valueBytes := handleBytes(1, c)
		valueBytesAsInt, err := strconv.Atoi(string(valueBytes[0]))

		if err != nil {
			fmt.Println("some kind of error")
		}

		valueByteSizeStringOfDigits := handleBytes(valueBytesAsInt, c)
		valueByteSizeStringOfDigitsAsInt, err = strconv.Atoi(string(valueByteSizeStringOfDigits[0]))

		if err != nil {
			fmt.Println("some kind of error")
		}

		AcutalValue := handleBytes(valueByteSizeStringOfDigitsAsInt, c)
		AcutalValueAsString = string(AcutalValue)
	}

	populatedGlobalTCPObj := GlobalTCPObj{
		Command:       string(command),
		keyBytes:      keyBytesAsInt,
		keyByteSize:   byteSizeStringOfDigitsAsInt,
		Key:           AcutalKeyValueAsString,
		valueBytes:    valueBytesAsInt,
		valueByteSize: valueByteSizeStringOfDigitsAsInt,
		Value:         AcutalValueAsString,
	}

	return populatedGlobalTCPObj
}
