#include <stdio.h>
#include "ur-wrap.hpp"
#include "ur.h"

char *bytewords_encode(uint8_t *bin, size_t bin_sz)
{
    auto encoded = ur::Bytewords::encode(ur::Bytewords::style::uri, ur::ByteVector(bin, bin + bin_sz));
    int n = encoded.length();

    char *encoded_string = new char[n + 1];
    strcpy(encoded_string, encoded.c_str());

    return encoded_string;
}

uint8_t *bytewords_decode(size_t *bin_sz, char *words)
{
    std::string words_string(words);
    auto buf = ur::Bytewords::decode(ur::Bytewords::style::uri, words_string);
    *bin_sz = buf.size();

    uint8_t *bin = new uint8_t[*bin_sz + 1];
    memcpy(bin, &buf[0], *bin_sz);

    return bin;
}
