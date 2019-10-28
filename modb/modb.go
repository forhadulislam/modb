package modb

import "sync"

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

func (storage *MoDB) Get() (err error){

	return
}

func (storage *MoDB) Delete() (err error){

	return
}

func (storage *MoDB) Expire() (err error){

	return
}

func (storage *MoDB) Close() (err error){

	return
}

func (storage *MoDB) Size() (size int, err error){

	return
}