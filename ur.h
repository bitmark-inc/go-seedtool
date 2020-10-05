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
    char *ur_encode_seed(uint8_t *bin, size_t bin_sz, size_t max_fragment_length);
    uint8_t *ur_decode_seed(size_t *bin_sz, size_t part_sz, char **parts);
#ifdef __cplusplus
}
#endif

#endif
