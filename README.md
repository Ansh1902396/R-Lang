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
