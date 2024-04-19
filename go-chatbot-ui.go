package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NDoolan360/go-chatbot-ui/client"
	"github.com/NDoolan360/go-chatbot-ui/server"
)

func main() {
	http.HandleFunc("/web-socket", http.HandlerFunc(server.HandleWebSocket))
	http.HandleFunc("/", http.HandlerFunc(handleIndex))

	port := ":3000"
	fmt.Printf("Server listening on port:\thttp://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("http.ListenAndServe(%s, nil) failed with: %v", port, err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		client.GetIndex(w, r)
	} else {
		publicDir := http.Dir("client/public")
		fileHandler := http.StripPrefix("/", http.FileServer(publicDir))
		fileHandler.ServeHTTP(w, r)
	}
}
