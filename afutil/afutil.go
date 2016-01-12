// Package afutil includes utility functions for afto.
package afutil

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
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

// IsDeb returns whether the string is a deb file with regex.
func IsDeb(filename string) bool {
	re := regexp.MustCompile(`(\.deb)`)
	return re.MatchString(filename)
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
