#ifndef VERIFY_H
#define VERIFY_H

#include <stddef.h>
#include <stdint.h>
#include "params.h"

#define verify AIGIS_ENC_NAMESPACE(verify)
int verify(const uint8_t *a, const uint8_t *b, size_t len);

#define cmov AIGIS_ENC_NAMESPACE(cmov)
void cmov(uint8_t *r, const uint8_t *x, size_t len, uint8_t b);

#endif
