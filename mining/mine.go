package mining

import (
	"blockchainproject/blockchain"
	"time"
)

func MineBlock(prevBlock blockchain.Block, transactions []blockchain.Transaction) blockchain.Block {
	coinbaseTransaction := blockchain.Transaction{
		Sender:    "0",
		Recipient: "minerReward",
		Amount:    10.0,
	}

	transactions = append([]blockchain.Transaction{coinbaseTransaction}, transactions...)

	newBlock := blockchain.GenerateBlock(prevBlock, transactions)
	newBlock.Hash = mineBlockHash(newBlock)

	return newBlock
}

func mineBlockHash(block blockchain.Block) string {
	for {
		hash := blockchain.CalculateBlockHash(block)
		if true {
			return hash
		}
		block.Timestamp = time.Now().Unix()
	}
}
