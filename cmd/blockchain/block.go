package blockchain

import (
	"time"
)

type Block struct {
	PrevHash     []byte
	Hash         []byte
	Transactions []*Transaction
	Timestamp    int64
}

func CreateBlock(prevHash []byte, txs []*Transaction) *Block {
	block := &Block{prevHash, []byte{}, txs, time.Now().Unix()}
	p := NewProof(block)
	hash := p.Run()
	block.Hash = hash

	return block
}

func GenesisBlock() *Block {
	return CreateBlock(nil, []*Transaction{})
}
