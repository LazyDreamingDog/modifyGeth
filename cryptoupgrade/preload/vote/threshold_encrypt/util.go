package thresholdencrypt

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
)

func lagrange_id_at_zero(id uint64, x_index []uint64) *Fr {
	len := len(x_index)
	res := NewFr().One()
	id_fr := FrFromInt(int(id))
	for i := 0; i < len; i++ {
		x_index_fr := FrFromInt(int(x_index[i]))
		res = NewFr().Mul(res, x_index_fr)
		res = NewFr().Mul(res, NewFr().Inverse(NewFr().Sub(x_index_fr, id_fr)))
	}
	return res
}
