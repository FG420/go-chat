package tests

import (
	"log"
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestGenesis(t *testing.T) {
	gen := blockchain.GenesisBlock()

	log.Println(gen)
}

func TestCreateNewBlock(t *testing.T) {
	gen := blockchain.GenesisBlock()
	newB := blockchain.CreateBlock(gen, append(gen.Transactions, blockchain.Transaction{FromPubKey: nil, ToPubKey: nil, Data: "New Block"}))

	log.Println("New Block -> ", newB)
}
