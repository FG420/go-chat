package tests

import (
	"testing"

	"github.com/GF420/go-chat/cmd/blockchain"
)

func TestAddBlock(t *testing.T) {
	bc := blockchain.Inizialized()

	genB := blockchain.GenesisBlock()
	bc.AddBlock(genB)

	newB := blockchain.CreateBlock(genB, []blockchain.Transaction{{FromPubKey: nil, ToPubKey: nil, Data: "Block 2"}})
	bc.AddBlock(newB)

	newB2 := blockchain.CreateBlock(&bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{{FromPubKey: nil, ToPubKey: nil, Data: "Block 3"}})
	bc.AddBlock(newB2)

	bc.Format()

}
