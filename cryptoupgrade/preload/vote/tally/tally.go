package tally

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote"
)

func AggregateBallotMain(res vote.BallotCipher, encryption_ballot vote.BallotCipher) {
	group1 := NewG1()
	// if len(res) != len(encryption_ballot) {
	// 	return nil, errors.New("wrong ballot syntex")
	// }
	for i := 0; i < len(res); i++ {
		res[i].C1 = group1.Add(group1.New(), res[i].C1, encryption_ballot[i].C1)
		res[i].C2 = group1.Add(group1.New(), res[i].C2, encryption_ballot[i].C2)
	}
	// return res, nil
}
