package pqcgo

var PQCSignType = map[string]int{
	"aigis_sig": 0,
	"dilithium": 1,
	"ml_dsa":    2,
	"slh_dsa":   3,
}

const (
	AIGIS_SIG3_PUBLICKEYBYTES         = 1568
	DILITHIUM3_PUBLICKEYBYTES         = 1952
	ML_DSA_65_PUBLICKEYBYTES          = 1952
	SLH_DSA_SHAKE_192f_PUBLICKEYBYTES = 48

	AIGIS_SIG3_SECRETKEYBYTES         = 3888
	DILITHIUM3_SECRETKEYBYTES         = 4000
	ML_DSA_65_SECRETKEYBYTES          = 4032
	SLH_DSA_SHAKE_192f_SECRETKEYBYTES = 96

	AIGIS_SIG3_SIGBYTES         = 3046
	DILITHIUM3_SIGBYTES         = 3293
	ML_DSA_65_SIGBYTES          = 3309
	SLH_DSA_SHAKE_192f_SIGBYTES = 35664
)

var PUBLICKEYBYTES = []int{
	AIGIS_SIG3_PUBLICKEYBYTES,
	DILITHIUM3_PUBLICKEYBYTES,
	ML_DSA_65_PUBLICKEYBYTES,
	SLH_DSA_SHAKE_192f_PUBLICKEYBYTES,
}

var SECRETKEYBYTES = []int{
	AIGIS_SIG3_SECRETKEYBYTES,
	DILITHIUM3_SECRETKEYBYTES,
	ML_DSA_65_SECRETKEYBYTES,
	SLH_DSA_SHAKE_192f_SECRETKEYBYTES,
}

var SIGNATUREBYTES = []int{
	AIGIS_SIG3_SIGBYTES,
	DILITHIUM3_SIGBYTES,
	ML_DSA_65_SIGBYTES,
	SLH_DSA_SHAKE_192f_SIGBYTES,
}
