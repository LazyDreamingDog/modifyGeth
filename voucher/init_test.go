package voucher

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestABI(t *testing.T) {
	fmt.Printf("BalanceOf: %v\n", BalanceOf)

	fmt.Printf("VoucherAddress: %v\n", common.MutiVoucherAddress)

	tokenName := "BitCoin"
	a := make([]byte, hex.EncodedLen(len(tokenName)))
	hex.Encode(a, []byte(tokenName))
	fmt.Printf("a:string %s,data %v\n", a, a)

	hex := common.Bytes2Hex([]byte(tokenName))
	fmt.Printf("hex: %v\n", hex)

	var fixedArray [20]byte
	copy(fixedArray[20-len([]byte(tokenName)):], []byte(tokenName))
	fmt.Printf("fixedArray:string %s,data %v\n", fixedArray, fixedArray)

	fmt.Printf("string(fixedArray[:]): %v\n", string(fixedArray[:]))
}

func encodingString(s string) [20]byte {
	var result [20]byte
	copy(result[:], []byte(s))
	return result
}

// truncate a string
func decodeString(array [20]byte) string {
	i := 0
	for ; i < 20; i++ {
		if array[i] == 0 {
			break
		}
	}
	return string(array[:i])
}

func TestToken(t *testing.T) {
	tokenName := "BitCoin"
	array := encodingString(tokenName)
	fmt.Printf("array: s:%s v:%v hex:%v\n", array, array, common.Bytes2Hex(array[:]))

	decode := decodeString(array)
	fmt.Printf("decode: %v\n", decode)

	for _, c := range tokenName {
		fmt.Printf("c: %v\n", c)
	}

}
