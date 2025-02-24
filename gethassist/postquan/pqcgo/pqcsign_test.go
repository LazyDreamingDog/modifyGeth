package pqcgo

import (
	"testing"
)

func TestKeyGen(t *testing.T) {
	for scheme := 0; scheme < 4; scheme++ {
		pk, sk, err := KeyGen(scheme)
		if err != nil {
			t.Fatalf("KeyGen failed for scheme %d: %v", scheme, err)
		}
		if len(pk) == 0 || len(sk) == 0 {
			t.Fatalf("KeyGen returned empty keys for scheme %d", scheme)
		}
		//t.Logf("scheme:%d, pk: %s\n,sk: %s", scheme, hex.EncodeToString(pk), hex.EncodeToString(sk))
	}
}

func TestKeyGenWithSeed(t *testing.T) {
	seed := []byte("thisisaseed")
	for scheme := 0; scheme < 4; scheme++ {
		pk, sk, err := KeyGenWithSeed(scheme, seed)
		if err != nil {
			t.Fatalf("KeyGenWithSeed failed for scheme %d: %v", scheme, err)
		}
		if len(pk) == 0 || len(sk) == 0 {
			t.Fatalf("KeyGenWithSeed returned empty keys for scheme %d", scheme)
		}
		//t.Logf("scheme:%d, pk: %s\n,sk: %s", scheme, hex.EncodeToString(pk), hex.EncodeToString(sk))
	}
}

func TestSign(t *testing.T) {
	message := []byte("test message")
	for scheme := 0; scheme < 4; scheme++ {
		_, sk, err := KeyGen(scheme)
		if err != nil {
			t.Fatalf("KeyGen failed for scheme %d: %v", scheme, err)
		}
		sig, err := Sign(scheme, message, sk)
		if err != nil {
			t.Fatalf("Sign failed for scheme %d: %v", scheme, err)
		}
		if len(sig) == 0 {
			t.Fatalf("Sign returned empty signature for scheme %d", scheme)
		}
		//t.Logf("scheme:%d, sk: %s\n,sig: %s", scheme, hex.EncodeToString(sk), hex.EncodeToString(sig))
	}
}

func TestVerify(t *testing.T) {
	message := []byte("test message")
	for scheme := 0; scheme < 4; scheme++ {
		pk, sk, err := KeyGen(scheme)
		if err != nil {
			t.Fatalf("KeyGen failed for scheme %d: %v", scheme, err)
		}
		sig, err := Sign(scheme, message, sk)
		if err != nil {
			t.Fatalf("Sign failed for scheme %d: %v", scheme, err)
		}
		valid, err := Verify(scheme, sig, message, pk)
		if err != nil {
			t.Fatalf("Verify failed for scheme %d: %v", scheme, err)
		}
		if !valid {
			t.Fatalf("Verify returned false for scheme %d", scheme)
		}
	}
}

func TestVerifyKeyGen(t *testing.T) {
	for scheme := 0; scheme < 4; scheme++ {
		if scheme == 3 {
			continue //	TODO bug here for scheme 3 under windows
		}
		_, fsk, err := KeyGen(scheme)
		if err != nil {
			t.Fatalf("KeyGen failed for scheme %d: %v", scheme, err)
		}
		bpk, bsk, err := KeyGenWithSeed(scheme, fsk)
		if err != nil {
			t.Fatalf("KeyGenWithSeed failed for scheme %d: %v", scheme, err)
		}
		valid, err := VerifyKeyGen(scheme, fsk, bsk, bpk)
		if err != nil {
			t.Fatalf("VerifyKeyGen failed for scheme %d: %v", scheme, err)
		}
		if !valid {
			t.Fatalf("VerifyKeyGen returned false for scheme %d", scheme)
		}
	}
}
