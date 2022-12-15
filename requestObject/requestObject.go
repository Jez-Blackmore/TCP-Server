package requestObject

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

type GlobalTCPObj struct {
	Command       string
	KeyBytes      int
	KeyByteSize   int
	Key           string
	ValueBytes    int
	ValueByteSize int
	Value         string
}

var (
	populatedGlobalTCPObj GlobalTCPObj
)

func handleBytes(num int, c net.Conn) ([]byte, string) {

	buf := make([]byte, num)

	_, err := c.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return make([]byte, 1), "error"
	}

	return buf, ""
}

func NewHandlerObj(c net.Conn) (GlobalTCPObj, error) {

	command, stringInfo := handleBytes(3, c)

	if stringInfo != "" {

		populatedGlobalTCPObj := GlobalTCPObj{
			Command: "EOF",
		}
		return populatedGlobalTCPObj, nil
	}

	if string(command) == "bye" {
		populatedGlobalTCPObj := GlobalTCPObj{
			Command: string(command),
		}

		return populatedGlobalTCPObj, nil
	}
	keyBytes, _ := handleBytes(1, c)
	keyBytesAsInt, err := strconv.Atoi(string(keyBytes[0]))

	if err != nil {
		fmt.Println("1. Error : ", err)
	}

	keyByteSizeStringOfDigits, _ := handleBytes(keyBytesAsInt, c)

	byteSizeStringOfDigitsAsInt, err := strconv.Atoi(string(keyByteSizeStringOfDigits))

	if err != nil {
		fmt.Println("2. Error: ")
		return GlobalTCPObj{}, errors.New("Not numeric values")
	}

	AcutalKeyValue, _ := handleBytes(byteSizeStringOfDigitsAsInt, c)

	AcutalKeyValueAsString := string(AcutalKeyValue)

	var valueByteSizeStringOfDigitsAsInt int
	var valueBytesAsInt int
	var AcutalValueAsString string

	if string(command) == "put" {

		valueBytes, _ := handleBytes(1, c)

		/* fmt.Printf("%v ", string(valueBytes[0])) */
		valueBytesAsInt, err = strconv.Atoi(string(valueBytes[0]))

		if err != nil {
			fmt.Println("3. Error: ", err)
		}

		valueByteSizeStringOfDigits, _ := handleBytes(valueBytesAsInt, c)

		valueByteSizeStringOfDigitsAsInt, err = strconv.Atoi(string(valueByteSizeStringOfDigits))

		if err != nil {
			fmt.Println("4. Error: ", err)
		}

		AcutalValue, _ := handleBytes(valueByteSizeStringOfDigitsAsInt, c)
		AcutalValueAsString = string(AcutalValue)
	}

	populatedGlobalTCPObj := GlobalTCPObj{
		Command:       string(command),
		KeyBytes:      keyBytesAsInt,
		KeyByteSize:   byteSizeStringOfDigitsAsInt,
		Key:           AcutalKeyValueAsString,
		ValueBytes:    valueBytesAsInt,
		ValueByteSize: valueByteSizeStringOfDigitsAsInt,
		Value:         AcutalValueAsString,
	}

	return populatedGlobalTCPObj, nil
}
