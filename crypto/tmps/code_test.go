package tmps

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"log"
	"testing"
)

type Student struct {
	Name *string
	Age  *int
}

func TestCode(t *testing.T) {
	name := "pzc"
	age := 18
	st := Student{
		Name: &name,
		Age:  &age,
	}
	output, _ := json.Marshal(st)
	t.Log(string(output))
	_ = json.Unmarshal(output, st)
	t.Log(st)
}

func TestJson(t *testing.T) {
	original := new(bn256.G1)
	_, original, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate random G1 point: %v", err)
	}

	// 将 G1 转换为 G1Alias
	originalAlias := (*G1Alias)(original)

	// 序列化为 JSON
	jsonData, err := json.Marshal(originalAlias)
	if err != nil {
		log.Fatalf("Failed to marshal G1Alias: %v", err)
	}
	fmt.Printf("Serialized JSON: %s\n", jsonData)

	// 从 JSON 反序列化回 G1Alias
	var deserialized G1Alias
	err = json.Unmarshal(jsonData, &deserialized)
	if err != nil {
		log.Fatalf("Failed to unmarshal G1Alias: %v", err)
	}

	// 验证反序列化后的值是否与原始值一致
	if original.String() != ((*bn256.G1)(&deserialized)).String() {
		log.Fatalf("Mismatch after JSON unmarshal: expected %s, got %s",
			original.String(), ((*bn256.G1)(&deserialized)).String())
	}

	fmt.Println("JSON serialization and deserialization successful, values match!")

}
