package memcache

import (
	"encoding/json"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
)

var ErrCacheMiss = memcache.ErrCacheMiss

type singleton struct {
	client *memcache.Client
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			client: memcache.New("localhost:11211"),
		}
	})
	return instance
}

func (s *singleton) Set(key string, value interface{}, ttlSecond int32) error {
	byt, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.client.Set(&memcache.Item{
		Key:        key,
		Value:      byt,
		Expiration: ttlSecond,
	})
}

func (s *singleton) Get(key string) (interface{}, error) {
	item, err := s.client.Get(key)
	if err != nil {
		return nil, err
	}

	var value interface{}
	if err := json.Unmarshal(item.Value, &value); err != nil {
		return false, err
	}

	return value, nil
}
