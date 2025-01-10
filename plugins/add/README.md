# Extism Zig PDK Plugin

See more documentation at https://github.com/extism/zig-pdk and
[join us on Discord](https://extism.org/discord) for more help.

## Compile the Plug-in

```
zig build
```

## Running the Plug-in

```
extism call ./zig-out/bin/zig-pdk-template.wasm greet --input "Christian"
# => Hello, Christian!
```