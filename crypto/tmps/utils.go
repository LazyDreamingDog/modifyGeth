package tmps

import (
	"fmt"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/ethereum/go-ethereum/log"
)

// 将bn256.G2转换为string
func PrintG2(g2 *bn256.G2, str string) {
	// 序列化为字节数组
	data := g2.Marshal()

	// 打印字节数组
	formattedData := fmt.Sprintf("%x", data)
	log.Debug(str)
	log.Debug("G2 value in bytes: \n", formattedData)

}

func PrintG1(g1 *bn256.G1, str string) {
	// 序列化为字节数组
	data := g1.Marshal()

	// 打印字节数组
	formattedData := fmt.Sprintf("%x", data)
	log.Debug(str)
	log.Debug("G1 value in bytes: \n", formattedData)
}

func PrintGT(gt *bn256.GT, str string) {
	// 序列化为字节数组
	data := gt.Marshal()

	// 打印字节数组
	formattedData := fmt.Sprintf("%x", data)
	log.Debug(str)
	log.Debug("GT value in bytes: \n", formattedData)
}
