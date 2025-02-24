package pqcgo

import (
	"strconv"
	"testing"
)

func BenchmarkKeyGen(b *testing.B) {
	for scheme := 0; scheme < 4; scheme++ {
		b.Run(strconv.Itoa(scheme), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = KeyGen(scheme)
			}
		})
	}
}

func BenchmarkKeyGenWithSeed(b *testing.B) {
	seed := []byte("thisisaseed")
	for scheme := 0; scheme < 4; scheme++ {
		b.Run(strconv.Itoa(scheme), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = KeyGenWithSeed(scheme, seed)
			}
		})
	}
}

func BenchmarkSign(b *testing.B) {
	message := []byte("test message")
	for scheme := 0; scheme < 4; scheme++ {
		b.Run(strconv.Itoa(scheme), func(b *testing.B) {
			_, sk, _ := KeyGen(scheme)
			for i := 0; i < b.N; i++ {
				_, _ = Sign(scheme, message, sk)
			}
		})
	}
}

func BenchmarkVerify(b *testing.B) {
	message := []byte("test message")
	for scheme := 0; scheme < 4; scheme++ {
		b.Run(strconv.Itoa(scheme), func(b *testing.B) {
			pk, sk, _ := KeyGen(scheme)
			sig, _ := Sign(scheme, message, sk)
			for i := 0; i < b.N; i++ {
				_, _ = Verify(scheme, sig, message, pk)
			}
		})
	}
}

func BenchmarkVerifyKeyGen(b *testing.B) {
	for scheme := 0; scheme < 4; scheme++ {
		b.Run(strconv.Itoa(scheme), func(b *testing.B) {
			_, fsk, _ := KeyGen(scheme)
			bpk, bsk, _ := KeyGenWithSeed(scheme, fsk)
			for i := 0; i < b.N; i++ {
				_, _ = VerifyKeyGen(scheme, fsk, bsk, bpk)
			}
		})
	}
}
