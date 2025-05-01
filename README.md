<img src="docs/attachments/verso_logo_banner.svg" />

# verso

Simple [CHANGELOG]("https://keepachangelog.com") [semver]("https://semver.org/") extractor library written in Go language. If you are interested in a CLI using verso have a look at [verso.cli](https://github.com/hkionline/verso.cli)

## Usage

Changelogs can be parsed either by supplying a file path or bytes from a file, assumed to contain a changelog formatted data.

***Parse a file from file path:***
```go
import "github.com/hkionline/verso"

changelog, err := verso.ParsePath("path/to/CHANGELOG.md")
```

***Parse a file from bytes:***
```go
import "github.com/hkionline/verso"

var bytes []byte

// code to read a file, e.g. stdin

changelog, err := verso.ParseBytes(bytes)
```
