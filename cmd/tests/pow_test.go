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

	log.Println("POW Block -> ", pow.Block)
	log.Println("POW Target -> ", pow.Target)
	log.Println("HASH -> ", hash)
}
