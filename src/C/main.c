#include <stdio.h>
#include <stdbool.h>

#include "saplinglib.h"

int main() {
	// rust_generate_wallet function takes four parameters
	// 1) nohd:				set it to false, if you don't want HD wallet
	// 2) zcount:			the number of sapling addresses you want to generate
	// 3) seed:				the user specified passphrase, which gives the same address everytime if given the same passphrase
	// 4) is_iguana_seed:	set this to true if you want the output to always give a deterministic address based on user specified seed phrase
	// 5) coinType:			cointype is picked from src/chainparam.cpp file of the cryptocurrency. Example, zcash uses 133 for mainnet, 1 for testnet, and komodo uses 141 for mainnet
	bool nohd = false;
	int zcount = 1;
	char *seed = "user specified seed phrase";
	bool is_iguana_seed = true;
	int cointype = 141;

	char * from_rust = rust_generate_wallet(nohd, zcount, seed, is_iguana_seed, cointype);
	char *stri = from_rust;
	printf("%s", stri);
	rust_free_string(from_rust);

	return 0;
}

// From root directory of saplinglib execute this command:

// LINUX:
// 
// git clone https://github.com/satindergrewal/saplinglib
// cd saplinglib
// curl https://sh.rustup.rs -sSf | sh
// source $HOME/.cargo/env
// cargo build --target=x86_64-unknown-linux-gnu --release
// gcc ./src/C/main.c -I./src -L./target/x86_64-unknown-linux-gnu/release -lsaplinglib -lpthread -ldl -lm -o main

// CROSS-COMPILING FOR WINDOWS USING MINGW:
//
// git clone https://github.com/satindergrewal/saplinglib
// cd saplinglib
// curl https://sh.rustup.rs -sSf | sh
// source $HOME/.cargo/env
// rustup target add x86_64-pc-windows-gnu
// cargo build --target=x86_64-pc-windows-gnu --release
// x86_64-w64-mingw32-gcc ./src/C/main.c -I./src -L./target/x86_64-pc-windows-gnu/release -lsaplinglib -lws2_32 -luserenv -o main.exe

// macOS x86_64:
//
// git clone https://github.com/satindergrewal/saplinglib
// cd saplinglib
// curl https://sh.rustup.rs -sSf | sh
// source $HOME/.cargo/env
// cargo build --target=x86_64-apple-darwin --release
// gcc ./src/C/main.c -I./src -L./target/x86_64-apple-darwin/release -lsaplinglib -framework Security -o main

// macOS arm64:
//
// git clone https://github.com/satindergrewal/saplinglib
// cd saplinglib
// curl https://sh.rustup.rs -sSf | sh
// source $HOME/.cargo/env
// cargo build --target=aarch64-apple-darwin --release
// gcc ./src/C/main.c -I./src -L./target/aarch64-apple-darwin/release -lsaplinglib -framework Security -o main