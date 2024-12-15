package blockchain

import (
	"log"
)

type Blockchain struct {
	LastHash []byte
	Blocks   []Block
}

func (chain *Blockchain) AddBlock(b *Block) *Blockchain {
	var block Block = *b

	chain.Blocks = append(chain.Blocks, block)
	chain.LastHash = block.Hash

	return chain
}

func (chain *Blockchain) Format() {
	for _, block := range chain.Blocks {
		log.Printf("\nLast Hash ->\t%x\nBlocks:\n\t- PrevHash ->\t%x\n\t- Hash ->\t%x\n\t- Timestamp ->\t%d\n\t- Transactions ->\t%s",
			chain.LastHash,
			block.PrevHash,
			block.Hash,
			block.Timestamp,
			block.Transactions)
	}
}

func Inizialized() *Blockchain {
	// var chain Blockchain
	// gen := GenesisBlock()
	return &Blockchain{}
}
