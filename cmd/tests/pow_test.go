package tests

import (
	"log"
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestProofOfWork(t *testing.T) {
	b := blockchain.GenesisBlock()
	pow := blockchain.NewProof(b)
	hash := pow.Run()

	log.Println("POW -> ", pow)
	log.Println("HASH -> ", hash)
}
