package tests

import (
	"log"
	"testing"
	"time"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestNewWallet(t *testing.T) {
	bc := blockchain.Inizialize()
	pino := blockchain.NewWallet()
	// gino,bc := blockchain.NewWallet()

	// bc := blockchain.Inizialize()

	log.Println(bc)

	log.Println("Pino -> ", pino)
	// log.Println("Gino -> ", gino)
}

func TestSend(t *testing.T) {
	bc := blockchain.Inizialize()
	bc1 := blockchain.Inizialize()
	pino := blockchain.NewWallet()
	gino := blockchain.NewWallet()

	data := "Chi sei?"
	tx := pino.Send(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b)

	data = "Sei un pino o un abete?"
	tx = pino.Send(gino.PubKey, data)
	b1 := blockchain.CreateBlock(&bc1.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b1)

	bc.Format()
}

func TestMultipleSend(t *testing.T) {
	bc := blockchain.Inizialize()
	pino := blockchain.NewWallet()
	gino := blockchain.NewWallet()

	data := "ciao"
	tx := pino.Send(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)

	time.Sleep(500)
	data1 := "che scemo sei"
	tx1 := pino.Send(gino.PubKey, data1)
	b.AddTransaction(tx1, pino.PubKey)

	time.Sleep(500)
	data2 := "però ti voglio bene!"
	tx2 := pino.Send(gino.PubKey, data2)
	b.AddTransaction(tx2, pino.PubKey)

	bc.AddBlock(b)
	bc.Format()

}
