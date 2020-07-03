package main

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

func index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Implement
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	
	
	log.Fatal(http.ListenAndServe(":8080", router))
}