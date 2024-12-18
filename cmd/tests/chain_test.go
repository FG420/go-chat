package tests

import (
	"log"
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestAddBlock(t *testing.T) {
	var bc blockchain.Blockchain
	bc.Init()

	newB := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{{FromPubKey: nil, ToPubKey: nil, Data: "Block 2"}})
	bc.AddBlock(newB)

	newB2 := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{{FromPubKey: nil, ToPubKey: nil, Data: "Block 3"}})
	bc.AddBlock(newB2)

	bc.Format()
	log.Println(bc)
}
