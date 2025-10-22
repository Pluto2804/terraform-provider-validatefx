package validate

import (
	"regexp"
)

// semverRegex defines a regex for validating semantic version strings
var semverRegex = regexp.MustCompile(`^v?([0-9]+)\.([0-9]+)\.([0-9]+)(?:-[0-9A-Za-z\-.]+)?(?:\+[0-9A-Za-z\-.]+)?$`)

// Checks if s is a valid semantic version
// (supports optional 'v', pre-release, and build metadata)
func IsSemver(s string) bool {
	return semverRegex.MatchString(s)
}
