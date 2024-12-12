package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	PrevHash     []byte
	Hash         []byte
	Transactions []Transaction
	Timestamp    int64
}

func CreateBlock(prevBlock *Block, txs []Transaction) *Block {
	block := &Block{prevBlock.Hash, []byte{}, txs, time.Now().Unix()}
	p := NewProof(block)
	hash := p.Run()
	block.Hash = hash

	return block
}

func GenesisBlock() *Block {
	gen := Block{
		PrevHash:     make([]byte, 0),
		Hash:         nil,
		Transactions: []Transaction{{nil, nil, "Genesis Block"}},
		Timestamp:    time.Now().Unix(),
	}

	sGb := string(gen.PrevHash) + fmt.Sprint(gen.Timestamp) + string(len(gen.Transactions))
	hash := sha256.Sum256([]byte(sGb))
	gen.Hash = hash[:]

	return &gen
}
