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

func NewWallet() *Wallet {
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	public := append(private.X.Bytes(), private.Y.Bytes()...)

	return &Wallet{public, private}
}

// func GetWallets() Wallets {

// }
