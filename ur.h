// ur.h

#ifndef __BYTEWORDS_HPP__
#define __BYTEWORDS_HPP__

#include <stdint.h>

#ifdef __cplusplus
extern "C"
{
#endif
    char *bytewords_encode(uint8_t *bin, size_t bin_sz);
    uint8_t *bytewords_decode(size_t *bin_sz, char *words);
#ifdef __cplusplus
}
#endif

#endif
