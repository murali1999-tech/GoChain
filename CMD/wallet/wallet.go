package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

type Wallet struct {
	Address  string
	PrivateKey *ecdsa.PrivateKey
}

func NewWallet() (*Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	address := hex.EncodeToString(sha256.Sum256(pubKey)[:])
	return &Wallet{
		Address:    address,
		PrivateKey: privateKey,
	}, nil
}

func (w *Wallet) Sign(msg []byte) (string, error) {
	if w.PrivateKey == nil {
		return "", errors.New("private key is missing")
	}

	hash := sha256.Sum256(msg)
	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, hash[:])
	if err != nil {
		return "", err
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return hex.EncodeToString(signature), nil
}

func (w *Wallet) Verify(msg []byte, signature string) bool {
	if w.PrivateKey == nil {
		return false
	}

	signatureBytes, err := hex.DecodeString(signature)
	if err != nil || len(signatureBytes) != 64 {
		return false
	}

	rBytes := signatureBytes[:32]
	sBytes := signatureBytes[32:]

	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	hash := sha256.Sum256(msg)
	return ecdsa.Verify(&w.PrivateKey.PublicKey, hash[:], r, s)
}

