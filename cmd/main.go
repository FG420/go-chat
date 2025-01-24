package main

import (
	"log"
	"net/http"

	"github.com/GF420/go-chat/cmd/server"
)

func enableHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
}

func main() {
	log.Println("Server starting at port 8080...")

	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		enableHeaders(&w)
		server.LoginHandler(w, req)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, req *http.Request) {
		enableHeaders(&w)
		server.LogoutHandler(w, req)
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
