package blockchain

type Transaction struct {
	FromPubKey []byte
	ToPubKey   []byte
	Data       []byte
}
