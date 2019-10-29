package server

import (
	"encoding/json"
	"fmt"
	"log"
	"modb/modb"
	"net/http"
)
const MoDBPort = "34443"

type MoDBServer struct {
	Method string
	Key string
	Value string
}

func Serve(db *modb.MoDB) {
	http.HandleFunc("/db", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {

			decoder := json.NewDecoder(request.Body)
			var moData MoDBServer
			err := decoder.Decode(&moData)
			if err != nil {
				log.Println(err)
			}
			fmt.Fprintf(writer, "%s", "<h2>MoDB POST</h2>" )
			fmt.Fprintf(writer, "%s", moData.Key)
			if(moData.Method == "SET"){
				db.Set( []byte(moData.Key), []byte(moData.Value) )
			}
			log.Println(db)

		} else if request.Method == "PUT"  {

			fmt.Fprintf(writer, "%s", "<h2>MoDB GET</h2>" )
			decoder := json.NewDecoder(request.Body)
			var moData MoDBServer
			error := decoder.Decode(&moData)
			if error != nil {
				log.Println("error ", error)
			}

			if(moData.Method == "GET"){
				getdata, error := db.Get( []byte(moData.Key) )
				if error != nil {
					log.Println("error ", error)
				}
				log.Println( moData.Key )
				log.Println( string(getdata) )
			}
			log.Println(db)

		}else if request.Method == "DELETE"  {

			fmt.Fprintf(writer, "%s", "<h2>MoDB DELETE</h2>" )
			decoder := json.NewDecoder(request.Body)
			var moData MoDBServer
			err := decoder.Decode(&moData)
			if err != nil {
				log.Println(err)
			}

			if(moData.Method == "DELETE"){
				error := db.Delete( []byte(moData.Key) )
				if error != nil {
					log.Println(err)
				}
			}
			log.Println(db)

		}
	})
	fmt.Println("MoDB server is running at : http://127.0.0.1:" + MoDBPort + "/db")
	http.ListenAndServe(":" + MoDBPort, nil)
}