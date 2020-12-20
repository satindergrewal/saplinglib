
### Sapling address generation library

This library is based on Zcash's [librustzcash](https://github.com/zcash/librustzcash), and parts of this library are taken from the [zecpaperwallet](https://github.com/adityapk00/zecpaperwallet).

### Compiling library

To generate the linkable static library follow these steps:

#### Step 0:

You **must** [install Rust](https://www.rust-lang.org/tools/install) on your machine. You need to:

```shell
curl https://sh.rustup.rs -sSf | sh
```

#### Step 1:

To compile a static linker library for your machine, just execute the following command:

```shell
make
```

It will generate a `libsaplinglib.a` file.


### Redistributable static release libraries

To compile a static redistributable libraries, just execute the following command:

```shell
make dist
```

It will compile and copy the static libraries for Windows, Linux and MacOS to a new directory named `dist/`

```shell
➜  dist git:(master) ✗ tree
.
├── dist
│   ├── darwin
│   │   └── libsaplinglib.a
│   ├── darwin_arm64
│   │   └── libsaplinglib.a
│   ├── linux
│   │   └── libsaplinglib.a
│   └── win64
│       └── libsaplinglib.a

3 directories, 3 files
```


### C Example Code

You can use following C example to use the library:

Save the following contents in `csapling.c` file:

```C
#include <stdio.h>

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
```

Assuming you saved the `csapling.c` in the same directory where the library file `libsaplinglib.a` is located, compile it using the following command:

```shell
gcc csapling.c -I./src -L./ -lsaplinglib -lpthread -ldl -o csapling -framework Security
```

It will generate a binary named `csapling` which if you execute will give the following output:

```json
➜  saplinglib ./csapling
[
  {
    "num": 0,
    "address": "zs10znh5fxagl4z2efdy2rltgas9aahjscjuj3slsjyk96zfn5s8vu3sz6u8s4jkyl9zswp2ucm68j",
    "private_key": "secret-extended-key-main1q0eaejzrqqqqpqpl5kj9676vn6dx4ul0s8vc2xhqu3g2f22r8494l0sjkega75nvhupyasuxrfyj2usr2g8ru2uv4y8d88g3xtrhg0jvcuzgy50wp3dsdnfs3nxjaj2qvpswg93x0e5sety25d6ktcgzkc7ntxq9rg60mfcq9fh8gp97h8aw8ccvn74z68tps6d43ukww4f55k6rhm2322sc02ugq7tur0e9kpj34tevyeej4h38dfz6ktj4thtv3alg0eydkm4rrlssautyr",
    "seed": {
      "HDSeed": "fe50eb2add6c3e1ecc550f901fa737cbebab7b7a1dbf6827d4b0fd3521d2f93e",
      "path": "m/32'/141'/0'"
    }
  }
]
```

#### Cross-compiling C example code

Windows

```shell
x86_64-w64-mingw32-gcc ./src/C/main.c -I./src -L./target/x86_64-pc-windows-gnu/release -lsaplinglib -lws2_32 -luserenv -o main.exe
```


### Go Language Example Code

You can use following Go language example to use the library:

Save the following contents in `gosapling.go` file:


```golang
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
```

Assuming you saved the `gosapling.go` in the same directory where the library file `libsaplinglib.a` is located, compile it using the following command:

NOTE: You have to specify `CGO_CFLAGS` and `CGO_LDFLAGS` values to include the header `saplinglib.h` and `libsaplinglib.a` file to call functions from static library files.

```shell
env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/ -lsaplinglib -framework Security" go build -o gosapling gosapling.go
```

It will generate a binary named `gosapling` which if you execute will give the following output:


```json
➜  saplinglib ./gosapling
[
  {
    "num": 0,
    "address": "zs10znh5fxagl4z2efdy2rltgas9aahjscjuj3slsjyk96zfn5s8vu3sz6u8s4jkyl9zswp2ucm68j",
    "private_key": "secret-extended-key-main1q0eaejzrqqqqpqpl5kj9676vn6dx4ul0s8vc2xhqu3g2f22r8494l0sjkega75nvhupyasuxrfyj2usr2g8ru2uv4y8d88g3xtrhg0jvcuzgy50wp3dsdnfs3nxjaj2qvpswg93x0e5sety25d6ktcgzkc7ntxq9rg60mfcq9fh8gp97h8aw8ccvn74z68tps6d43ukww4f55k6rhm2322sc02ugq7tur0e9kpj34tevyeej4h38dfz6ktj4thtv3alg0eydkm4rrlssautyr",
    "seed": {
      "HDSeed": "fe50eb2add6c3e1ecc550f901fa737cbebab7b7a1dbf6827d4b0fd3521d2f93e",
      "path": "m/32'/141'/0'"
    }
  }
]
```

#### Cross-compiling Go example code

macOS x86_64:

```shell
env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin/ -lsaplinglib -framework Security" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build
```

macOS arm64:

```shell
env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64/ -lsaplinglib -framework Security" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build
```

Windows:
There must be extra parameter provided `CC` to specifiy which gcc binary to use to compile windows binary.

```shell
env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build
```