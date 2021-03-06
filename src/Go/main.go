package main

/*
#include "saplinglib.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	// rust_generate_wallet function takes four parameters
	// 1) nohd:				set it to false, if you don't want HD wallet
	// 2) zcount:			the number of sapling addresses you want to generate
	// 3) seed:				the user specified passphrase, which gives the same address everytime if given the same passphrase
	// 4) isIguanaSeed:		set this to true if you want the output to always give a deterministic address based on user specified seed phrase
	// 5) coinType:			cointype is picked from src/chainparam.cpp file of the cryptocurrency. Example, zcash uses 133 for mainnet, 1 for testnet, and komodo uses 141 for mainnet
	nohd := C.bool(false)
	zcount := C.uint(1)
	seed := C.CString("user specified seed phrase")
	isIguanaSeed := C.bool(true)
	coinType := C.uint(141)

	fromRust := C.CString("")
	defer C.free(unsafe.Pointer(fromRust))
	fromRust = C.rust_generate_wallet(nohd, zcount, seed, isIguanaSeed, coinType)
	fmt.Println(C.GoString(fromRust))
}

// Compile using this command:

// macOS x86_64:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build

// macOS arm64:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build

// Windows:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build

// Linux:
// env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build
