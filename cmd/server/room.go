package server

import (
	"bytes"
	"log"
	"strconv"
	"time"

	"math/rand"

	"github.com/GF420/go-chat/cmd/blockchain"
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

			b := blockchain.CreateBlock(&c.blockchain.Blocks[len(c.blockchain.Blocks)-1], c.blockchain.Blocks[len(c.blockchain.Blocks)-1].Transactions[len(c.blockchain.Blocks[len(c.blockchain.Blocks)-1].Transactions)-1])
			for _, chas := range chatters {
				tx := c.wallet.Send(chas.wallet.PubKey, message.Text)
				b.AddTransaction(tx, c.wallet.PubKey)
			}

			for _, tx := range b.Transactions {
				log.Println(tx.Data)
			}
		}
	}
}
