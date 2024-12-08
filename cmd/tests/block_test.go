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
