package server

import (
	"math/rand"
	"modb/modb"
	"testing"
	"time"
)

var myDB = modb.NewMoDB()
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

func TestServe(t *testing.T) {

	/*for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}*/
}