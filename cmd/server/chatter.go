package server

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/GF420/go-chat/cmd/blockchain"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type (
	Chatter struct {
		room       *Room
		conn       *websocket.Conn
		send       chan []byte
		wallet     *blockchain.Wallet
		blockchain *blockchain.Blockchain
	}

	Welcome struct {
		Greeting string
		User     []byte
	}

	MessageType struct {
		User      []byte
		Text      string
		Timestamp time.Time
	}
)

func (chatter *Chatter) Read() {
	defer func() {
		chatter.room.unregister <- chatter
		chatter.conn.Close()
	}()

	for {

		t, mess, err := chatter.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			} else {
				log.Panic(err)
			}
			break
		}
		log.Println(t, string(mess))

		mess = bytes.TrimSpace(bytes.Replace(mess, []byte{'\n'}, []byte{' '}, -1))
		chatter.room.broadcast <- mess
	}
}

func (chatter *Chatter) Write() {
	ticker := time.NewTicker(180 * time.Second)
	defer func() {
		ticker.Stop()
		chatter.conn.Close()
	}()

	// log.Println("Chatter Wallet -> ", chatter.wallet)

	log.Println(chatter.room.chatters)
	// b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], &blockchain.Transaction{})
	// bc.AddBlock(b)

	for {
		select {
		case mess, ok := <-chatter.send:
			if !ok {
				chatter.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := chatter.conn.WriteJSON(MessageType{User: chatter.wallet.PubKey, Text: string(mess), Timestamp: time.Now()}); err != nil {
				return
			}
			// writer, err := chatter.conn.NextWriter(websocket.TextMessage)
			// if err != nil {
			// 	return
			// }
			// writer.Write(mess)
			// tx := chatter.wallet.Send(wallet.PubKey, <-chatter.send)
			// b.AddTransaction(tx)
			// log.Println("Block -> ", b)

			for cha := range chatter.room.chatters {
				if cha.wallet != chatter.wallet {
					tx := chatter.wallet.Send(cha.wallet.PubKey, mess)
					chatter.blockchain.AddBlock(blockchain.CreateBlock(
						&chatter.blockchain.Blocks[len(chatter.blockchain.Blocks)-1], tx))
				}
			}

			// log.Println(chatter.blockchain)

			// for i := 0; i < len(chatter.send); i++ {
			// 	writer.Write([]byte{'\n'})
			// 	writer.Write(<-chatter.send)
			// }

			// if err := writer.Close(); err != nil {
			// 	return
			// }
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

	chatter := &Chatter{room: room, conn: conn, send: make(chan []byte, 256), wallet: blockchain.NewWallet(), blockchain: blockchain.Inizialize()}
	chatter.room.register <- chatter

	greet := Welcome{
		Greeting: "Welcome",
		User:     chatter.wallet.PubKey,
	}

	chatter.conn.WriteJSON(greet)
	go chatter.Read()
	go chatter.Write()
}
