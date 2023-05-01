package wallet

import (
	"crypto/rand"
	"errors"
	"github.com/tyler-smith/go-bip39"
)

const MnemonicLength = 24 // The number of words in a BIP-39 mnemonic phrase

type Mnemonic struct {
	Phrase string
}

func NewMnemonic() (*Mnemonic, error) {
	// Generate a new random BIP-39 mnemonic phrase with 128 bits of entropy
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}
	phrase, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	return &Mnemonic{Phrase: phrase}, nil
}

func NewMnemonicFromString(str string) (*Mnemonic, error) {
	// Ensure that the provided string is a valid BIP-39 mnemonic phrase
	if !bip39.IsMnemonicValid(str) {
		return nil, errors.New("invalid mnemonic phrase")
	}
	return &Mnemonic{Phrase: str}, nil
}

func (m *Mnemonic) String() string {
	return m.Phrase
}
