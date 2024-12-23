package tests

import (
	"log"
	"testing"
	"time"

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

func TestSend(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	gino, bc1 := blockchain.NewWallet()

	bc.Init()
	bc1.Init()

	data := "Sei un mona"
	tx := pino.Send(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b)

	data = "Sei un pino o un abete?"
	tx = pino.Send(gino.PubKey, data)
	b1 := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)
	bc.AddBlock(b1)

	bc.Format()
}

func TestMultipleSend(t *testing.T) {
	pino, bc := blockchain.NewWallet()
	gino, _ := blockchain.NewWallet()

	bc.Init()
	data := "mona coglione"
	tx := pino.Send(gino.PubKey, data)
	b := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], tx)

	time.Sleep(500)
	data1 := "cazzone che sei, scemo"
	tx1 := pino.Send(gino.PubKey, data1)
	b.AddTransaction(tx1)

	time.Sleep(500)
	data2 := "però ti voglio pene!"
	tx2 := pino.Send(gino.PubKey, data2)
	b.AddTransaction(tx2)

	bc.AddBlock(b)
	bc.Format()

}
