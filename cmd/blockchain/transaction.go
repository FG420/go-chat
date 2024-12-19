package blockchain

type (
	Transaction struct {
		FromPubKey []byte
		ToPubKey   []byte
		Data       any
	}
	Transactions []*Transaction
)

func (tx *Transaction) NewTransaction(fromKey, toKey []byte, data any) *Transaction {
	return &Transaction{fromKey, toKey, data}
}

func (tx *Transaction) ValidateTx() bool {
	return false
}
