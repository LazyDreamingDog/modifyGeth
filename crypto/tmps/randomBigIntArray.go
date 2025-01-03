package tmps

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// Prime is a prime over which we form a basic field: 36u⁴+36u³+24u²+6u+1.
var Prime = bigFromBase10("21888242871839275222246405745257275088548364400416034343698204186575808495617")

func GenerateRandomBigIntArray(max *big.Int, d int) []*big.Int {
	result := make([]*big.Int, d+1)

	for i := 0; i <= d; i++ {

		randomInt, _ := rand.Int(rand.Reader, max)
		result[i] = randomInt
	}

	return result
}

func GenerateChallenges(str string, N int) []*big.Int {
	challenges := make([]*big.Int, N)
	for i := 0; i < N; i++ {
		now := time.Now().UnixNano()
		combinedStr := str + fmt.Sprintf("%d", now) + fmt.Sprintf("%d", i)
		randomBytes := make([]byte, len(combinedStr))
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}
		randomInt := new(big.Int).SetBytes(randomBytes)
		challenges[i] = randomInt
	}
	return challenges
}

func bigFromBase10(s string) *big.Int {
	n, _ := new(big.Int).SetString(s, 10)
	return n
}
