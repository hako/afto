// Package deb provides a structure for a parsing a debian control file.
package deb

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// AFTODO: An interface.

// Control represents a structure of a debian/cydia tweak control file.
type Control struct {
	packageID     string
	name          string
	depends       string
	version       string
	arch          string
	description   string
	homepage      string
	depiction     string
	maintainer    string
	author        string
	sponsor       string
	section       string
	installedsize int
}

// Package represents a structure of a tweak Packages file.
type Package struct {
	packageID     string
	version       string
	arch          string
	maintainer    string
	installedsize int
	depends       string
	filename      string
	size          int
	md5sum        string
	sha1          string
	sha256        string
	section       string
	description   string
	author        string
	sponsor       string
	name          string
}

var (
	//DefaultHomePage changes the default homepage of a package.
	DefaultHomePage = "None"

	//DefaultSponsor changes the default sponsor of a package.
	DefaultSponsor = "None"
)

// NewControl creates a new Format struct for the control file.
func NewControl() *Control {
	return &Control{}
}

// NewPackage creates a new Package struct to hold the Package data.
func NewPackage() *Package {
	return &Package{}
}

// Package returns the Cydia package name identifier. (com.example.tweakexample)
func (c *Control) Package() string {
	return c.packageID
}

// Name returns the Cydia tweak name. (TweakExample)
func (c *Control) Name() string {
	return c.name
}

// Depends returns the Cydia tweak dependencies and dependents. (mobilesubstrate...)
func (c *Control) Depends() string {
	return c.depends
}

// Version returns the Cydia tweak version. (0.0.1)
func (c *Control) Version() string {
	return c.version
}

// Arch returns the Cydia tweak compiled architecture. (iphoneos-arm)
func (c *Control) Arch() string {
	return c.arch
}

// Description returns the Cydia tweak description. (An awesome MobileSubstrate tweak!)
func (c *Control) Description() string {
	return c.description
}

// Maintainer returns the Cydia tweak maintainer name. (foo)
func (c *Control) Maintainer() string {
	return c.maintainer
}

// Author returns the Cydia tweak author name. (foo)
func (c *Control) Author() string {
	return c.author
}

// Section returns the Cydia tweak section. (Tweaks)
func (c *Control) Section() string {
	return c.section
}

// InstalledSize returns an estimate of the total amount
// of disk space required to install the tweak (88)
func (c *Control) InstalledSize() int {
	return c.installedsize
}

// Homepage returns the Cydia tweak authors homepage. (http://example.com)
func (c *Control) Homepage() string {
	return c.homepage
}

// Depiction returns the depiction URL of a cydia tweak. (http://example.com/depiction)
func (c *Control) Depiction() string {
	return c.depiction
}

// ParseString parses a string p returns a *Control.
func (c *Control) ParseString(p string) (*Control, error) {
	parseMap := parse(p)
	if len(parseMap) == 0 {
		return nil, errors.New("Unable to parse control file.")
	}

	if _, exists := parseMap["Homepage"]; exists != true {
		parseMap["Homepage"] = DefaultHomePage
	}
	if _, exists := parseMap["Sponsor"]; exists != true {
		parseMap["Sponsor"] = DefaultSponsor
	}

	// Necessary string conversion.
	is, sterr := strconv.Atoi(parseMap["Installed-Size"])
	if sterr != nil {
		return nil, sterr
	}

	cntrl := &Control{
		packageID:     parseMap["Package"],
		name:          parseMap["Name"],
		depends:       parseMap["Depends"],
		version:       parseMap["Version"],
		arch:          parseMap["Architecture"],
		description:   parseMap["Description"],
		homepage:      parseMap["Homepage"],
		depiction:     parseMap["Depiction"],
		maintainer:    parseMap["Maintainer"],
		author:        parseMap["Author"],
		sponsor:       parseMap["Sponsor"],
		section:       parseMap["Section"],
		installedsize: is,
	}

	return cntrl, nil
}

// parse parse the given control file f and returns a *Control struct.
func parse(f string) map[string]string {
	var matches = make(map[string]string)

	keys := regexp.MustCompile(`(?m:^\w\S+)`)    // eg. should match: Packages:
	values := regexp.MustCompile(`(?m:( )\b.+)`) // eg. should match: com.yourcompany.tweakexample

	for c, v := range values.FindAllString(f, -1) {
		field := strings.Replace(keys.FindAllString(f, -1)[c], ":", "", -1)
		matches[field] = strings.Replace(v, " ", "", 1) // replace at the start.
	}

	return matches
}
