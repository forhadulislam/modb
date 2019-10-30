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
			fmt.Fprintf(writer, "%s : %s", moData.Method, moData.Key)

			// SET
			if(moData.Method == "SET"){
				error := db.Set( []byte(moData.Key), []byte(moData.Value) )
				if error != nil {
					log.Println("Error ", error)
				}
			}

			// GET
			if(moData.Method == "GET"){
				getdata, error := db.Get( []byte(moData.Key) )
				if error != nil {
					log.Println("Error ", error)
				}
				log.Println( "Key: ", moData.Key )
				log.Println( "Value: ", string(getdata) )
			}

			// DELETE
			if(moData.Method == "DELETE"){
				error := db.Delete( []byte(moData.Key) )
				if error != nil {
					log.Println(error)
				}
			}

			log.Println(db)

		} else{

			fmt.Fprintf(writer, "%s", "<h2>Method not Allowed</h2>" )

		}
	})
	fmt.Println("MoDB server is running at : http://127.0.0.1:" + MoDBPort + "/db")
	http.ListenAndServe(":" + MoDBPort, nil)
}