package server

import (
	"strconv"
	"time"

	"math/rand"
)

type Room struct {
	ID        string
	chatters  map[*Chatter]bool
	broadcast chan struct {
		Message []byte
		Sender  *Chatter
	}
	register   chan *Chatter
	unregister chan *Chatter
}

func NewRoom() *Room {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	return &Room{
		ID: strconv.Itoa(r.Int()),
		broadcast: make(chan struct {
			Message []byte
			Sender  *Chatter
		}),
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

// 		case messageAndSender := <-r.broadcast:
// 			for chatter := range r.chatters {
// 				if chatter != messageAndSender.Sender {
// 					select {
// 					case chatter.send <- messageAndSender.Message:
// 					default:
// 						close(chatter.send)
// 						delete(r.chatters, chatter)
// 					}
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

		case messageAndSender := <-r.broadcast:
			for chatter := range r.chatters {
				if chatter != messageAndSender.Sender {
					select {
					case chatter.send <- messageAndSender.Message:
					default:
						close(chatter.send)
						delete(r.chatters, chatter)
					}
				}
			}
		}
	}
}
