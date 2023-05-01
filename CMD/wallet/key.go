package wallet

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

const KeyLength = 64 // The length of the hexadecimal string representation of a key

type Key struct {
	Key string
}

func NewKey() (*Key, error) {
	// Generate a new random key using 256 bits of entropy
	keyBytes := make([]byte, 32)
	_, err := rand.Read(keyBytes)
	if err != nil {
		return nil, err
	}
	key := hex.EncodeToString(keyBytes)
	return &Key{Key: key}, nil
}

func NewKeyFromString(str string) (*Key, error) {
	// Ensure that the provided key is a valid hexadecimal string with the correct length
	if len(str) != KeyLength {
		return nil, errors.New("invalid key")
	}
	_, err := hex.DecodeString(str)
	if err != nil {
		return nil, errors.New("invalid key")
	}
	return &Key{Key: str}, nil
}

func (k *Key) String() string {
	return k.Key
}
