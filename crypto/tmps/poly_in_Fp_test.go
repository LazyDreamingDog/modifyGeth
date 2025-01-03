package tmps

import (
	"math/big"
	"testing"
	"time"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/sirupsen/logrus"
)

func TestProofVerify(t *testing.T) {
	// 设置日志级别
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// 文件数
	N := 2

	// 读取第一个文件并转换为多项式
	poly1 := &Polynomial{}
	poly1.prime = Prime
	poly1, err := poly1.FileToPolynomial("/Users/panzhuochen/offchain-storage/testdata/hello1")
	if err != nil {
		t.Fatalf("Failed to load polynomial from file 1: %v", err)
	}
	d1 := len(poly1.coefficients) - 1
	logrus.Infof("Starting test for polynomial degree (file1): %d", d1)

	totalBytes1 := 0
	for _, coeff := range poly1.GetCoefficients() {
		totalBytes1 += len(coeff.Bytes())
	}
	logrus.WithFields(logrus.Fields{
		"degree":    d1,
		"bf_M_size": totalBytes1,
	}).Info("bf_M generated from file 1")

	// 读取第二个文件并转换为多项式
	poly2 := &Polynomial{}
	poly2.prime = Prime
	poly2, err = poly2.FileToPolynomial("/Users/panzhuochen/offchain-storage/testdata/hello2")
	if err != nil {
		t.Fatalf("Failed to load polynomial from file 2: %v", err)
	}
	d2 := len(poly2.coefficients) - 1
	logrus.Infof("Starting test for polynomial degree (file2): %d", d2)

	totalBytes2 := 0
	for _, coeff := range poly2.GetCoefficients() {
		totalBytes2 += len(coeff.Bytes())
	}
	logrus.WithFields(logrus.Fields{
		"degree":    d2,
		"bf_M_size": totalBytes2,
	}).Info("bf_M generated from file 2")

	// 取最大的 degree 作为公共 degree
	d := d1
	if d2 > d {
		d = d2
	}
	logrus.Infof("Using max degree: %d", d)

	// Setup 阶段 for poly1
	startSetup1 := time.Now()
	com := SetupCommon(d, N)
	// 创建一个新的多项式，将原先的多项式系数复制过去，并且补0
	poly1_padded := NewPolynomial(padCoefficients(poly1.coefficients, d+1), Prime)
	pk1, ek1 := Setup(poly1_padded, com)
	logrus.WithFields(logrus.Fields{
		"degree":     d,
		"setup_time": time.Since(startSetup1),
	}).Info("Setup completed for file 1")

	// Setup 阶段 for poly2
	startSetup2 := time.Now()
	// 创建一个新的多项式，将原先的多项式系数复制过去，并且补0
	poly2_padded := NewPolynomial(padCoefficients(poly2.coefficients, d+1), Prime)
	pk2, ek2 := Setup(poly2_padded, com)
	logrus.WithFields(logrus.Fields{
		"degree":     d,
		"setup_time": time.Since(startSetup2),
	}).Info("Setup completed for file 2")

	// Chal 阶段
	startChal := time.Now()
	chal := GenerateChallenges("Challenge", N)
	challenge1, vk1 := Chal(pk1, chal)
	challenge2, vk2 := Chal(pk2, chal)
	logrus.WithFields(logrus.Fields{
		"degree":    d,
		"chal_time": time.Since(startChal),
	}).Info("Challenge generated")

	helper(pk1)

	// Prove 阶段 for poly1
	startProve1 := time.Now()
	pi1 := Prove(challenge1, ek1)
	logrus.WithFields(logrus.Fields{
		"degree":     d,
		"prove_time": time.Since(startProve1),
	}).Info("Proof generated for file 1")

	// Prove 阶段 for poly2
	startProve2 := time.Now()
	pi2 := Prove(challenge2, ek2)
	logrus.WithFields(logrus.Fields{
		"degree":     d,
		"prove_time": time.Since(startProve2),
	}).Info("Proof generated for file 2")

	// Verify 阶段 for poly1
	startVerify1 := time.Now()
	output, _ := Encode(pk1, challenge1, pi1, vk1)
	pk1, challenge1, pi1, vk1, _ = Decode(output)
	b1 := Verify(pk1, challenge1, pi1, vk1)
	logrus.WithFields(logrus.Fields{
		"degree":      d,
		"verify_time": time.Since(startVerify1),
		"result":      b1,
	}).Info("Verification completed for file 1")

	// Verify 阶段 for poly2
	startVerify2 := time.Now()
	b2 := Verify(pk2, challenge2, pi2, vk2)
	logrus.WithFields(logrus.Fields{
		"degree":      d,
		"verify_time": time.Since(startVerify2),
		"result":      b2,
	}).Info("Verification completed for file 2")

	// 聚合和验证
	logrus.Info("Starting aggregation...")
	startAgg := time.Now()
	pi_a, vk_a := Aggregation([]*Pi{pi1, pi2}, []*Vk{vk1, vk2})
	aggresTime := time.Since(startAgg)

	startAggVerify := time.Now()
	b_a := Verify(pk1, challenge1, pi_a, vk_a)
	aggVerifyTime := time.Since(startAggVerify)

	logrus.WithFields(logrus.Fields{
		"degree":            d,
		"aggregation_time":  time.Since(startAgg),
		"agg_time":          aggresTime,
		"agg_verify_time":   aggVerifyTime,
		"agg_verify_result": b_a,
	}).Info("Aggregation and verification completed")

	logrus.Info("------------------------------------------------")

	// 添加断言来检查验证结果，确保它为 true
	if !b1 {
		t.Errorf("Verification failed for file 1, degree %d", d)
	}
	if !b2 {
		t.Errorf("Verification failed for file 2, degree %d", d)
	}
	if !b_a {
		t.Errorf("Aggregation verification failed for degree %d", d)
	}
}

