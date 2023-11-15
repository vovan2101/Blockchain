package blockchain

import (
	"fmt"
)

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

func calculateTransactionHash(transaction Transaction) string {
	data := fmt.Sprintf("%s%s%f", transaction.Sender, transaction.Recipient, transaction.Amount)
	return CalculateHash(data)
}
