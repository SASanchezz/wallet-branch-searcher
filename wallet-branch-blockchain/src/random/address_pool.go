package random

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"wallet-branch-blockchain/src/common"
)

func generateAddressPool(amount int) []*common.Address {
	addressPool := make([]*common.Address, amount)

	for i := 0; i < amount; i++ {
		addressPool[i] = generateRandomAddress()
	}

	return addressPool
}

func generateRandomAddress() *common.Address {
	addressLength := 18

	randomBytes := make([]byte, addressLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	rawAddress := hex.EncodeToString(randomBytes)
	address := "0x" + rawAddress

	return common.StringToAddress(address)
}

var mock_addresses = []*common.Address{
	common.StringToAddress(strings.Repeat("0", 20)),
	common.StringToAddress(strings.Repeat("1", 20)),
	common.StringToAddress(strings.Repeat("2", 20)),
	common.StringToAddress(strings.Repeat("3", 20)),
	common.StringToAddress(strings.Repeat("4", 20)),
	common.StringToAddress(strings.Repeat("5", 20)),
}
