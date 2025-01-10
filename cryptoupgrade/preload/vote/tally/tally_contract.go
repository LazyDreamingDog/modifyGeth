package tally

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote"
)

type AggrBallot struct{}

func (v *AggrBallot) TargetFunc() interface{} { return interface{}(AggregateBallot) }
func (v *AggrBallot) RequiredGas() uint64     { return 100 }
func (v *AggrBallot) GetTypeList() ([]string, []string) {
	itype := []string{"bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes"}
	otype := []string{"bytes", "bytes", "bytes"}
	return itype, otype
}

func AggregateBallot(v1res, v2res, v3res, v1, v2, v3, g_bytes []byte) ([]byte, []byte, []byte) {
	group1 := NewG1()
	g, _ := group1.FromBytes(g_bytes)
	var ballot_cipher vote.BallotCipher
	cipher, _ := vote.NewSingleElemCipher().Deserialize(v1)
	ballot_cipher = append(ballot_cipher, cipher)
	cipher, _ = vote.NewSingleElemCipher().Deserialize(v2)
	ballot_cipher = append(ballot_cipher, cipher)
	cipher, _ = vote.NewSingleElemCipher().Deserialize(v3)
	ballot_cipher = append(ballot_cipher, cipher)
	cipher = vote.NewSingleElemCipher()
	cipher.C1, cipher.C2 = group1.MulScalar(group1.New(), g, NewFr().Zero()), group1.MulScalar(group1.New(), g, NewFr().Zero())
	ballot_cipher = append(ballot_cipher, cipher)

	var ballot_cipher_res vote.BallotCipher
	cipher, _ = vote.NewSingleElemCipher().Deserialize(v1res)
	ballot_cipher_res = append(ballot_cipher_res, cipher)
	cipher, _ = vote.NewSingleElemCipher().Deserialize(v2res)
	ballot_cipher_res = append(ballot_cipher_res, cipher)
	cipher, _ = vote.NewSingleElemCipher().Deserialize(v3res)
	ballot_cipher_res = append(ballot_cipher_res, cipher)
	cipher = vote.NewSingleElemCipher()
	cipher.C1, cipher.C2 = group1.MulScalar(group1.New(), g, NewFr().Zero()), group1.MulScalar(group1.New(), g, NewFr().Zero())
	ballot_cipher_res = append(ballot_cipher_res, cipher)

	ballot_cipher_res, err := ballot_cipher_res.Mul(ballot_cipher)
	if err != nil {
		return nil, nil, nil
	}
	return ballot_cipher_res[0].Serialize(), ballot_cipher_res[1].Serialize(), ballot_cipher_res[2].Serialize()
}

type AggrShare struct{}

func (v *AggrShare) TargetFunc() interface{} { return interface{}(AggregateShare) }
func (v *AggrShare) RequiredGas() uint64     { return 100 }
func (v *AggrShare) GetTypeList() ([]string, []string) {
	itype := []string{"bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes"}
	otype := []string{"bytes", "bytes", "bytes"}
	return itype, otype
}

func AggregateShare(share1res_bytes, share2res_bytes, share3res_bytes, share1_bytes, share2_bytes, share3_bytes, g_bytes []byte) ([]byte, []byte, []byte) {
	group1 := NewG1()

	share1res, _ := group1.FromBytes(share1res_bytes)
	share, _ := group1.FromBytes(share1_bytes)
	share1res = group1.Add(group1.New(), share1res, share)

	share2res, _ := group1.FromBytes(share2res_bytes)
	share, _ = group1.FromBytes(share2_bytes)
	share2res = group1.Add(group1.New(), share2res, share)

	share3res, _ := group1.FromBytes(share3res_bytes)
	share, _ = group1.FromBytes(share3_bytes)
	share1res = group1.Add(group1.New(), share3res, share)

	return group1.ToBytes(share1res), group1.ToBytes(share2res), group1.ToBytes(share3res)
}
