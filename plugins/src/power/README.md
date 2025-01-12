# Extism Go PDK Plugin

See more documentation at https://github.com/extism/go-pdk and
[join us on Discord](https://extism.org/discord) for more help.

## Compile the Plug-in

```
tinygo build -target wasi -o power.wasm main.go
```

## Running the Plug-in

```
extism call plugin.wasm greet --input "Christian" --wasi
```