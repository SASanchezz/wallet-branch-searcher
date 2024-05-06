package main

import (
	"wallet-branch-blockchain/src/core"
	"wallet-branch-blockchain/src/random"
	"wallet-branch-blockchain/src/transaction"
)

func generateTransactions(amount int) {
	ts := transaction.New()

	transactionArgs := random.GetRandomTransactions(amount)
	for _, transactionArg := range *transactionArgs {
		transaction := ts.GenerateTransaction(transactionArg)
		transaction.Hash = core.GetHash(&transaction)
		ts.SaveTransaction(transaction)
	}
}
