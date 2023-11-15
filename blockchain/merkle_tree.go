package blockchain

import (
	"fmt"
)

func CreateMerkleTree(transactions []Transaction) string {
	var merkleTree []string

	for _, tx := range transactions {
		merkleTree = append(merkleTree, CalculateHash(fmt.Sprintf("%s%s%f", tx.Sender, tx.Recipient, tx.Amount)))
	}

	for len(merkleTree) > 1 {
		var newLevel []string

		for i := 0; i < len(merkleTree); i += 2 {
			hashPair := merkleTree[i]

			if i+1 < len(merkleTree) {
				hashPair += merkleTree[i+1]
			}

			newLevel = append(newLevel, CalculateHash(hashPair))
		}

		merkleTree = newLevel
	}
	return merkleTree[0]
}
