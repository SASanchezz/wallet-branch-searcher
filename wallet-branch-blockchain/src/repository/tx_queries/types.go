package tx_queries

import (
	"math/big"
	"wallet-branch-blockchain/src/common"
)

type GBranchParams struct {
	From   *common.Address
	To     *common.Address
	Limit  *int64
	Before *int64
	After  *int64
}

type NodeData struct {
	From                 *string  `json:"from"`
	To                   *string  `json:"to"`
	Hash                 *string  `json:"hash"`
	Gas                  *uint64  `json:"gas"`
	GasPrice             *big.Int `json:"gasPrice"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
	Value                *big.Int `json:"value"`
	Timestamp            *uint64  `json:"timestamp"`
	Nonce                *uint64  `json:"nonce"`
}

type RelationshipData struct {
	From *common.Address `json:"from"`
	To   *common.Address `json:"to"`
}

type InterrelatedAddresses struct {
	FromAddresses []string `json:"from"`
	ToAddresses   []string `json:"to"`
}

type Branch []*NodeData

type Branches []*Branch
