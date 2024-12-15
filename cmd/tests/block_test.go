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
	newB := blockchain.CreateBlock(gen, []blockchain.Transaction{{FromPubKey: nil, ToPubKey: nil, Data: "New Block"}})

	log.Println("New Block -> ", newB)
}

func TestSerializeBlock(t *testing.T) {
	gen := blockchain.GenesisBlock()
	newB := blockchain.CreateBlock(gen, []blockchain.Transaction{{FromPubKey: nil, ToPubKey: nil, Data: "Block 2"}})
	hashedB := newB.Serialize()

	log.Println("Hashed Block -> ", hashedB)
}

func TestDeserializeBlock(t *testing.T) {
	gen := blockchain.GenesisBlock()

	encodedB := gen.Serialize()
	log.Println("Encoded Block -> ", encodedB)

	decodedB := blockchain.Deserialize(encodedB)
	log.Println("Decoded Block -> ", decodedB)
}
