package vote

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
)

// func flat_3bytes32(byte_matrix [3][32]byte) []byte {
// 	var flat bytes.Buffer
// 	for i := 0; i < len(byte_matrix); i++ {
// 		flat.Write(byte_matrix[i][:])
// 	}
// 	return flat.Bytes()
// }

// func flat_6bytes32(byte_matrix [6][32]byte) []byte {
// 	var flat bytes.Buffer
// 	for i := 0; i < len(byte_matrix); i++ {
// 		flat.Write(byte_matrix[i][:])
// 	}
// 	return flat.Bytes()
// }

// func bytes_to_cipher(cipher_bytes [6][32]byte) (*SingleElemCipher, error) {
// 	group1 := NewG1()
// 	var point_g1_bytes bytes.Buffer
// 	for i := 0; i < 3; i++ {
// 		point_g1_bytes.Write(cipher_bytes[i][:])
// 	}
// 	c1_point, err := group1.FromBytes(point_g1_bytes.Bytes())
// 	if err != nil {
// 		return nil, err
// 	}

// 	point_g1_bytes.Reset()
// 	for i := 3; i < 6; i++ {
// 		point_g1_bytes.Write(cipher_bytes[i][:])
// 	}
// 	c2_point, err := group1.FromBytes(point_g1_bytes.Bytes())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &SingleElemCipher{C1: c1_point, C2: c2_point}, nil
// }

type VerifyBallot struct{}

func (v *VerifyBallot) RequiredGas() uint64     { return 100 }
func (v *VerifyBallot) TargetFunc() interface{} { return interface{}(VerifyBallotZKP) }
func (v *VerifyBallot) GetTypeList() ([]string, []string) {
	itype := []string{"bytes", "bytes", "bytes", "bytes", "bytes", "bytes"}
	otype := []string{"bool"}
	return itype, otype
}

func VerifyBallotZKP(v1, v2, v3, zkp_bytes, g_bytes, pk_bytes []byte) bool {
	group1 := NewG1()

	g, _ := group1.FromBytes(g_bytes)
	pk, _ := group1.FromBytes(pk_bytes)

	var ballot_cipher BallotCipher

	cipher, _ := NewSingleElemCipher().Deserialize(v1)
	ballot_cipher = append(ballot_cipher, cipher)
	cipher, _ = NewSingleElemCipher().Deserialize(v2)
	ballot_cipher = append(ballot_cipher, cipher)
	cipher, _ = NewSingleElemCipher().Deserialize(v3)
	ballot_cipher = append(ballot_cipher, cipher)
	cipher = NewSingleElemCipher()
	cipher.C1, cipher.C2 = group1.MulScalar(group1.New(), g, NewFr().Zero()), group1.MulScalar(group1.New(), g, NewFr().Zero())
	ballot_cipher = append(ballot_cipher, cipher)

	ballot_zkp, _ := NewBallotZKP().Deserialize(zkp_bytes)

	return VerifyBallotZKPMain(ballot_cipher, ballot_zkp, g, pk, pk)

}
