package verso

import "time"

// Semver struct contains semantic version information
type Semver struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
	Date       time.Time
}
