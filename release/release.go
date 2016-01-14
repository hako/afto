// Package release provides a structure for a repo Release file.
package release

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

// NewRelease creates a new Release struct for a Release file.
func NewRelease() *Release {
	return &Release{}
}

// Origin return the name of the Cydia repository.
func (r *Release) Origin() string {
	return r.origin
}

// Label returns the respository and section of the packages.
func (r *Release) Label() string {
	return r.label
}

// Suite returns whether the package is beta, unstable or stable.
func (r *Release) Suite() string {
	return r.suite
}

// Version returns the version number of the repo. (saruik questioned it's purpose...)
func (r *Release) Version() int {
	return r.version
}

// Codename returns the codename of the repo such as 'tangelo' or 'hakobyt'
func (r *Release) Codename() string {
	return r.codename
}

// Arch returns repo with all the supported architectures. (iphoneos-arm for 2.x or darwin-arm or 1.1.x)
func (r *Release) Arch() string {
	return r.arch
}

// Components returns the repo components name usually main.
func (r *Release) Components() string {
	return r.Components()
}

// Description returns the Cydia repo description.
func (r *Release) Description() string {
	return r.description
}
