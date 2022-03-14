package main

import (
	"log"
	"net/http"
	"os"
)

func mirrorHandler(writer http.ResponseWriter, request *http.Request) {
	request.Write(writer)
}

func main() {
	// Port is last argument
	port := "8080"
	for _, arg := range os.Args {
		port = arg
	}
	log.Println("Listening on :" + port)
	http.HandleFunc("/", mirrorHandler)
	http.ListenAndServe(":"+port, nil)
}
