package store

type StoreMain struct {
	key map[string]StructValueObject
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

func (s *StoreMain) GetRequest(key string) StructValueObject {

	return s.key[key]

}

func (s *StoreMain) PutRequest(key string, value string) StructValueObject {

	return s.key[key]

}

func (s *StoreMain) DeleteRequest(key string) StructValueObject {

	return s.key[key]

}
