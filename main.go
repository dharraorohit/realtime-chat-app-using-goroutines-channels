package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dharraorohit/scalable-chat-app/srvhandlers"
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	messageHandler := &srvhandlers.MessageHandler{}

	r := mux.NewRouter()
	r.Handle("/", messageHandler)
	return r
}

func main() {
	router := getRouter()
	server := http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Print("Starting Server...")
	log.Fatal(server.ListenAndServe())
}
