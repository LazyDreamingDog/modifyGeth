package preload

import (
	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote"
	tally "github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote/tally"
	thresholdencrypt "github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote/threshold_encrypt"
)

type PreLoadAlgorithm interface {
	RequiredGas() uint64
	TargetFunc() interface{}
	GetTypeList() ([]string, []string)
}

// Attention: algorithm must be captial
var PreLoadAlgorithmMap = map[string]PreLoadAlgorithm{
	"VerifyDecryptionShareZKP": &thresholdencrypt.VerifyShare{},
	"AggregateShare":           &tally.AggrShare{},
	"VerifyBallotZKP":          &vote.VerifyBallot{},
	"AggregateBallot":          &tally.AggrBallot{},
}

func IsPreLoad(algoName string) (PreLoadAlgorithm, bool) {
	p, ok := PreLoadAlgorithmMap[algoName]
	return p, ok
}
