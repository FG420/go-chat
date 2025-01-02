package server

import (
	"bytes"
	"log"
	"net/http"
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

	Chatter struct {
		room *Room
		conn *websocket.Conn
		send chan []byte
	}
)

func (wsh WebSocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// var res Response

	room := newRoom()
	go room.run()

	conn, err := wsh.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("Error during connection upgrade: %s", err)
		return
	}
	defer conn.Close()

	chatter := &Chatter{room: room, conn: conn, send: make(chan []byte, 256)}

	chatter.room.register <- chatter

	log.Println(chatter.room.register)
	// for {
	// 	chatter.Write()
	// 	chatter.Read()
	// }

	// log.Println(chatter.room.chatters)

	// res.Text = "U joined a room!"
	// res.Timestamp = time.Now()
	// conn.WriteJSON(res)

	// log.Println("remote address: ", conn.RemoteAddr().String())

	// for {
	// 	mess := readMessages(conn, res)
	// 	if strings.Trim(string(mess), "\n") == "" {
	// 		res.Text = "Say something!"
	// 		res.Timestamp = time.Now()
	// 		err := conn.WriteJSON(res)
	// 		if err != nil {
	// 			log.Printf("Error: %s", err)
	// 			return
	// 		}
	// 		continue
	// 	}
	// 	resClient(conn, res, mess)
	// }
}

// func readMessages(conn *websocket.Conn, res Response) []byte {
// 	mt, mess, err := conn.ReadMessage()
// 	if err != nil {
// 		log.Printf("Error during sending messages: %s", err)
// 		return nil
// 	}

// 	if mt == websocket.BinaryMessage {
// 		// err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))
// 		res.Text = "Server doesn't support binary messages"
// 		res.Timestamp = time.Now()
// 		err = conn.WriteJSON(res)
// 		if err != nil {
// 			log.Printf("Error: %s", err)
// 		}
// 		return nil
// 	}

// 	log.Printf("Received message %s", string(mess))
// 	return mess
// }

// func resClient(conn *websocket.Conn, res Response, mess []byte) {
// 	log.Println("start responding to client...")
// 	res.Text = string(mess)
// 	res.Timestamp = time.Now()
// 	err := conn.ReadJSON(res)
// 	if err != nil {
// 		log.Printf("Error when respondin to client: %s", err)
// 		return
// 	}
// }

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
			}
			break
		}

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
	log.Println("ciao")

	// for chatter := range chatter.room.chatters {
	// 	mess, ok := <-chatter.send
	// 	if !ok {
	// 		chatter.conn.WriteMessage(websocket.CloseMessage, []byte{})
	// 		return
	// 	}

	// 	writer, err := chatter.conn.NextWriter(websocket.TextMessage)
	// 	if err != nil {
	// 		return
	// 	}
	// 	writer.Write(mess)

	// 	for i := 0; i < len(chatter.send); i++ {
	// 		writer.Write([]byte{'\n'})
	// 		writer.Write(<-chatter.send)
	// 	}
	// }

	for {
		select {
		case mess, ok := <-chatter.send:
			if !ok {
				chatter.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := chatter.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			writer.Write(mess)
			log.Println(string(mess))

			for i := 0; i < len(chatter.send); i++ {
				writer.Write([]byte{'\n'})
				writer.Write(<-chatter.send)
				log.Println(<-chatter.send)
			}
		case <-ticker.C:
			chatter.conn.SetWriteDeadline(time.Now().Add(15 * time.Second))
			if err := chatter.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
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
