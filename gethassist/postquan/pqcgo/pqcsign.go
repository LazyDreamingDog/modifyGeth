package pqcgo

/*
#cgo CFLAGS: -I${SRCDIR}/libs/include
#cgo windows LDFLAGS: ${SRCDIR}/libs/lib/win/libpqmagic.a
#cgo linux LDFLAGS: ${SRCDIR}/libs/lib/linux/libpqmagic.a
#include <stdint.h>
#include <stdlib.h>
#include "./pqcsign_wrapper.h"
*/
import "C"
import (
	"bytes"
	"errors"
	"unsafe"
)

func KeyGen(scheme int) ([]byte, []byte, error) {
	pk := make([]byte, PUBLICKEYBYTES[scheme])
	sk := make([]byte, SECRETKEYBYTES[scheme])
	ret := C.keyGen(C.int(scheme), (*C.uint8_t)(unsafe.Pointer(&pk[0])), (*C.uint8_t)(unsafe.Pointer(&sk[0])))
	if ret != 0 {
		return nil, nil, errors.New("key generation failed")
	}
	return pk, sk, nil
}

func KeyGenWithSeed(scheme int, seed []byte) ([]byte, []byte, error) {
	pk := make([]byte, PUBLICKEYBYTES[scheme])
	sk := make([]byte, SECRETKEYBYTES[scheme])
	ret := C.keyGenWithSeed(C.int(scheme), (*C.uint8_t)(unsafe.Pointer(&pk[0])), (*C.uint8_t)(unsafe.Pointer(&sk[0])), (*C.uint8_t)(unsafe.Pointer(&seed[0])), C.size_t(len(seed)))
	if ret != 0 {
		return nil, nil, errors.New("key generation with seed failed")
	}
	return pk, sk, nil
}

func Sign(scheme int, message []byte, sk []byte) ([]byte, error) {
	sig := make([]byte, SIGNATUREBYTES[scheme]+10)
	siglen := C.size_t(0)
	ret := C.sign(C.int(scheme), (*C.uint8_t)(unsafe.Pointer(&sig[0])), &siglen, (*C.uint8_t)(unsafe.Pointer(&message[0])), C.size_t(len(message)), (*C.uint8_t)(unsafe.Pointer(&sk[0])))
	if ret != 0 {
		return nil, errors.New("signing failed")
	}
	return sig[:siglen], nil
}

func Verify(scheme int, sig []byte, message []byte, pk []byte) (bool, error) {
	ret := C.verify(C.int(scheme), (*C.uint8_t)(unsafe.Pointer(&sig[0])), C.size_t(len(sig)), (*C.uint8_t)(unsafe.Pointer(&message[0])), C.size_t(len(message)), (*C.uint8_t)(unsafe.Pointer(&pk[0])))
	if ret != 0 {
		return false, errors.New("verification failed")
	}
	return true, nil
}

func VerifyKeyGen(scheme int, fsk []byte, bsk []byte, bpk []byte) (bool, error) {
	ret := C.VerifyKeyGen(C.int(scheme), (*C.uint8_t)(unsafe.Pointer(&fsk[0])), (*C.uint8_t)(unsafe.Pointer(&bsk[0])), (*C.uint8_t)(unsafe.Pointer(&bpk[0])))
	if ret != 0 {
		return false, errors.New("key generation verification failed")
	}
	return true, nil
}

func VerifyPubKeyIndex(sheme int, prepk []byte, curpk []byte) (bool, error) {
	pk, _, err := KeyGenWithSeed(sheme, prepk)
	return bytes.Equal(pk, curpk), err
}
