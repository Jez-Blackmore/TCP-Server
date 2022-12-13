package store

import "fmt"

type Key string

type StoreMain struct {
	key           map[string]StructValueObject
	putChannel    chan PutChannel
	deleteChannel chan DeleteChannel
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
		key: map[string]StructValueObject{},
	}

	return Store
}

func (s *StoreMain) Monitor() {
	for {
		select {
		case putVal := <-s.putChannel:

			putVal.ResponseChannelPut <- ResponseChannel{Value: "1", key: "1"}

		case deleteVal := <-s.deleteChannel:

			deleteVal.ResponseChannelDelete <- ResponseChannel{Value: "1", key: "1"}
		}
	}
}

func (s *StoreMain) GetRequest(key string) StructValueObject {

	fmt.Printf("%v", s.key[key])
	return s.key[key]

}

func (s *StoreMain) PutRequest(key string, value string) StructValueObject {

	s.putChannel <- PutChannel{key: "1", Value: "1"}

	confirm := <-s.putChannel

	fmt.Printf("%v ", confirm)

	return s.key[key]

}

func (s *StoreMain) DeleteRequest(key string) StructValueObject {

	s.deleteChannel <- DeleteChannel{key: "1", Value: "1"}

	confirm := <-s.deleteChannel

	fmt.Printf("%v ", confirm)

	return s.key[key]

}
