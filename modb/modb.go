package modb

import (
	"errors"
	"log"
	"sync"
)

type MoDB struct {
	storage map[string][]byte
	lock sync.RWMutex
}

func NewMoDB() *MoDB{
	return &MoDB{
		storage: make(map[string][]byte),
	}
}

func (storage *MoDB) Set(key []byte, value []byte) (err error){
	storage.lock.Lock()
	defer storage.lock.Unlock()

	storage.storage[string(key)] = value
	return
}

func (storage *MoDB) Get(key []byte) (value []byte, err error){
	storage.lock.Lock()
	defer storage.lock.Unlock()

	value, ok := storage.storage[string(key)]
	if(ok){
		return value, nil
	}
	return nil, errors.New("Invalid key")
}

func (storage *MoDB) Has(key []byte) (hasKey bool){
	storage.lock.Lock()
	defer storage.lock.Unlock()

	_, ok := storage.storage[string(key)]
	if ok{
		return true
	}

	return false
}

func (storage *MoDB) Delete(key []byte) (err error){
	storage.lock.Lock()
	defer storage.lock.Unlock()

	hasKey := storage.Has(key)
	log.Println(string(key))
	if hasKey {
		log.Println(string(key))
		delete(storage.storage, string(key) )
		return nil
	}
	return errors.New("Invalid key")
}

func (storage *MoDB) Expire() (err error){
	return
}

func (storage *MoDB) Close() (err error){
	return
}

func (storage *MoDB) Size() (size int, err error){
	return len(storage.storage), nil
}