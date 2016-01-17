// Package release provides a structure for a repo Release file.
package release

import (
	"strconv"
)

// Release represents a structure of a repo Release file.
type Release struct {
	origin      string
	label       string
	suite       string
	version     int
	codename    string
	arch        string
	components  string
	description string
}

// NewRelease creates a new Release struct for a Release file with default values.
func NewRelease() *Release {
	return &Release{
		suite:      "beta",
		arch:       "iphoneos-arm",
		components: "main",
	}
}

// Origin returns the name of the Cydia repository.
func (r Release) Origin() string {
	return r.origin
}

// Label returns the respository and section of the packages.
func (r Release) Label() string {
	return r.label
}

// Suite returns whether the package is beta, unstable or stable.
func (r Release) Suite() string {
	return r.suite
}

// Version returns the version number of the repo. (saruik questioned it's purpose...)
func (r Release) Version() int {
	return r.version
}

// Codename returns the codename of the repo such as 'tangelo' or 'hakobyt'
func (r Release) Codename() string {
	return r.codename
}

// Arch returns repo with all the supported architectures. (iphoneos-arm for 2.x or darwin-arm or 1.1.x)
func (r Release) Arch() string {
	return r.arch
}

// Components returns the repo components name usually main.
func (r Release) Components() string {
	return r.components
}

// Description returns the Cydia repo description.
func (r Release) Description() string {
	return r.description
}

// SetOrigin sets the name of the Cydia repository.
func (r *Release) SetOrigin(origin string) {
	r.origin = origin
}

// SetLabel sets the respository and section of the packages.
func (r *Release) SetLabel(label string) {
	r.label = label
}

// SetSuite sets whether the package is beta, unstable or stable.
func (r *Release) SetSuite(suite string) {
	r.suite = suite
}

// SetVersion sets the version number of the repo. (saruik questioned it's purpose...)
func (r Release) SetVersion(version int) {
	r.version = version
}

// SetCodename sets the codename of the repo such as 'tangelo' or 'hakobyt'
func (r *Release) SetCodename(codename string) {
	r.codename = codename
}

// SetArch sets repo with all the supported architectures. (iphoneos-arm for 2.x or darwin-arm or 1.1.x)
func (r *Release) SetArch(arch string) {
	r.arch = arch
}

// SetComponents sets the repo components name usually main.
func (r *Release) SetComponents(comp string) {
	r.components = comp
}

// SetDescription sets the Cydia repo description.
func (r *Release) SetDescription(desc string) {
	r.description = desc
}

// Generate creates a release file from the Release struct.
func (r Release) Generate() string {
	var release = `Origin: ` + r.origin + `
Label: ` + r.label + `
Suite: ` + r.suite + `
Version: ` + strconv.Itoa(r.version) + `
Codename: ` + r.codename + `
Architectures: ` + r.arch + `
Components: ` + r.components + `
Description: ` + r.description + `
`
	return release
}
