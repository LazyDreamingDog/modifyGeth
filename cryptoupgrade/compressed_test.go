package cryptoupgrade

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// * The test will create a copy in current folder.
func TestCompress(t *testing.T) {
	compressFile := "./compressed.go"
	compressedString, err := compressFileToString(compressFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// rootpath is the path of the file to be executed
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("get current workfolder err:%v\n", err)
		return
	} else {
		fmt.Printf("currentPath: %v\n", rootPath)
	}
	outputPath := rootPath + "/decompressed" + "_" + "main.go"
	err = decompressStringToFile(compressedString, outputPath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// * Decompressed check
func TestDecompress(t *testing.T) {
	decodedData := "Hello world!"
	rootPath, err := os.Getwd()
	outputPath := rootPath + "/decompressed" + "_" + "main.go"
	err = os.WriteFile(outputPath, []byte(decodedData), 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// * Use this test to get compressed string, as input for geth.js command line
func TestCompressOnly(t *testing.T) {
	compressFile := "./testgo/add.go"
	compressedString, err := compressFileToString(compressFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("compressedString: %s\n", compressedString)
}

func convertStringToBytes10(input string) [10]byte {
	var result [10]byte
	inputBytes := []byte(input)
	copy(result[:], inputBytes)
	return result
}

func TestBytes(t *testing.T) {
	itype := "int"
	// Convert string to [10]byte
	itypeBytes := convertStringToBytes10(itype)
	// Convert [10]byte to Hex. Solidiy
	// When passing parameters to the solidity compiler, the byte needs to be converted to hexadecimal
	result2 := "0x" + common.Bytes2Hex(itypeBytes[:])
	fmt.Printf("result2: %v\n", result2)
}
