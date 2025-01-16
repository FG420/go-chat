package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/GF420/go-chat/cmd/blockchain"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type (
	Chatter struct {
		room       *Room
		conn       *websocket.Conn
		send       chan []byte
		wallet     *blockchain.Wallet
		blockchain *blockchain.Blockchain
	}

	Message struct {
		User      []byte    `json:"user"`
		Text      string    `json:"text"`
		Timestamp time.Time `json:"timestamp"`
	}
)

func (chatter *Chatter) Read() {
	defer func() {
		chatter.room.unregister <- chatter
		chatter.conn.Close()
	}()

	for {
		_, mess, err := chatter.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				chatter.conn.WriteJSON(Message{User: chatter.wallet.PubKey, Text: "This User left the room", Timestamp: time.Now()})
			} else {
				log.Printf("error: %v", err)
				chatter.conn.WriteJSON(Message{User: chatter.wallet.PubKey, Text: "This User left the room", Timestamp: time.Now()})
			}
			break
		}

		mess = bytes.TrimSpace(bytes.Replace(mess, []byte{'\n'}, []byte{' '}, -1))
		chatter.room.broadcast <- struct {
			Message []byte
			Sender  *Chatter
		}{
			Message: mess,
			Sender:  chatter,
		}

	}
}

func (chatter *Chatter) Write() {
	ticker := time.NewTicker(3600 * time.Second)
	defer func() {
		ticker.Stop()
		chatter.conn.Close()
	}()

	log.Println(chatter.room.chatters)

	for {
		select {
		case mess, ok := <-chatter.send:
			if !ok {
				chatter.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			message := Message{}
			writer, err := chatter.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			for c := range chatter.room.chatters {
				log.Println(!bytes.Equal(chatter.wallet.PubKey, c.wallet.PubKey))
				if !bytes.Equal(chatter.wallet.PubKey, c.wallet.PubKey) {

					message.User = chatter.wallet.PubKey
					message.Text = string(mess)
					message.Timestamp = time.Now()
				}
			}

			byted, err := json.Marshal(message)
			if err != nil {
				log.Panic(err)
			}
			writer.Write(byted)

			for i := 0; i < len(chatter.send); i++ {
				writer.Write([]byte{'\n'})
				writer.Write(<-chatter.send)
			}

			if err := writer.Close(); err != nil {
				return
			}

		case <-ticker.C:
			chatter.conn.SetWriteDeadline(time.Now().Add(15 * time.Second))
			if err := chatter.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func Inizialiaze(room *Room, w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("Error during connection upgrade: %s", err)
		return
	}

	chatter := &Chatter{
		room:       room,
		conn:       conn,
		send:       make(chan []byte, 256),
		wallet:     blockchain.NewWallet(),
		blockchain: blockchain.Inizialize(),
	}

	chatter.room.register <- chatter

	// text := "Connected"
	greet := Message{
		User: chatter.wallet.PubKey,
		Text: "Connected to Room " + chatter.room.ID,
		// Timestamp: time.Now(),
	}
	chatter.conn.WriteJSON(greet)

	// Broadcast del messaggio di benvenuto agli altri utenti
	// chatter.room.broadcast <- struct {
	// 	Message []byte
	// 	Sender  *Chatter
	// }{
	// 	Message: []byte(text),
	// 	Sender:  chatter,
	// }

	go chatter.Read()
	go chatter.Write()
}

// func Inizialiaze(room *Room, w http.ResponseWriter, req *http.Request) {

// 	conn, err := upgrader.Upgrade(w, req, nil)
// 	if err != nil {
// 		log.Printf("Error during connection upgrade: %s", err)
// 		return
// 	}

// 	chatter := &Chatter{room: room, conn: conn, send: make(chan []byte, 256),
// 		wallet: blockchain.NewWallet(), blockchain: blockchain.Inizialize()}
// 	chatter.room.register <- chatter

// 	greet := Message{
// 		User:      chatter.wallet.PubKey,
// 		Text:      "Welcome Muddafakka",
// 		Timestamp: time.Now(),
// 	}
// 	chatter.conn.WriteJSON(greet)

// 	go chatter.Read()
// 	go chatter.Write()
// }
