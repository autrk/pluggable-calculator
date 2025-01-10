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

* [tidy](#devbox-run-tidy)

## Environment

```sh
GOPATH="${HOME}/go/"
PATH="${PATH}:${HOME}/go/bin"
```

## Shell Init Hook
The Shell Init Hook is a script that runs whenever the devbox environment is instantiated. It runs 
on `devbox shell` and on `devbox run`.
```sh
export "GOROOT=$(go env GOROOT)"
```

## Packages

* [go@latest](https://www.nixhub.io/packages/go)
* [cobra-cli@latest](https://www.nixhub.io/packages/cobra-cli)

## Script Details

### devbox run tidy
```sh
go mod tidy
```
&ensp;



<!-- gen-readme end -->
