// Copyright © 2018-2019 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package saplinglib

/*
#include "saplinglib.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
*/
import "C"

import (
	"encoding/json"
	"unsafe"
)

// IguanaSaplingAddress data type to parse saping address output from saplinglib
type IguanaSaplingAddress []struct {
	Num        int    `json:"num"`
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	Seed       struct {
		HDSeed string `json:"HDSeed"`
		Path   string `json:"path"`
	} `json:"seed"`
}

// GetZAddress generates a shielded sapling address using a seed phrase
func GetZAddress(nohd bool, zcount uint, iguanaSeed string, isIguanaSeed bool, coinType uint) IguanaSaplingAddress {

	var zaddrs IguanaSaplingAddress

	// rust_generate_wallet function takes four parameters
	// 1) nohd:				set it to false, if you don't want HD wallet
	// 2) zcount:			the number of sapling addresses you want to generate
	// 3) seed:				the user specified passphrase, which gives the same address everytime if given the same passphrase
	// 4) isIguanaSeed:		set this to true if you want the output to always give a deterministic address based on user specified seed phrase
	// 5) coinType:			cointype is picked from src/chainparam.cpp file of the cryptocurrency. Example, zcash uses 133 for mainnet, 1 for testnet, and komodo uses 141 for mainnet
	_nohd := C.bool(nohd)
	_zcount := C.uint(zcount)
	_seed := C.CString(iguanaSeed)
	_isIguanaSeed := C.bool(isIguanaSeed)
	_coinType := C.uint(141)

	fromRust := C.CString("")
	defer C.free(unsafe.Pointer(fromRust))
	fromRust = C.rust_generate_wallet(_nohd, _zcount, _seed, _isIguanaSeed, _coinType)
	// fmt.Println(C.GoString(fromRust))

	zaddrBytes := []byte(C.GoString(fromRust))

	if err := json.Unmarshal(zaddrBytes, &zaddrs); err != nil {
		panic(err)
	}
	// fmt.Println(zaddrs)
	return zaddrs
}
