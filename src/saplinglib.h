#include <stdbool.h>

char * rust_generate_wallet(bool nohd, unsigned int zcount, const char* entropy, bool iguana_seed, unsigned int cointype);
void   rust_free_string(char* s);