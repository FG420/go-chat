package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	PrevHash     []byte
	Hash         []byte
	Transactions []*Transaction
	Timestamp    int64
}

// func (b *Block) AddTransaction(tx *Transaction) *Block {
// }

func (b *Block) Serialize() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		log.Panic("Error during encoding")
	}

	return buf.Bytes()
}

func Deserialize(data []byte) *Block {
	var b Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&b)
	if err != nil {
		log.Panic("error during decoding")
	}

	return &b
}

func CreateBlock(prevBlock *Block, tx *Transaction) *Block {
	var txs []*Transaction
	txs = append(txs, tx)
	block := &Block{prevBlock.Hash, []byte{}, txs, time.Now().Unix()}
	p := NewProof(block)
	hash := p.Run()
	block.Hash = hash

	return block
}

func GenesisBlock() *Block {
	gen := Block{
		PrevHash:     nil,
		Hash:         nil,
		Transactions: []*Transaction{{nil, nil, "Blockchain Inizialized"}},
		Timestamp:    time.Now().Unix(),
	}

	sGb := string(gen.PrevHash) + fmt.Sprint(gen.Timestamp) + fmt.Sprint(len(gen.Transactions))
	hash := sha256.Sum256([]byte(sGb))
	gen.Hash = hash[:]

	return &gen
}
