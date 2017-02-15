// Package deb provides a structure for a parsing a debian control file.
package deb

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// DpkgInterface is an interface for control files and Package files in a existing repo.
type DpkgInterface interface {
	Package() string
	Name() string
	Depends() string
	Arch() string
	Description() string
	Maintainer() string
	Author() string
	Section() string
	Version() string
	InstalledSize() int
	Depiction() string
	Homepage() string
	Sponsor() string
	Filename() string
	Size() int
	MD5Sum() string
	SHA1() string
	SHA256() string

	ParseString()
}

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

// Packages represents a structure of a tweak Packages file.
type Packages struct {
	packageID     string
	version       string
	arch          string
	maintainer    string
	installedsize int
	homepage      string
	depends       string
	filename      string
	size          int
	md5sum        string
	sha1          string
	sha256        string
	section       string
	description   string
	depiction     string
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

// NewPackages creates a new Package struct to hold the Package data.
func NewPackages() *Packages {
	return &Packages{}
}

// Package

// Package returns the Cydia package name identifier. (com.example.tweakexample)
func (c *Control) Package() string {
	return c.packageID
}

// Package returns the Cydia package name identifier. (com.example.tweakexample)
func (p *Packages) Package() string {
	return p.packageID
}

// Name

// Name returns the Cydia tweak name. (TweakExample)
func (c *Control) Name() string {
	return c.name
}

// Name returns the Cydia tweak name. (TweakExample)
func (p *Packages) Name() string {
	return p.name
}

// Depends

// Depends returns the Cydia tweak dependencies and dependents. (mobilesubstrate...)
func (c *Control) Depends() string {
	return c.depends
}

// Depends returns the Cydia tweak dependencies and dependents. (mobilesubstrate...)
func (p *Packages) Depends() string {
	return p.depends
}

// Version

// Version returns the Cydia tweak version. (0.0.1)
func (c *Control) Version() string {
	return c.version
}

// Version returns the Cydia tweak version. (0.0.1)
func (p *Packages) Version() string {
	return p.version
}

// Architecture

// Arch returns the Cydia tweak compiled architecture. (iphoneos-arm)
func (c *Control) Arch() string {
	return c.arch
}

// Arch returns the Cydia tweak compiled architecture. (iphoneos-arm)
func (p *Packages) Arch() string {
	return p.arch
}

// Description

// Description returns the Cydia tweak description. (An awesome MobileSubstrate tweak!)
func (c *Control) Description() string {
	return c.description
}

// Description returns the Cydia tweak description. (An awesome MobileSubstrate tweak!)
func (p *Packages) Description() string {
	return p.description
}

// Maintainer

// Maintainer returns the Cydia tweak maintainer name. (foo)
func (c *Control) Maintainer() string {
	return c.maintainer
}

// Maintainer returns the Cydia tweak maintainer name. (foo)
func (p *Packages) Maintainer() string {
	return p.maintainer
}

// Author

// Author returns the Cydia tweak author name. (foo)
func (c *Control) Author() string {
	return c.author
}

// Author returns the Cydia tweak author name. (foo)
func (p *Packages) Author() string {
	return p.author
}

// Section

// Section returns the Cydia tweak section. (Tweaks)
func (c *Control) Section() string {
	return c.section
}

// Section returns the Cydia tweak section. (Tweaks)
func (p *Packages) Section() string {
	return p.section
}

// Installed Size

// InstalledSize returns an estimate of the total amount
// of disk space required to install the tweak (88)
func (c *Control) InstalledSize() int {
	return c.installedsize
}

// InstalledSize returns an estimate of the total amount
// of disk space required to install the tweak (88)
func (p *Packages) InstalledSize() int {
	return p.installedsize
}

// Homepage

// Homepage returns the Cydia tweak authors homepage. (http://example.com)
func (c *Control) Homepage() string {
	return c.homepage
}

// Homepage returns the Cydia tweak authors homepage. (http://example.com)
func (p *Packages) Homepage() string {
	return p.homepage
}

// Depiction

// Depiction returns the depiction URL of a cydia tweak. (http://example.com/depiction)
func (c *Control) Depiction() string {
	return c.depiction
}

// Depiction returns the depiction URL of a cydia tweak. (http://example.com/depiction)
func (p *Packages) Depiction() string {
	return p.depiction
}

// Sponsor

// Sponsor returns the sponsor author of cydia tweak. (<sponsor> http://example.com)
func (c *Control) Sponsor() string {
	return c.sponsor
}

// Sponsor returns the sponsor author of cydia tweak. (<sponsor> http://example.com)
func (p *Packages) Sponsor() string {
	return p.sponsor
}

// Filename returns the filepath of the cydia tweak.
func (p *Packages) Filename() string {
	return p.filename
}

// Size returns the total size in bytes of a single cydia tweak.
func (p *Packages) Size() int {
	return p.size
}

// MD5Sum returns the MD5 of a single cydia tweak.
func (p *Packages) MD5Sum() string {
	return p.md5sum
}

// SHA1 returns the SHA1 of a single cydia tweak.
func (p *Packages) SHA1() string {
	return p.sha1
}

// SHA256 returns the SHA256 of a single cydia tweak.
func (p *Packages) SHA256() string {
	return p.sha256
}

// ParseString parses a string f returns a *Control struct.
func (c *Control) ParseString(f string) (*Control, error) {
	parseMap := parse(f)
	if len(parseMap) == 0 {
		return nil, errors.New("unable to parse control file")
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

// ParseString parses a string f returns a *Packages struct.
func (p *Packages) ParseString(f string) (*Packages, error) {
	parseMap := parse(f)
	if len(parseMap) == 0 {
		return nil, errors.New("unable to parse Packages file")
	}

	if _, exists := parseMap["Homepage"]; exists != true {
		parseMap["Homepage"] = DefaultHomePage
	}
	if _, exists := parseMap["Sponsor"]; exists != true {
		parseMap["Sponsor"] = DefaultSponsor
	}

	// Necessary string conversion for Installed-Size.
	is, sterr := strconv.Atoi(parseMap["Installed-Size"])
	if sterr != nil {
		return nil, sterr
	}

	// Necessary string conversion for Installed-Size.
	size, serr := strconv.Atoi(parseMap["Size"])
	if serr != nil {
		return nil, serr
	}

	pkgs := &Packages{
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
		size:          size,
		md5sum:        parseMap["MD5sum"],
		sha1:          parseMap["SHA1"],
		sha256:        parseMap["SHA256"],
		filename:      parseMap["Filename"],
	}

	return pkgs, nil
}

// parse transforms a given control or packages file f and returns a key value map.
func parse(f string) map[string]string {
	var matches = make(map[string]string)

	keys := regexp.MustCompile(`(?m:^\w\S+)`)    // eg. should match: Packages:
	values := regexp.MustCompile(`(?m:( )\S.+)`) // eg. should match: com.yourcompany.tweakexample

	for c, v := range values.FindAllString(f, -1) {
		field := strings.Replace(keys.FindAllString(f, -1)[c], ":", "", -1)
		matches[field] = strings.Replace(v, " ", "", 1) // replace at the start.
	}

	return matches
}
