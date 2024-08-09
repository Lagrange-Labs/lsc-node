package core

import "fmt"

// Version is the version of the node.
type Version struct {
	// Major is the major version.
	Major uint32 `json:"major"`
	// Minor is the minor version.
	Minor uint32 `json:"minor"`
	// Patch is the patch version.
	Patch uint32 `json:"patch"`
}

// NewVersion creates a new version.
func NewVersion(major, minor, patch uint32) *Version {
	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

// String returns the version string.
func (v *Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Compare compares two versions.
func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		return int(v.Major) - int(other.Major)
	}
	if v.Minor != other.Minor {
		return int(v.Minor) - int(other.Minor)
	}
	if v.Patch != other.Patch {
		return int(v.Patch) - int(other.Patch)
	}
	return 0
}

// GetVersion returns the version from the string.
func GetVersion(version string) (*Version, error) {
	var major, minor, patch uint32
	_, err := fmt.Sscanf(version, "v%d.%d.%d", &major, &minor, &patch)
	if err != nil {
		return nil, err
	}
	return NewVersion(major, minor, patch), nil
}
