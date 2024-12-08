package tests

import (
	"log"
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestNewWallet(t *testing.T) {
	pino := blockchain.NewWallet()
	log.Println(pino.PubKey)
	log.Println(pino.PrivKey)
}
