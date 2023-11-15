package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(data string) string {
	hashInBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashInBytes[:])
}

func ConcatenateStrings(strings ...string) string {
	var result string
	for _, str := range strings {
		result += str
	}
	return result
}
