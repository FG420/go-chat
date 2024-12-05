package blockchain

type Block struct {
	Hash        []byte
	PrevHash    []byte
	Transaction *Transaction
	Timestamp   int64
}

// func AddBlock(prevHash []byte, tx *Transaction) *Block {
// 	block := &Block{[]byte{}, prevHash, tx, time.Now().Unix()}
// 	pow := NewProof(block)

// 	block.Hash =
// }
