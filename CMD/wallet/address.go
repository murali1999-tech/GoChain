package wallet

import (
	"encoding/hex"
	"errors"
)

const AddressLength = 64 // The length of the hexadecimal string representation of an address

type Address struct {
	Address string
}

func NewAddress(addr string) (*Address, error) {
	// Ensure that the provided address is a valid hexadecimal string with the correct length
	if len(addr) != AddressLength {
		return nil, errors.New("invalid address")
	}
	_, err := hex.DecodeString(addr)
	if err != nil {
		return nil, errors.New("invalid address")
	}
	return &Address{Address: addr}, nil
}

func (a *Address) String() string {
	return a.Address
}
