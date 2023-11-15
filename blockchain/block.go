package blockchain

import (
	"fmt"
	"time"
)

type Block struct {
	Index       int
	Timestamp   int64
	Transaction []Transaction
	PrevHash    string
	Hash        string
}

var Blockchain []Block

func CalculateBlockHash(block Block) string {
	var transactionHashes []string

	for _, tx := range block.Transaction {
		transactionHash := calculateTransactionHash(tx)
		transactionHashes = append(transactionHashes, transactionHash)
	}

	record := fmt.Sprintf("%d%d%s%s", block.Index, block.Timestamp, block.PrevHash, ConcatenateStrings(transactionHashes...))
	return CalculateHash(record)
}

func GenerateBlock(prevBlock Block, transactions []Transaction) Block {
	var newBlock Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Transaction = transactions
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = CreateMerkleTree(transactions)

	return newBlock
}
