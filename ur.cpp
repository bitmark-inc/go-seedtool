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

char *ur_encode_seed(uint8_t *bin, size_t bin_sz, size_t max_fragment_length)
{
    ur::ByteVector bytes = ur::ByteVector(bin, bin + bin_sz);

    auto u = ur::UR("crypto-seed", bytes);
    auto encoder = ur::UREncoder(u, max_fragment_length);
    size_t seq_len;
    if (encoder.is_single_part())
    {
        seq_len = 1;
    }
    else
    {
        seq_len = encoder.seq_len();
    }

    std::string result = "";
    while (!encoder.is_complete())
    {
        result += encoder.next_part();
        result += "\n";
    }

    // initialize an array of char with a size without the last "\n"
    char *encoded_string = new char[result.size()];
    strcpy(encoded_string, result.c_str());

    // replace the last char by the null-terminated string
    encoded_string[result.size() - 1] = '\0';

    return encoded_string;
}

uint8_t *ur_decode_seed(size_t *bin_sz, size_t part_sz, char **parts)
{
    std::vector<ur::UR> ur_shares;

    ur::URDecoder *decoder = new ur::URDecoder();
    for (uint32_t i = 0; i < part_sz; i++)
    {
        // printf("c part [%d] %p\n", i, parts[i]);
        std::string part(parts[i]);
        decoder->receive_part(part);
        if (decoder->is_failure())
        {
            // error handling
        }
    }

    uint8_t *bin;

    if (decoder->is_success())
    {
        ur::UR result = decoder->result_ur();
        ur::ByteVector buf = result.cbor();
        *bin_sz = buf.size();

        bin = new uint8_t[*bin_sz + 1];
        memcpy(bin, &buf[0], *bin_sz);
    }

    return bin;
}
