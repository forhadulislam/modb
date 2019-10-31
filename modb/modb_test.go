package modb

import (
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

var myDB = NewMoDB()
var totalData = 100000
var keyList = make([]string, 0)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func SeedString(length int) string {
	return StringWithCharset(length, charset)
}

func TestMoDB_Set(t *testing.T) {
	for i:=0; i<totalData; i++ {
		t.Run("Insert data", func(t *testing.T) {
			key := []byte(SeedString(12))
			value := []byte(SeedString(50))
			keyList = append(keyList, string(key))
			err := myDB.Set(key, value)
			if err != nil {
				t.Errorf("Set() error = %v", err)
			}
		})
	}

	dbSize, _ := myDB.Size()
	assert.Equal(t, totalData, dbSize, "Data length should be equal to totalData length")
	assert.Equal(t, totalData, len(keyList), "keyList length should be equal to totalData length")
}

func TestMoDB_Delete(t *testing.T) {
	dbSize, _ := myDB.Size()
	assert.Equal(t, totalData, dbSize, "Data length should be equal to totalData length")
	assert.Equal(t, totalData, len(keyList), "keyList length should be equal to totalData length")

	for i:=0; i<totalData; i++ {
		t.Run("Deleting from DB", func(t *testing.T) {
			myDB.Delete([]byte(keyList[i]))
		})
	}
	dbSize, _ = myDB.Size()
	assert.Equal(t, 0, dbSize, "keyList length should be equal to totalData length")
}

func TestMoDB_Close(t *testing.T) {
	type fields struct {
		storage map[string][]byte
		lock    sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &MoDB{
				storage: tt.fields.storage,
				lock:    tt.fields.lock,
			}
			if err := storage.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



func TestMoDB_Expire(t *testing.T) {
	type fields struct {
		storage map[string][]byte
		lock    sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &MoDB{
				storage: tt.fields.storage,
				lock:    tt.fields.lock,
			}
			if err := storage.Expire(); (err != nil) != tt.wantErr {
				t.Errorf("Expire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoDB_Get(t *testing.T) {
	type fields struct {
		storage map[string][]byte
		lock    sync.RWMutex
	}
	type args struct {
		key []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &MoDB{
				storage: tt.fields.storage,
				lock:    tt.fields.lock,
			}
			gotValue, err := storage.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestMoDB_Has(t *testing.T) {
	type fields struct {
		storage map[string][]byte
		lock    sync.RWMutex
	}
	type args struct {
		key []byte
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHasKey bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &MoDB{
				storage: tt.fields.storage,
				lock:    tt.fields.lock,
			}
			if gotHasKey := storage.Has(tt.args.key); gotHasKey != tt.wantHasKey {
				t.Errorf("Has() = %v, want %v", gotHasKey, tt.wantHasKey)
			}
		})
	}
}

func TestMoDB_Size(t *testing.T) {
	type fields struct {
		storage map[string][]byte
		lock    sync.RWMutex
	}
	tests := []struct {
		name     string
		fields   fields
		wantSize int
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &MoDB{
				storage: tt.fields.storage,
				lock:    tt.fields.lock,
			}
			gotSize, err := storage.Size()
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSize != tt.wantSize {
				t.Errorf("Size() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestNewMoDB(t *testing.T) {
	tests := []struct {
		name string
		want *MoDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMoDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMoDB() = %v, want %v", got, tt.want)
			}
		})
	}
}