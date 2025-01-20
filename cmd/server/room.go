package server

import (
	"bytes"
	"log"
	"strconv"
	"time"

	"math/rand"
)

type Room struct {
	ID       string
	chatters map[*Chatter]bool
	// broadcast chan []byte
	// broadcast chan struct {
	// 	Message []byte
	// 	Sender  *Chatter
	// }
	broadcast  chan Message
	register   chan *Chatter
	unregister chan *Chatter
}

func NewRoom() *Room {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	return &Room{
		ID:        strconv.Itoa(r.Int()),
		broadcast: make(chan Message),
		// broadcast: make(chan struct {
		// 	Message []byte
		// 	Sender  *Chatter
		// }),
		// broadcast:  make(chan []byte),
		register:   make(chan *Chatter),
		unregister: make(chan *Chatter),
		chatters:   make(map[*Chatter]bool),
	}
}

// func (r *Room) Run() {
// 	for {
// 		select {
// 		case chatter := <-r.register:
// 			r.chatters[chatter] = true

// 		case chatter := <-r.unregister:
// 			if _, ok := r.chatters[chatter]; ok {
// 				delete(r.chatters, chatter)
// 				close(chatter.send)
// 			}

// 		case mess := <-r.broadcast:
// 			for chatter := range r.chatters {
// 				select {
// 				case chatter.send <- mess:
// 				default:
// 					close(chatter.send)
// 					delete(r.chatters, chatter)
// 				}

// 			}
// 		}
// 	}
// }

func (r *Room) Run() {
	for {
		select {
		case chatter := <-r.register:
			r.chatters[chatter] = true

		case chatter := <-r.unregister:
			if _, ok := r.chatters[chatter]; ok {
				delete(r.chatters, chatter)
				close(chatter.send)
			}

		case message := <-r.broadcast:
			var c *Chatter
			var chatters []*Chatter
			for chatter := range r.chatters {
				if !bytes.Equal(chatter.wallet.PubKey, message.User) {
					chatters = append(chatters, chatter)
					select {
					case chatter.send <- message:
					default:
						close(chatter.send)
						delete(r.chatters, chatter)
					}
				} else {
					c = chatter
				}
			}

			for _, chas := range chatters {
				tx := c.wallet.Send(chas.wallet.PubKey, message.Text)
				log.Println(tx)
			}

		}
	}
}
