/* Based on the public domain implementation in
 * crypto_hash/keccakc512/simple/ from http://bench.cr.yp.to/supercop.html
 * by Ronny Van Keer
 * and the public domain "TweetFips202" implementation
 * from https://twitter.com/tweetfips202
 * by Gilles Van Assche, Daniel J. Bernstein, and Peter Schwabe */
// Modified by HaveYouTall

#include <stddef.h>
#include <stdint.h>
#include "hash.h"

#define NROUNDS 24
#define ROL(a, offset) ((a << offset) ^ (a >> (64-offset)))

/*************************************************
* Name:        sm3_256
*
* Description: SM3-256 with non-incremental API
*
* Arguments:   - uint8_t *h:        pointer to output (32 bytes)
*              - const uint8_t *in: pointer to input
*              - size_t inlen:      length of input in bytes
**************************************************/
void sm3_256(uint8_t h[32], const uint8_t *in, size_t inlen)
{
  sm3_extended(h, 32, in, inlen);
}

/*************************************************
* Name:        sm3_512
*
* Description: SM3-512 with non-incremental API
*
* Arguments:   - uint8_t *h:        pointer to output (64 bytes)
*              - const uint8_t *in: pointer to input
*              - size_t inlen:      length of input in bytes
**************************************************/
void sm3_512(uint8_t *h, const uint8_t *in, size_t inlen)
{
  sm3_extended(h, 64, in, inlen);
}
