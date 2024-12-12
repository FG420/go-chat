package blockchain

type Transaction struct {
	FromPubKey []byte
	ToPubKey   []byte
	Data       string
}

func NewTransaction(fromKey, toKey []byte, data string) *Transaction {
	return &Transaction{fromKey, toKey, data}
}
