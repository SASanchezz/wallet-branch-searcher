package random

import (
	"math/big"
	"math/rand"
	"time"
	"wallet-branch-blockchain/src/common"
)

var addressPool = generateAddressPool(300)

func GetRandomAddress() *common.Address {
	return addressPool[rand.Intn(len(addressPool))]
}

func GetRandomUint64() *uint64 {
	value := rand.Uint64()
	return &value
}

func GetRandomBigInt() *big.Int {
	n := int64(rand.Int())

	return big.NewInt(n)
}

func GetRandomTimestamp() *uint64 {
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC)

	duration := time.Duration(rand.Int63n(end.Unix()-start.Unix())) * time.Second

	randomTime := uint64(start.Add(duration).Unix())

	return &randomTime
}

func GetRandomBytes(length int) []byte {
	byteArray := make([]byte, length)
	_, err := rand.Read(byteArray)

	if err != nil {
		panic(err)
	}

	return byteArray
}
