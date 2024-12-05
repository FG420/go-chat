package blockchain

import (
	"math/big"
	"math/rand"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	source := rand.Intn(100 - 20)
	t := big.NewInt(1)
	t.Lsh(t, uint(source))

	return &ProofOfWork{b, t}
}
