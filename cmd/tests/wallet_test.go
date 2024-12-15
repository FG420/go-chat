package tests

import (
	"log"
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestNewWallet(t *testing.T) {
	pino := blockchain.NewWallet()
	gino := blockchain.NewWallet()
	log.Println("Pino -> ", pino)
	log.Println("Gino -> ", gino)
}
