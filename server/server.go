package server

import (
	"fmt"
	"net/http"
)

const MoDBPort = "34443"

func dbFunc( w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
		fmt.Fprintf(w, "%s", "<h2>MoDB POST</h2>" )
	}
	fmt.Fprintf(w, "%s", "<h2>MoDB Home</h2>" )
}

func Serve() {

	http.HandleFunc("/db", dbFunc)
	fmt.Println("MoDB server is running at port : " + MoDBPort)
	http.ListenAndServe(":" + MoDBPort, nil)

}