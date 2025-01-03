package tmps

import (
	"encoding/json"
	"fmt"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

type InputData struct {
	Pk *Pk      `json:"pk"`
	C  *big.Int `json:"c"`
	Pi *Pi      `json:"pi"`
	Vk *Vk      `json:"Vk"`
}

// 定义 bn256.G1 的类型别名
type G1Alias bn256.G1

// 实现 JSON 序列化接口
func (g *G1Alias) MarshalJSON() ([]byte, error) {
	if g == nil {
		return []byte("null"), nil
	}

	// 使用 Marshal 将 G1 转为二进制数据
	data := (*bn256.G1)(g).Marshal()
	return json.Marshal(data)
}

// 实现 JSON 反序列化接口
func (g *G1Alias) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	// 从 JSON 中解析出二进制数据
	var binData []byte
	if err := json.Unmarshal(data, &binData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	point := new(bn256.G1)
	if _, err := point.Unmarshal(binData); err != nil {
		return fmt.Errorf("failed to unmarshal G1 point: %w", err)
	}

	*g = G1Alias(*point)
	return nil
}

// 定义 bn256.G2 的类型别名
type G2Alias bn256.G2

// 实现 JSON 序列化接口
func (g *G2Alias) MarshalJSON() ([]byte, error) {
	if g == nil {
		return []byte("null"), nil
	}

	// 使用 Marshal 将 G2 转为二进制数据
	data := (*bn256.G2)(g).Marshal()
	return json.Marshal(data)
}

// 实现 JSON 反序列化接口
func (g *G2Alias) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	// 从 JSON 中解析出二进制数据
	var binData []byte
	if err := json.Unmarshal(data, &binData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	point := new(bn256.G2)
	if _, err := point.Unmarshal(binData); err != nil {
		return fmt.Errorf("failed to unmarshal G2 point: %w", err)
	}

	*g = G2Alias(*point)
	return nil
}

// 定义 bn256.GT 的类型别名
type GTAlias bn256.GT

// 实现 JSON 序列化接口
func (g *GTAlias) MarshalJSON() ([]byte, error) {
	if g == nil {
		return []byte("null"), nil
	}

	// 使用 Marshal 将 GT 转为二进制数据
	data := (*bn256.GT)(g).Marshal()
	return json.Marshal(data)
}

// 实现 JSON 反序列化接口
func (g *GTAlias) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	// 从 JSON 中解析出二进制数据
	var binData []byte
	if err := json.Unmarshal(data, &binData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	point := new(bn256.GT)
	if _, err := point.Unmarshal(binData); err != nil {
		return fmt.Errorf("failed to unmarshal GT point: %w", err)
	}
	*g = GTAlias(*point)
	return nil
}

// BigIntAlias 为 big.Int 提供 JSON 支持
type BigIntAlias big.Int

// 实现 JSON 序列化接口
func (b *BigIntAlias) MarshalJSON() ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}
	return json.Marshal((*big.Int)(b).String())
}

// 实现 JSON 反序列化接口
func (b *BigIntAlias) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	i := new(big.Int)
	if _, ok := i.SetString(s, 10); !ok {
		return fmt.Errorf("invalid big.Int format: %s", s)
	}

	*b = BigIntAlias(*i)
	return nil
}

// encode 将函数 Verify 的入参打包为字节数组
func Encode(pk *Pk, c *big.Int, pi *Pi, vk *Vk) ([]byte, error) {
	inputData := InputData{
		Pk: pk,
		C:  c,
		Pi: pi,
		Vk: vk,
	}

	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// decode 将字节数组还原为 Verify 函数的入参
func Decode(input []byte) (*Pk, *big.Int, *Pi, *Vk, error) {
	var inputData InputData
	err := json.Unmarshal(input, &inputData)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	return inputData.Pk, inputData.C, inputData.Pi, inputData.Vk, nil
}
