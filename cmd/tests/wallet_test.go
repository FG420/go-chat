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

// func TestWalletNewTx(t *testing.T) {
// 	pino, bc := blockchain.NewWallet()
// 	gino, bc1 := blockchain.NewWallet()

// 	bc.Init()
// 	bc1.Init()

// 	data := "Sei un coglione"
// 	b := pino.SendData(&bc.Blocks[len(bc.Blocks)-1], gino.PubKey, data)
// 	b1 := gino.ReceivedData(&bc1.Blocks[len(bc1.Blocks)-1], pino.PubKey, data)

// 	log.Println("sender -> ", bc.AddBlock(b))
// 	log.Println("receiver -> ", bc1.AddBlock(b1))
// }

func TestSend(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	gino, bc1 := blockchain.NewWallet()

	bc.Init()
	bc1.Init()

	data := "Sei un coglione"
	tx := pino.Send(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b)

	data = "Sei un ebete o abete?"
	tx = pino.Send(gino.PubKey, data)
	b1 := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b1)

	bc.Format()
}

func TestReceive(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	gino, bc1 := blockchain.NewWallet()

	bc.Init()
	bc1.Init()

	data := "Sei un coglione"
	tx := pino.Receive(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b)

	data = "Sei un ebete o abete?"
	tx = pino.Receive(gino.PubKey, data)
	b1 := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b1)

	bc.Format()
}

func TestSendReceive(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	gino, bc1 := blockchain.NewWallet()

	bc.Init()
	bc1.Init()

	data := "Sei un coglione"
	tx := pino.Send(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b)

	data = "Sei un ebete o abete?"
	tx = pino.Receive(gino.PubKey, data)
	b1 := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b1)

	bc.Format()

}
