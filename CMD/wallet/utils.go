package wallet

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

// CalculateHash returns the SHA-256 hash of the given data
func CalculateHash(data interface{}) string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(hash[:])
}

// ContainsString checks if a string is present in a slice of strings
func ContainsString(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
