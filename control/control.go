// Package control provides a structure for a tweak control file.
package control

import (
	// "errors"
	// "io"
	"net/url"
)

// Format represents a structure of a tweak control file.
type Format struct {
	packageID   string
	name        string
	depends     string
	version     string
	arch        string
	description string
	homepage    url.URL
	depiction   url.URL
	maintainer  string
	author      string
	sponsor     string
	section     string
}

// newFormat() creates a new tweak control parser format.
func newFormat() *Format {
	return &Format{}
}

// Package returns the Cydia package name identifier. (com.example.tweakexample)
func (f *Format) Package() string {
	return f.packageID
}

// Name returns the Cydia tweak name. (TweakExample)
func (f *Format) Name() string {
	return f.name
}

// Depends returns the Cydia tweak dependencies and dependents. (mobilesubstrate...)
func (f *Format) Depends() string {
	return f.depends
}

// Version returns the Cydia tweak version. (0.0.1)
func (f *Format) Version() string {
	return f.version
}

// Arch returns the Cydia tweak compiled architecture. (iphoneos-arm)
func (f *Format) Arch() string {
	return f.arch
}

// Description returns the Cydia tweak description. (An awesome MobileSubstrate tweak!)
func (f *Format) Description() string {
	return f.description
}

// Maintainer returns the Cydia tweak maintainer name. (foo)
func (f *Format) Maintainer() string {
	return f.maintainer
}

// Author returns the Cydia tweak author name. (foo)
func (f *Format) Author() string {
	return f.author
}

// Section returns the Cydia tweak section. (Tweaks)
func (f *Format) Section() string {
	return f.section
}

// // parseControlFile parses the given control file and returns a *Format struct.
// func parseControlFile() (*Format, errors) {
// }
