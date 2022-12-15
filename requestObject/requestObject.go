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

func handleBytes(num int, c net.Conn) []byte {

	buf := make([]byte, num)

	_, err := c.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	return buf
}

func NewHandlerObj(c net.Conn) (GlobalTCPObj, error) {
	fmt.Println("RAAAAAAAAAAAAAANNNNNNNNNNNNNN")
	command := handleBytes(3, c)

	/* 	fmt.Printf("1. command = %s\n", command) */
	fmt.Print("command === ", string(command)+"end")
	if string(command) == "bye" {
		populatedGlobalTCPObj := GlobalTCPObj{
			Command: string(command),
		}

		return populatedGlobalTCPObj, nil
	}
	keyBytes := handleBytes(1, c)
	keyBytesAsInt, err := strconv.Atoi(string(keyBytes[0]))

	if err != nil {
		fmt.Println("1error : ", err)
	}
	/* 	fmt.Printf("2. key as int = %v\n", keyBytesAsInt) */

	keyByteSizeStringOfDigits := handleBytes(keyBytesAsInt, c)

	byteSizeStringOfDigitsAsInt, err := strconv.Atoi(string(keyByteSizeStringOfDigits))

	if err != nil {
		fmt.Println("some kind of error")
		return GlobalTCPObj{}, errors.New("Not numeric values")
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

		fmt.Printf("%v ", string(valueBytes[0]))
		valueBytesAsInt, err = strconv.Atoi(string(valueBytes[0]))

		if err != nil {
			fmt.Println("2some kind of error", err)
		}

		valueByteSizeStringOfDigits := handleBytes(valueBytesAsInt, c)
		print("22222what is this ==== " + fmt.Sprint(string(valueByteSizeStringOfDigits)) + "end")
		valueByteSizeStringOfDigitsAsInt, err = strconv.Atoi(string(valueByteSizeStringOfDigits))

		if err != nil {
			fmt.Println("3some kind of error", err)
		}

		AcutalValue := handleBytes(valueByteSizeStringOfDigitsAsInt, c)
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

	fmt.Print("OBJ ==== ", populatedGlobalTCPObj)
	return populatedGlobalTCPObj, nil
}
