package blockchain

type Transaction struct {
	FromPubKey []byte
	ToPubKey   []byte
	Data       []byte
}

func NewTransaction(fromKey, toKey, data []byte) *Transaction {
	return &Transaction{fromKey, toKey, data}
}
