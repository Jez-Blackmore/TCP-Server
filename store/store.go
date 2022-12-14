package store

import (
	"errors"
	"fmt"
)

type Key string

type StoreMain struct {
	key           map[string]StructValueObject
	PutChannel    chan PutChannel
	DeleteChannel chan DeleteChannel
}

type PutChannel struct {
	Value              string `json:"value"`
	key                Key
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
	Key   string `json:"key"`
	Value string `json:"value"`
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

			s.key[string(putVal.key)] = StructValueObject{Value: putVal.Value, Key: string(putVal.key)}

			/* fmt.Print("key: ", putVal.key, " value: ", putVal.Value, " object ", s.key[string(putVal.key)]) */

			putVal.ResponseChannelPut <- ResponseChannel{Value: s.key[string(putVal.key)].Value, key: Key(s.key[string(putVal.key)].Key)}

		case deleteVal := <-s.DeleteChannel:

			delete(s.key, string(deleteVal.key))

			deleteVal.ResponseChannelDelete <- ResponseChannel{Value: "1", key: "1"}
		}
	}
}

func (s *StoreMain) GetRequest(key string) string {

	var valueToShow string

	for keyVal, value := range s.key {
		if string(keyVal) == key {
			valueToShow = value.Value
		}
	}

	fmt.Printf("%v", s.key[key])
	return valueToShow
}

func (s *StoreMain) PutRequest(key string, value string) ResponseChannel {

	responseChan := make(chan ResponseChannel)

	s.PutChannel <- PutChannel{key: Key(key), Value: value, ResponseChannelPut: responseChan}

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
