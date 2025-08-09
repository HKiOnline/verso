package verso

import (
	"fmt"
	"strings"
	"time"
)

// Semver struct contains semantic version information
type Semver struct {
	Major      int       `json:"major"`
	Minor      int       `json:"minor"`
	Patch      int       `json:"patch"`
	PreRelease string    `json:"preRelease"`
	Build      string    `json:"build"`
	Date       time.Time `json:"date"`
}

// String implements [fmt.Stringer] interface. It returns Semver fields' values as a string.
//
// Major, Minor and Patch are mandatory fields and if all three are in their default values,
// an empty string is returned. PreRelease, Build and Date fields are optional and will not be
// printed if they are in their default values.
func (s Semver) String() string {

	if s.Major <= 0 && s.Minor <= 0 && s.Patch <= 0 {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteString("[")
	sb.WriteString(fmt.Sprint(s.Major))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(s.Minor))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(s.Patch))

	if s.PreRelease != "" {
		sb.WriteString("-")
		sb.WriteString(fmt.Sprint(s.PreRelease))
	}

	if s.Build != "" {
		sb.WriteString("+")
		sb.WriteString(fmt.Sprint(s.Build))
	}

	sb.WriteString("]")

	if !s.Date.IsZero() {
		sb.WriteString(" ")
		sb.WriteString(s.Date.Format(time.DateOnly))
	}

	return sb.String()
}
