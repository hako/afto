// Package release provides a structure for a repo Release file.
package release

import (
	"crypto/md5"
	"fmt"
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
	signatures  []MD5Signature
}

// MD5Signature represents a signed repo Release file.
type MD5Signature struct {
	sum         string
	size        int
	packageName string
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

// AddPackageSignature appends an MD5 Signature of Packages & Packages.bz2 in the Release file.
// It should be in the form of:
// MD5Sum:
//  <hash> <size in bytes> Packages
//  <hash> <size in bytes> Packages.bz2
func (r *Release) AddPackageSignature(pkgs []byte, pkgbz2 []byte) {
	pkgsum := fmt.Sprintf("%x", md5.Sum(pkgs))
	pkgbz2sum := fmt.Sprintf("%x", md5.Sum(pkgbz2))

	pkgsig := MD5Signature{pkgsum, len(pkgs), "Packages"}
	pkgbz2sig := MD5Signature{pkgbz2sum, len(pkgbz2), "Packages.bz2"}

	r.signatures = []MD5Signature{pkgsig, pkgbz2sig}
}

// Generate creates a release file from the Release struct.
// It appends a signature at the end of the release file.
func (r Release) Generate() string {

	var release = `Origin: ` + r.origin + `
Label: ` + r.label + `
Suite: ` + r.suite + `
Version: ` + strconv.Itoa(r.version) + `
Codename: ` + r.codename + `
Architectures: ` + r.arch + `
Components: ` + r.components + `
Description: ` + r.description + `
MD5Sum:
`

	for _, s := range r.signatures {
		release += fmt.Sprintf(" %s %d %s\n", s.sum, s.size, s.packageName)
	}

	return release
}
