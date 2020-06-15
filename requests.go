package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server()
}

func simpleRequest() {
	response, _ := http.Get("https://google.com")

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Done!\n", body[0])
}

func handler() {
	http.HandleFunc("/welcome", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("Welcome!", request.URL.RawPath)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func server() {
	server := http.Server{
		Addr:           "127.0.0.1:7000",
		Handler:        serverHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening at " + server.Addr)
	err := server.ListenAndServe()
	log.Println("Error Log >>", err)
}

func serverHandler() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Println("Request Log >>", request.Host, request.Method)
	})
}
