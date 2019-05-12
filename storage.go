package goui

import (
	"encoding/json"
	"github.com/fipress/fml"
	"sync"
)

const filename = "store"

type storage struct {
	store *fml.FML
}

var store *storage
var once sync.Once

func Storage() *storage {
	once.Do(func() {
		f, err := fml.Load(filename)
		if err != nil {
			Log("storage - open file failed:", err)
			f = fml.NewFml()
		}
		store = &storage{f}
	})
	return store
}

func (s *storage) Put(key string, v interface{}) {
	s.store.SetValue(key, v)
	s.store.WriteToFile(filename)
}

func (s *storage) GetInt(key string) int {
	return s.store.GetInt(key)
}

func (s *storage) GetString(key string) string {
	return s.store.GetString(key)
}

func (s *storage) PutStruct(key string, v interface{}) (err error) {
	b, err := json.Marshal(v)
	if err != nil {
		Log("Put struct - marshal error:", err)
		return
	}
	s.store.SetValue(key, string(b))
	s.store.WriteToFile(filename)
	return
}

func (s *storage) GetStruct(key string, v interface{}) (err error) {
	str := s.store.GetString(key)
	if str != "" {
		err = json.Unmarshal([]byte(str), v)
		if err != nil {
			Log("Get struct - unmarshal error:", err)
			return
		}
	}
	return
}
