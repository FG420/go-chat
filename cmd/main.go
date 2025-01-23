package main

import (
	"log"
	"net/http"

	"github.com/GF420/go-chat/cmd/server"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	log.Println("Server starting at port 8080...")

	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		enableCors(&w)
		server.LoginHandler(w, req)
	})

	room := server.NewRoom()
	go room.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.Inizialiaze(room, w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
