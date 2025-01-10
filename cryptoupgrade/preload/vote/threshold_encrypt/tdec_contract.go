package thresholdencrypt

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote"
)

type VerifyShare struct{}

func (v *VerifyShare) RequiredGas() uint64     { return 100 }
func (v *VerifyShare) TargetFunc() interface{} { return interface{}(VerifyDecryptionShare) }
func (v *VerifyShare) GetTypeList() ([]string, []string) {
	itype := []string{"uint64", "uint64", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes", "bytes"}
	otype := []string{"bool"}
	return itype, otype
}

func VerifyDecryptionShare(node_id uint64,
	x_index []uint64,
	cipher1, cipher2, cipher3,
	decryption_share1, decryption_share2, decryption_share3,
	zkp1, zkp2, zkp3,
	commitee_pk_bytes, g_bytes []byte) bool {
	group1 := NewG1()
	g, _ := group1.FromBytes(g_bytes)
	commitee_pk, _ := group1.FromBytes(commitee_pk_bytes)
	cipher, _ := vote.NewSingleElemCipher().Deserialize(cipher1)
	share, _ := group1.FromBytes(decryption_share1)
	share_zkp, _ := NewDecryptionShareZKP().Deserialize(decryption_share1)
	return VerifyDecryptionShareMain(node_id, x_index, cipher, share, share_zkp, commitee_pk, g)
}
