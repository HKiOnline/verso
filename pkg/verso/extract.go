package verso

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func extract(content string) (Changelog, error) {

	changelog := Changelog{
		Versions: []Semver{},
	}

	// Regular expression to match version headers
	re := regexp.MustCompile(`## \[([\d\.]+)(?:-([\da-zA-Z\-\.]+))?(?:\+([\da-zA-Z\-\.]+))?\] - (\d{4}-\d{2}-\d{2})`)
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {

		// A match must have five and exactly five elements
		if len(match) != 5 {
			continue // Skip invalid matches
		}

		versionStr := match[1]
		preRelease := match[2]
		build := match[3]
		dateStr := match[4]

		// Parse version string
		parts := strings.Split(versionStr, ".")
		if len(parts) < 3 {
			continue // Skip invalid version strings
		}

		major, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		patch, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		// Parse date
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}

		// Create Semver struct
		semver := Semver{
			Major:      major,
			Minor:      minor,
			Patch:      patch,
			PreRelease: preRelease,
			Build:      build,
			Date:       date,
		}

		changelog.Versions = append(changelog.Versions, semver)
	}

	return changelog, nil
}
