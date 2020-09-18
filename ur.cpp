#include <stdio.h>
#include "ur-wrap.hpp"
#include "ur.h"

char *bytewords_encode(void *bytes, size_t len)
{
    uint8_t *start = (uint8_t *)bytes;
    auto encoded = ur::Bytewords::encode(ur::Bytewords::style::uri, ur::ByteVector(start, start + len));
    int n = encoded.length();

    char *encoded_string = new char[n + 1];
    strcpy(encoded_string, encoded.c_str());

    return encoded_string;
}

void *bytewords_decode(size_t *bin_sz, char *encoded_hex)
{
    std::string encoded_hex_string(encoded_hex);
    auto buf = ur::Bytewords::decode(ur::Bytewords::style::uri, encoded_hex_string);
    *bin_sz = buf.size();

    uint8_t *bin = (uint8_t *)malloc(*bin_sz + 1);
    memcpy(bin, &buf[0], *bin_sz);

    return bin;
}