// padCoefficients 函数用于将系数补齐到指定的长度
func padCoefficients(coeffs []*big.Int, targetLen int) []*big.Int {
	paddedCoeffs := make([]*big.Int, targetLen)
	for i := 0; i < targetLen; i++ {
		if i < len(coeffs) {
			paddedCoeffs[i] = coeffs[i]
		} else {
			paddedCoeffs[i] = big.NewInt(0)
		}
	}
	return paddedCoeffs
}

func helper(pk *Pk) {
	//计算g1的5次方
	g1_5 := OneG1().ScalarMult(pk.GetG1(), big.NewInt(5))
	g2_9 := OneG2().ScalarMult(pk.GetG2(), big.NewInt(9))
	g1_6 := OneG1().ScalarMult(pk.GetG1(), big.NewInt(6))
	g2_8 := OneG2().ScalarMult(pk.GetG2(), big.NewInt(8))
	PrintG1(g1_5, "g1^5")
	PrintG2(g2_9, "g2^9")
	PrintG1(g1_6, "g1^6")
	PrintG2(g2_8, "g2^8")
	p1_37 := bn256.Pair(pk.GetG1(), new(bn256.G2).ScalarMult(pk.GetG2(), big.NewInt(37)))
	p5_9 := bn256.Pair(g1_5, g2_9)
	p6_8 := bn256.Pair(g1_6, g2_8)
	p1_9 := bn256.Pair(OneG1(), g2_9)
	PrintGT(p1_37, "e(g1,g2)^37")
	PrintGT(p5_9, "e(g1^5,g2^9)")
	PrintGT(p6_8, "e(g1^6,g2^8)")
	PrintGT(p1_9, "e(1,g2^9)")
	p1_102 := bn256.Pair(pk.GetG1(), new(bn256.G2).ScalarMult(pk.GetG2(), big.NewInt(102)))
	PrintGT(p1_102, "e(g1,g2)^102")
}
