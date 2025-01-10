# Extism TS PDK Plugin

See more documentation at https://github.com/extism/js-pdk and
[join us on Discord](https://extism.org/discord) for more help.

## Install the dependency

```
curl -O https://raw.githubusercontent.com/extism/js-pdk/main/install.sh
sh install.sh
```

## Compile the Plug-in

```
npm run build
```

## Running the Plug-in

```
extism call ./dist/plugin.wasm greet --input "Christian" --wasi
# => Hello, Christian
```