all: build

dist: darwin darwin_arm64 win64 linux64

build: src/lib.rs Cargo.toml
	cargo build --lib --release
	@cp -av target/release/libsaplinglib.a .
	@rm -rf target

# Test binaries using static library file
cbin: build
	gcc src/C/main.c -I./src -L./ -lsaplinglib -lpthread -ldl -o csapling -framework Security
	./csapling
	@rm -rf target
	@rm -rf Cargo.lock

gobin: build
	go build -o gosapling src/Go/main.go
	./gosapling
	@rm -rf target
	@rm -rf Cargo.lock

# Distributable release libraries
darwin:
	@mkdir -p dist/darwin
	rustup target add x86_64-apple-darwin
	cargo build --target=x86_64-apple-darwin --release
	@cp -av target/x86_64-apple-darwin/release/libsaplinglib.a dist/darwin/
	@rm -rf target
	@rm -rf Cargo.lock

darwin_arm64:
	@mkdir -p dist/darwin_arm64
	rustup target add aarch64-apple-darwin
	cargo build --target=aarch64-apple-darwin --release
	@cp -av target/aarch64-apple-darwin/release/libsaplinglib.a dist/darwin_arm64/
	@rm -rf target
	@rm -rf Cargo.lock

win64:
	@mkdir -p dist/win64
	rustup target add x86_64-pc-windows-gnu
	cargo build --target=x86_64-pc-windows-gnu --release
	@cp -av target/x86_64-pc-windows-gnu/release/libsaplinglib.a dist/win64/
	@rm -rf target
	@rm -rf Cargo.lock

linux64:
	@mkdir -p dist/linux
	rustup target add x86_64-unknown-linux-gnu
	cargo build --target=x86_64-unknown-linux-gnu --release
	@cp -av target/x86_64-unknown-linux-gnu/release/libsaplinglib.a dist/linux/
	@rm -rf target
	@rm -rf Cargo.lock


clean:
	@rm -rf target
	@rm -rf dist
	@rm -rf csapling
	@rm -rf gosapling
	@rm -rf Cargo.lock
	@rm -rf libsaplinglib.a