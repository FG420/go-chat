package blockchain

type Transaction struct {
	FromPubKey []byte
	ToPubKey   []byte
	Data       any
}

func (tx *Transaction) NewTransaction(fromKey, toKey []byte, data any) *Transaction {
	return &Transaction{fromKey, toKey, data}
}
