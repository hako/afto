// Package afutil includes utility functions for afto.
package afutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"

	"github.com/hako/afto/release"
)

// ParseDir checks to see if a directory has the required files for a cydia repo.
// might move to separate util package.
func ParseDir(indir string) (bool, error) {
	var reqfiles = []string{"Packages", "Release"}
	var valid = false
	var reqcount = 0

	files, err := ioutil.ReadDir(indir)
	if err != nil {
		return valid, errors.New("directory " + indir + " not found or valid.")
	}
	for _, file := range files {
		for _, reqfile := range reqfiles {
			if file.Name() == reqfile {
				reqcount++
			}
		}
	}

	if reqcount < len(reqfiles) {
		count := len(reqfiles) - reqcount
		return valid, errors.New(strconv.Itoa(count) + " required files missing. (Need Packages and Release)")
	}
	valid = true
	return valid, nil
}

// CheckDPKG checks the host system has the dpkg command installed.
func CheckDPKG() error {
	_, err := exec.LookPath("dpkg")
	if err != nil {
		return errors.New("unable to find required command 'dpkg'")
	}
	return nil
}

// CheckBzip2 checks the host system has the bzip2 command installed. (Sometimes this happens.)
func CheckBzip2() error {
	_, err := exec.LookPath("bzip2")
	if err != nil {
		return errors.New("unable to find required command 'bzip2'")
	}
	return nil
}

// GetRepo checks if the string dir matches a valid repo in the current directory.
// This can be from a name or a path.
func GetRepo(dir string) (string, error) {
	// Check if the input directory has the valid files for a cydia repo.
	_, direrr := ParseDir(dir)
	if direrr != nil {
		return "", errors.New(direrr.Error())
	}

	// Get the absolute path from dir.
	finalPath, abserr := filepath.Abs(dir)
	if abserr != nil {
		return "", errors.New(abserr.Error())
	}

	return finalPath, nil
}

// ParseDeb parses a deb file and lists all the fields in the deb file.
// It uses the command dpkg --field *.deb to
func ParseDeb(debName string) error {
	// Run dpkg --field *.deb
	fields, parseerr := exec.Command("dpkg", "-field", debName).Output()
	if parseerr != nil {
		return parseerr
	}
	fmt.Println(fields)
	return nil
}

// BzipPackages compresses the 'Packages' file Pacakges.bz2.
// Note: The stlib package "compress/bzip2" does not support compression.
func BzipPackages() error {
	// Run bzip2 -kf Packages.
	_, bzcmderr := exec.Command("bzip2", "-kf", "Packages").Output()
	if bzcmderr != nil {
		return bzcmderr
	}
	return nil
}

// CheckDeb checks if the user has deb files ready to go to the repo.
func CheckDeb() ([]string, error) {
	cwdir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(cwdir)
	if err != nil {
		return nil, err
	}

	var deb []string
	for _, file := range files {
		if IsDeb(file.Name()) == true {
			deb = append(deb, file.Name())
		}
	}
	// Return error if ultimately no deb files are found.
	if len(deb) == 0 {
		return nil, errors.New("No .deb file(s) found. Unable to continue.")
	}
	return deb, nil
}

// CheckDebWithPath is like CheckDeb but a path is required.
func CheckDebWithPath(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var deb []string
	for _, file := range files {
		if IsDeb(file.Name()) == true {
			deb = append(deb, file.Name())
		}
	}
	// Return error if ultimately no deb files are found.
	if len(deb) == 0 {
		return nil, errors.New("No .deb file(s) found. Unable to continue.")
	}
	return deb, nil
}

// IsDeb returns whether the string is a deb file with regex.
func IsDeb(filename string) bool {
	re := regexp.MustCompile(`(\.deb)`)
	return re.MatchString(filename)
}

// ReleaseFile generates a release file based on origin, label, desc codename and suite.
// It is recommended to generate this file for hosting a repo.
func ReleaseFile(origin string, label string, desc string, codename string, suite string) (string, error) {
	r := release.NewRelease()
	r.SetOrigin(origin)
	r.SetLabel(label)
	r.SetDescription(desc)
	r.SetCodename(codename)
	r.SetSuite(suite)

	// Get Packages and Packages.bz2
	packages, err := ioutil.ReadFile("Packages")
	if err != nil {
		return "", err
	}
	packagesbz, err := ioutil.ReadFile("Packages.bz2")
	if err != nil {
		return "", err
	}

	r.AddPackageSignature(packages, packagesbz)
	return r.Generate(), nil
}

// DetectPlatform returns what host system the user is running.
func DetectPlatform() (string, error) {
	hostOS := runtime.GOOS
	switch hostOS {
	case "darwin":
		return "Since you're on a Mac, you can install dpkg via brew. \nIf you have brew installed type `brew install dpkg` and try again.", nil
	case "linux":
		return "Since you're on Linux, you can install dpkg by using `sudo apt-get install dpkg`.", nil
	case "windows":
		return "It looks like dpkg is not available on Windows. \nYou can stil use afto, you cannot build your repo without dpkg.", nil
	default:
		return "", errors.New("Unable to detect your OS.")
	}
}
