package model

import (
	"github.com/coreos/go-semver/semver"
)

var (
	// GitRepository is the git repository that was compiled.
	GitRepository string
	// GitCommit is the git commit that was compiled.
	GitCommit string
	// VersionPre indicates prerelease.
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
)

// Version is the specification version that the package types support.
func Version() semver.Version {
	return semver.Version{
		// VersionMajor is for an API incompatible changes.
		Major: 0,
		// VersionMinor is for functionality in a backwards-compatible manner.
		Minor: 1,
		// VersionPatch is for backwards-compatible bug fixes.
		Patch: 0,
		// VersionPre indicates prerelease.
		PreRelease: semver.PreRelease(VersionPre),
		// VersionDev indicates development branch. Releases will be empty string.
		Metadata: VersionDev,
	}
}
