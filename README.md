# R-Lang

R-Lang is a small interpreted language written in Go.

## Run The REPL

```sh
go run .
```

## Run A File

R-Lang source files use the `.rlang` extension.

```sh
go run . examples/hello.rlang
```

Run a larger algorithm example:

```sh
go run . examples/binary_search.rlang
```

After installing or downloading a release binary:

```sh
rlang examples/hello.rlang
```

## Release

Create and push a version tag to build binaries and publish a GitHub Release:

```sh
git tag -a v0.1.0 -m "R-Lang v0.1.0"
git push origin v0.1.0
```

The release workflow runs tests, builds macOS, Linux, and Windows binaries, and attaches them to the GitHub Release.

## VS Code Syntax Highlighting

This repo includes a local VS Code extension for `.rlang` files in `vscode-extension/`.

macOS/Linux:

```sh
mkdir -p ~/.vscode/extensions
ln -s "$(pwd)/vscode-extension" ~/.vscode/extensions/rlang-syntax
```

Then reload VS Code and open a `.rlang` file.
