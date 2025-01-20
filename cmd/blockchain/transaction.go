package blockchain

import (
	"bytes"
	"time"
)

type (
	Transaction struct {
		FromPubKey []byte
		ToPubKey   []byte
		Data       any
		Timestamp  int64
	}

	Transactions []*Transaction
)

func (tx *Transaction) NewTransaction(fromKey, toKey []byte, data any) *Transaction {
	return &Transaction{fromKey, toKey, data, time.Now().Unix()}
}

func (tx *Transaction) ValidateTx(key []byte) bool {
	if !bytes.Equal(tx.FromPubKey, key) {
		return false
	} else {
		return true
	}
}

// TODO
func (tx *Transaction) EncypedData() *Transaction {
	return tx
}
