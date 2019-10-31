package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var totalData = 10000
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
	testUrl := "http://127.0.0.1" + MoDBPort + "/db"
	for i:=0; i<totalData; i++ {
		t.Run("Insert data", func(t *testing.T) {
			requsetBody, err := json.Marshal(map[string]string{
				"Method": "SET",
				"Key": SeedString(12),
				"Value": SeedString(40),
			})
			if err != nil {
				fmt.Errorf("Cannot marshal code")
			}

			_, err = http.Post(
				testUrl,
				"text/plain",
				bytes.NewBuffer(requsetBody),
				)

			if err != nil {
				fmt.Errorf("Error happened")
			}
		})
	}

}