package osversion

import (
	"fmt"

	"github.com/blang/semver"
)

func GetSemanticVersion() (semver.Version, error) {
	str, err := GetString()
	if err != nil {
		return semver.Version{}, err
	}

	ver, err := semver.Make(str)
	if err != nil {
		return semver.Version{}, fmt.Errorf("invalid semantic version in %s: %w", str, err)
	}
	return ver, nil
}
