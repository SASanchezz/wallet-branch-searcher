package main

import (
	"wallet-branch-blockchain/src/api"
	"wallet-branch-blockchain/src/bootstrap"
	"wallet-branch-blockchain/src/listener"
)

const amount = 100

func main() {
	bootstrap.LoadEnv()
	// generateTransactions(amount)

	go api.Run()

	listener.Listen()
}
