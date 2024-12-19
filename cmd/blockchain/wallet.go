package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type (
	Wallet struct {
		PubKey  []byte
		PrivKey *ecdsa.PrivateKey
	}

	// Wallets struct {
	// 	Wallets []Wallet
	// }
)

func NewWallet() (*Wallet, *Blockchain) {
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	public := append(private.X.Bytes(), private.Y.Bytes()...)

	return &Wallet{public, private}, &Blockchain{}
}

func (w *Wallet) Send(toPubKey []byte, data any) *Transaction {
	var tx Transaction
	return tx.NewTransaction(w.PubKey, toPubKey, data)
}

func (w *Wallet) Receive(receiverKey []byte, data any) *Transaction {
	var tx Transaction
	return tx.NewTransaction(receiverKey, w.PubKey, data)
}

// func (w *Wallet) SendData(prevB *Block, receiverKey []byte, data any) *Block {
// 	return CreateBlock(prevB, []Transaction{{FromPubKey: w.PubKey, ToPubKey: receiverKey, Data: data}})
// }

// func (w *Wallet) ReceivedData(prevB *Block, senderKey []byte, data any) *Block {
// 	return CreateBlock(prevB, []Transaction{{FromPubKey: senderKey, ToPubKey: w.PubKey, Data: data}})
// }

// func GetWallets() Wallets {

// }
