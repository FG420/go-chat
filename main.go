package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	defer func() {
		log.Println("Closing connection")
		defer conn.Close()
	}()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error %s when reading message from client", err)
		}

		if mt == websocket.BinaryMessage {
			err = conn.WriteMessage(websocket.TextMessage, []byte("server don't support binary messages"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
		}
		log.Printf("Received message %s", string(message))

		if strings.TrimSpace(string(message)) != "start" {
			err = conn.WriteMessage(websocket.TextMessage, []byte("you didn't type the right word!"))
			if err != nil {
				log.Printf("Error %s when sending message to the client", err)
				return
			}
			continue
		}
		log.Println("start responding to client...")
		i := 1
		for {
			res := fmt.Sprintf("notification %d", i)
			err := conn.WriteMessage(websocket.TextMessage, []byte(res))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
				return
			}
			i = i + 1
			time.Sleep(2 * time.Second)
		}

	}
}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	mux := http.NewServeMux()

	mux.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
