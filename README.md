# verso

Simple [CHANGELOG]("https://keepachangelog.com") [semver]("https://semver.org/") extractor CLI and libray.

## Usage

### Usage as CLI

Get the latest version in the CHANGELOG.
```bash
verso latest
```

### Usage as library package

```go
import "github.com/hkionline/verso/pkg/verso"

changelog, err := verso.Parse("path/to/CHANGELOG.md")

```
