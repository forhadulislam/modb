package modb

import "sync"

type MoDB struct {
	storage map[string][]byte
	lock sync.RWMutex
}

func NewMoDB() *MoDB{

	return &MoDB{}
}


func (storage *MoDB) Set() (err error){

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