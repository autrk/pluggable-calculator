build:
	tinygo build -target wasi -o power.wasm main.go

test:
	extism call power.wasm greet --input "Christian" --wasi
