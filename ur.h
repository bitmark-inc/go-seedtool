// ur.h

#ifndef __BYTEWORDS_HPP__
#define __BYTEWORDS_HPP__

#include <stdint.h>

#ifdef __cplusplus
extern "C"
{
#endif

    typedef void *UR;
    char *bytewords_encode(void *bytes, size_t len);
    void *bytewords_decode(size_t *bin_sz, char *encoded_hex);

#ifdef __cplusplus
}
#endif

#endif
