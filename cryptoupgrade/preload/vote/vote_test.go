package vote

import (
	"testing"

	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"

	ec_elgamal "github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote/elgmal"
)

func Test(t *testing.T) {
	group1 := bls12381.NewG1()
	g := group1.One()
	// modular/(2^33)
	_, pk := ec_elgamal.Keygen(g)
	ballot, _ := GenerateBallot(3, 1)
	cipher := EncryptBallot(ballot, g, pk)
	zkp := GenerateBallotZKP(ballot, g, pk, pk)
	t.Logf("zkp size: %d\n", zkp.Size())
	t.Logf("ballot cipher size: %d\n", cipher.Size())

	zkp_bytes := zkp.Serialize()
	zkp_temp, _ := NewBallotZKP().Deserialize(zkp_bytes)
	t.Log(zkp.Equal(zkp_temp))

	v1 := cipher[0].Serialize()
	v2 := cipher[1].Serialize()
	v3 := cipher[2].Serialize()

	verResult := VerifyBallotZKP(v1, v2, v3, zkp.Serialize(), group1.ToBytes(g), group1.ToBytes(pk))
	t.Log(verResult)

	verResult = VerifyBallotZKPMain(cipher, zkp, g, pk, pk)
	t.Log(verResult)
	// res := g
	// target := group1.MulScalar(group1.New(), g, bls12381.FrFromInt(100000000))
	// start := time.Now()
	// for {
	// 	res = group1.Add(group1.New(), res, g)
	// 	if group1.Equal(res, target) {
	// 		break
	// 	}
	// }
	// elapsed := time.Since(start)
	// fmt.Printf("Operation took %s\n", elapsed)
}
