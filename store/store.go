package store

import (
	"errors"
	"fmt"
	"week3Project-TCP/requestObject"
)

type Key string

type StoreMain struct {
	key           map[string]StructValueObject
	PutChannel    chan PutChannel
	DeleteChannel chan DeleteChannel
}

type PutChannel struct {
	RequestObj         StructValueObject `json:"requestObj "`
	ResponseChannelPut chan ResponseChannel
}

type DeleteChannel struct {
	Value                 string `json:"value"`
	key                   Key
	ResponseChannelDelete chan ResponseChannel
}

type ResponseChannel struct {
	Value string `json:"value"`
	key   Key
}

type StructValueObject struct {
	Command string
	Key     string `json:"key"`
	Value   string `json:"value"`

	KeyBytes    int `json:"keyBytes"`
	KeyByteSize int `json:"keyBytesSize"`

	ValueBytes    int `json:"valueBytes"`
	ValueByteSize int `json:"valueByteSize"`
}

var (
	MainStoreMain StoreMain
)

func NewStoreMain() StoreMain {

	Store := StoreMain{
		key:           map[string]StructValueObject{},
		PutChannel:    make(chan PutChannel),
		DeleteChannel: make(chan DeleteChannel),
	}

	return Store
}

func (s *StoreMain) Monitor() {
	for {
		select {
		case putVal := <-s.PutChannel:

			s.key[string(putVal.RequestObj.Key)] = StructValueObject{Value: putVal.RequestObj.Value, Key: string(putVal.RequestObj.Key), KeyBytes: putVal.RequestObj.KeyBytes, KeyByteSize: putVal.RequestObj.KeyByteSize, ValueBytes: putVal.RequestObj.ValueBytes, ValueByteSize: putVal.RequestObj.ValueByteSize}

			putVal.ResponseChannelPut <- ResponseChannel{Value: s.key[string(putVal.RequestObj.Key)].Value, key: Key(s.key[string(putVal.RequestObj.Key)].Key)}

		case deleteVal := <-s.DeleteChannel:

			delete(s.key, string(deleteVal.key))

			deleteVal.ResponseChannelDelete <- ResponseChannel{Value: "1", key: "1"}
		}
	}
}

func (s *StoreMain) GetRequest(key string, r requestObject.GlobalTCPObj) string {

	var valueToShow string

	for keyVal, value := range s.key {
		if string(keyVal) == key {
			valueToShow = "val" + fmt.Sprint(value.ValueBytes) + fmt.Sprint(value.ValueByteSize) + value.Value
		}
	}

	fmt.Printf("%v", s.key[key])
	return valueToShow
}

func (s *StoreMain) PutRequest(r requestObject.GlobalTCPObj) ResponseChannel {

	responseChan := make(chan ResponseChannel)

	s.PutChannel <- PutChannel{RequestObj: StructValueObject{Key: r.Key, Value: r.Value, KeyBytes: r.KeyBytes, KeyByteSize: r.KeyByteSize, ValueBytes: r.ValueBytes, ValueByteSize: r.ValueByteSize}, ResponseChannelPut: responseChan}

	confirmObj := <-responseChan

	return confirmObj

}

func (s *StoreMain) DeleteRequest(key string) (ResponseChannel, error) {

	var valueToShow string

	for keyVal, value := range s.key {
		if string(keyVal) == key {
			valueToShow = value.Value
		}
	}

	if valueToShow == "" {
		return ResponseChannel{key: "", Value: ""}, errors.New("Not found")
	}

	responseChan := make(chan ResponseChannel)

	s.DeleteChannel <- DeleteChannel{key: Key(key), ResponseChannelDelete: responseChan}

	confirm := <-responseChan

	return confirm, nil

}
