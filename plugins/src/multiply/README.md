# Extism Rust PDK Plugin

See more documentation at https://github.com/extism/rust-pdk and
[join us on Discord](https://extism.org/discord) for more help.

## Compile the Plug-in

```
cargo build --target wasm32-unknown-unknown
```

## Running the Plug-in

```
extism call target/wasm32-unknown-unknown/debug/rust_pdk_template.wasm greet --input "Christian"
# => Hello, Christian!
```

## Optimize using release mode

You will get better performance and smaller .wasm binaries if you build your code using --release.

```
cargo build --target wasm32-unknown-unknown --release
```