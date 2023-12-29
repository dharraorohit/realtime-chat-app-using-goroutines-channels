package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dharraorohit/realtime-chat-app-using-goroutines-channels/srvhandlers"
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	webSocketHandler := srvhandlers.WebSocketHandler{
		Clients: make(map[int]*srvhandlers.Client),
	}

	r := mux.NewRouter()
	r.HandleFunc("/ws", webSocketHandler.HandleConnection)

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
