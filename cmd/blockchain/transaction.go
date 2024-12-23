package blockchain

import "time"

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

func (tx *Transaction) ValidateTx() bool {
	return false
}
