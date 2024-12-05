package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type (
	WebSocketHandler struct {
		upgrader websocket.Upgrader
	}

	Response struct {
		Text      string
		Timestamp time.Time
	}
)

func (wsh WebSocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var res Response

	c, err := wsh.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("Error during connection upgrade: %s", err)
		return
	}

	defer c.Close()
	res.Text = "U joined a room!"
	res.Timestamp = time.Now()
	c.WriteJSON(res)

	log.Println("remote address: ", c.RemoteAddr().String())

	for {

		mess := readMessages(c, res)
		if strings.Trim(string(mess), "\n") == "" {
			res.Text = "Say something!"
			res.Timestamp = time.Now()
			err := c.WriteJSON(res)
			if err != nil {
				log.Printf("Error: %s", err)
				return
			}
			continue
		}

		resClient(c, res, mess)
	}
}

func readMessages(c *websocket.Conn, res Response) []byte {

	mt, mess, err := c.ReadMessage()
	if err != nil {
		log.Printf("Error during sending messages: %s", err)
		return nil
	}

	if mt == websocket.BinaryMessage {
		// err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))
		res.Text = "Server doesn't support binary messages"
		res.Timestamp = time.Now()
		err = c.WriteJSON(res)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		return nil
	}

	log.Printf("Received message %s", string(mess))
	return mess
}

func resClient(c *websocket.Conn, res Response, mess []byte) {
	log.Println("start responding to client...")
	res.Text = string(mess)
	res.Timestamp = time.Now()
	err := c.WriteJSON(res)
	if err != nil {
		log.Printf("Error when respondin to client: %s", err)
		return
	}
}

func Inizialiaze() {
	wsh := WebSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	http.Handle("/", wsh)
	log.Println("Server starting at port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
