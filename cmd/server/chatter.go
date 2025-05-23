package server

import (
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
		send       chan Message
		wallet     *blockchain.Wallet
		blockchain *blockchain.Blockchain
	}

	Message struct {
		User      []byte `json:"user"`
		Text      string `json:"text"`
		Timestamp int64  `json:"timestamp"`
	}
)

func (chatter *Chatter) Read() {
	defer func() {
		chatter.room.unregister <- chatter
		chatter.conn.Close()
	}()

	for {
		var message Message

		err := chatter.conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		log.Println(message)
		chatter.room.broadcast <- message
	}
}

// func (chatter *Chatter) Read() {
// 	defer func() {
// 		chatter.room.unregister <- chatter
// 		chatter.conn.Close()
// 	}()

// 	for {
// 		_, mess, err := chatter.conn.ReadMessage()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("error: %v", err)
// 				chatter.conn.WriteJSON(Message{User: chatter.wallet.PubKey, Text: "This User left the room", Timestamp: time.Now()})
// 			} else {
// 				log.Printf("error: %v", err)
// 				chatter.conn.WriteJSON(Message{User: chatter.wallet.PubKey, Text: "This User left the room", Timestamp: time.Now()})
// 			}
// 			break
// 		}

// 		mess = bytes.TrimSpace(bytes.Replace(mess, []byte{'\n'}, []byte{' '}, -1))

// 		log.Println(string(mess))

// 		// if err := json.Unmarshal(mess, &mex); err != nil {
// 		// 	log.Panic(err)
// 		// }

// 		message := Message{User: chatter.wallet.PubKey, Text: string(mess), Timestamp: time.Now()}
// 		log.Println(message)

// 		chatter.room.broadcast <- message

// 	}
// }

func (chatter *Chatter) Write() {
	// var message Message

	ticker := time.NewTicker(9600 * time.Second)
	defer func() {
		ticker.Stop()
		chatter.conn.Close()
	}()

	for {
		select {
		case mess, ok := <-chatter.send:
			if !ok {
				chatter.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// log.Println(mess)

			// message.User = chatter.wallet.PubKey
			// message.Text = string(mess)
			// message.Timestamp = time.Now()

			// if err := chatter.conn.WriteJSON(mess); err != nil {
			// 	log.Panic(err)
			// }

			writer, err := chatter.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			// message.User = chatter.wallet.PubKey
			// message.Text = string(mess)
			// message.Timestamp = time.Now()

			byted, err := json.Marshal(mess)
			if err != nil {
				log.Panic(err)
			}
			writer.Write(byted)

			for i := 0; i < len(chatter.send); i++ {
				writer.Write([]byte{'\n'})
				// writer.Write(<-chatter.send)
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
		room: room,
		conn: conn,
		// send:       make(chan []byte, 256)
		send:       make(chan Message),
		wallet:     blockchain.NewWallet(),
		blockchain: blockchain.Inizialize(),
	}

	chatter.room.register <- chatter

	greet := Message{
		User:      chatter.wallet.PubKey,
		Text:      "Connected to Room " + chatter.room.ID,
		Timestamp: time.Now().Unix(),
	}
	chatter.conn.WriteJSON(greet)

	go chatter.Read()
	go chatter.Write()
}
