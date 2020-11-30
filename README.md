# Kota-ExConverter
CLI to convert specified jpg/png/gif format image to jpg/png/gif format

## GLOBAL OPTIONS

```
  --help, -h     show help (default: false)
  --version, -v  print the version (default: false)
```

## HOW TO USE
Install Homebrew if you have not yet

### INSTALLATION

```
$ brew tap kota-yata/kec
$ brew install kec
```

#### ARGUMENTS

```go
$ kec [source image file path] [distribution path]
// example
$ kec /Users/mike/Download/src.png /Users/mike/Download/dist/dist.jpg
```
- Both URI and relative are OK
- Since image format will be detected automatically, you don't need to specify it.

## DEVELOPMENT

```
$ git clone git@github.com:kota-yata/kec.git
```

before creating PR

```
$ go build
```
