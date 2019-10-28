package server

import (
	"fmt"
	"log"
	"modb/modb"
	"net/http"
)

const MoDBPort = "34443"

func Serve(db *modb.MoDB) {

	http.HandleFunc("/db", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			fmt.Fprintf(writer, "%s", "<h2>MoDB POST</h2>" )
			log.Println(db)
		}else {
			fmt.Fprintf(writer, "%s", "<h2>MoDB Home</h2>" )
			log.Println(db)
		}
	})
	fmt.Println("MoDB server is running at port : http://127.0.0.1:" + MoDBPort + "/db")
	http.ListenAndServe(":" + MoDBPort, nil)

}