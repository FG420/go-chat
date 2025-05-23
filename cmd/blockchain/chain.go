package blockchain

import (
	"bytes"
	"log"
)

type Blockchain struct {
	LastHash []byte
	Blocks   []Block
}

func (chain *Blockchain) AddBlock(b *Block) *Blockchain {
	var block Block = *b

	if !bytes.Equal(chain.LastHash, b.PrevHash) {
		log.Panic("Error the last hash and the previous block hash do not match!")
	}
	chain.Blocks = append(chain.Blocks, block)
	chain.LastHash = b.Hash
	return chain
}

func (chain *Blockchain) Format() {
	for _, block := range chain.Blocks {
		log.Printf("\nBlock:\n\t- PrevHash ->\t%x\n\t- Hash ->\t%x\n\t- Timestamp ->\t%d",
			block.PrevHash,
			block.Hash,
			block.Timestamp)

		for _, tx := range block.Transactions {
			log.Printf("\n\t- Transactions:\n\t\t- FromKey ->\t%x\n\t\t- ToKey ->\t%x\n\t\t- Data ->\t%s\n\t\t- Timestamp ->\t%d",
				tx.FromPubKey,
				tx.ToPubKey,
				tx.Data,
				tx.Timestamp)
		}
	}
}

func Inizialize() *Blockchain {
	var chain Blockchain
	return chain.AddBlock(GenesisBlock())
}

func (chain *Blockchain) Init() *Blockchain {
	return chain.AddBlock(GenesisBlock())
}
