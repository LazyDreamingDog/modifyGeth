#ifndef PQC_WRAPPER_H
#define PQC_WRAPPER_H

#include <stdint.h>
#include <stddef.h>
#include "./libs/include/pqmagic_api.h"

typedef enum {
    AIGIS_SIG,
    DILITHIUM,
    ML_DSA,
    SLH_DSA
} SignAlgType;

int keyGen(int  scheme, uint8_t *pk, uint8_t *sk);
int keyGenWithSeed(int  scheme, uint8_t *pk, uint8_t *sk, const uint8_t *presk,size_t prepklen);
int sign(int  scheme, uint8_t *sig, size_t *siglen, const uint8_t *m, size_t mlen, const uint8_t *sk);
int verify(int  scheme, const uint8_t *sig, size_t siglen, const uint8_t *m, size_t mlen, const uint8_t *pk);
int VerifyKeyGen(int  scheme, const uint8_t *fsk, const uint8_t *bsk, const uint8_t *bpk);

#endif // PQC_WRAPPER_H