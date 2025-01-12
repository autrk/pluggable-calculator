<!-- gen-readme start - generated by https://github.com/jetify-com/devbox/ -->
# go devbox

Instant, easy, and predictable go development environment

## Getting Started
This project uses [devbox](https://github.com/jetify-com/devbox) to manage its development environment.

Install devbox:
```sh
curl -fsSL https://get.jetpack.io/devbox | bash
```

Start the devbox shell:
```sh 
devbox shell
```

Run a script in the devbox environment:
```sh
devbox run <script>
```
## Scripts
Scripts are custom commands that can be run using this project's environment. This project has the following scripts:

* [build-add](#devbox-run-build-add)
* [build-multiply](#devbox-run-build-multiply)
* [build-power](#devbox-run-build-power)
* [build-substract](#devbox-run-build-substract)

## Environment

```sh
GOPATH="${HOME}/go/"
LIBRARY_PATH="/home/ckihm/workspace-ckl/wasm/pluggable-calculator/.devbox/nix/profile/default/lib"
PATH="${PATH}:${HOME}/go/bin"
RUSTUP_HOME="/home/ckihm/workspace-ckl/wasm/pluggable-calculator/.devbox/virtenv/rustup"
```

## Shell Init Hook
The Shell Init Hook is a script that runs whenever the devbox environment is instantiated. It runs 
on `devbox shell` and on `devbox run`.
```sh
test -z $DEVBOX_COREPACK_ENABLED || corepack enable --install-directory "/home/ckihm/workspace-ckl/wasm/pluggable-calculator/.devbox/virtenv/nodejs/corepack-bin/"
test -z $DEVBOX_COREPACK_ENABLED || export PATH="/home/ckihm/workspace-ckl/wasm/pluggable-calculator/.devbox/virtenv/nodejs/corepack-bin/:$PATH"
export "GOROOT=$(go env GOROOT)"
rustup default stable
cargo fetch
```

## Packages

* [go@latest](https://www.nixhub.io/packages/go)
* [cobra-cli@latest](https://www.nixhub.io/packages/cobra-cli)
* [extism-cli@latest](https://www.nixhub.io/packages/extism-cli)
* [zig@latest](https://www.nixhub.io/packages/zig)
* [rustup@latest](https://www.nixhub.io/packages/rustup)
* [nodejs@latest](https://www.nixhub.io/packages/nodejs)
* [tinygo@latest](https://www.nixhub.io/packages/tinygo)

## Script Details

### devbox run build-add
```sh
pushd ./plugins/src/add > /dev/null
zig build
cp ./zig-out/bin/add.wasm ../../wasm
popd > /dev/null
```
&ensp;

### devbox run build-multiply
```sh
pushd ./plugins/src/multiply > /dev/null
cargo build --target wasm32-unknown-unknown --release
cp ./target/wasm32-unknown-unknown/release/multiply.wasm ../../wasm
popd > /dev/null
```
&ensp;

### devbox run build-power
```sh
pushd ./plugins/src/power > /dev/null
tinygo build -target wasi -o power.wasm main.go
cp ./power.wasm ../../wasm
popd > /dev/null
```
&ensp;

### devbox run build-substract
```sh
pushd ./plugins/src/substract > /dev/null
npm run build
cp ./dist/substract.wasm ../../wasm
popd > /dev/null
```
&ensp;



<!-- gen-readme end -->
