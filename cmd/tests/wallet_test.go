package tests

import (
	"log"
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestNewWallet(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	// gino,bc := blockchain.NewWallet()

	// bc := blockchain.Inizialize()

	log.Println(bc.Init())

	log.Println("Pino -> ", pino)
	// log.Println("Gino -> ", gino)
}

func TestWalletNewTx(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	gino, bc1 := blockchain.NewWallet()

	bc.Init()
	bc1.Init()

	data := "Sei un coglione"
	b := pino.SendData(&bc.Blocks[len(bc.Blocks)-1], gino.PubKey, data)
	b1 := gino.ReceivedData(&bc1.Blocks[len(bc1.Blocks)-1], pino.PubKey, data)

	log.Println("sender -> ", bc.AddBlock(b))
	log.Println("receiver -> ", bc1.AddBlock(b1))
}
