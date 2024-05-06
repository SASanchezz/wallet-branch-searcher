package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GTransaction(dbTx neo4j.ManagedTransaction, txHash *common.Hash) *neo4j.Record {
	params := map[string]interface{}{
		"hash": txHash.ToString(),
	}
	template := "MATCH (t:BaseTransaction) " +
		"WHERE t.hash = $hash " +
		"RETURN t " +
		"UNION " +
		"MATCH (t:Transaction) " +
		"WHERE t.hash = $hash " +
		"RETURN t"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_transaction.txt").
		Build()

	if result := query.Run(); len(result) == 0 {
		return nil
	} else {
		return result[0]
	}
}
