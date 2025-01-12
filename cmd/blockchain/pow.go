package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
	"math/rand"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func (pow *ProofOfWork) Run() []byte {
	var intHash big.Int
	var hash [32]byte
	random := rand.Intn(50)

	for random < math.MaxInt64 {
		data := pow.InitData(random)
		hash := sha256.Sum256(data)
		// fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			return hash[:]
		} else {
			random++
		}
	}

	return hash[:]
}

func (pow *ProofOfWork) InitData(num int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Hash,
			{byte(pow.Block.Timestamp)},
			ToHex(int64(num)),
		}, []byte{})
	return data
}

func NewProof(b *Block) *ProofOfWork {
	source := rand.Intn(20)
	t := big.NewInt(1)
	// log.Print(num)
	// log.Println(t, source)
	t.Lsh(t, uint(256-source))
	log.Println(t)

	return &ProofOfWork{b, t}
}

func ToHex(num int64) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return b.Bytes()
}
