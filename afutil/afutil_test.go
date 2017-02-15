package afutil

import (
	"os"
	"testing"
)

func setup() {
	os.Mkdir("tests", 0644)
}

// Testing non existence of not required files.
func TestNonExistence(t *testing.T) {
	setup()

	var paramTests = []struct {
		params string
	}{
		{"tests"},
	}

	for _, dir := range paramTests {
		valid, err := ParseDir(dir.params)
		if err == nil {
			t.Errorf("ParseDir(%q) failed test. \n\n\rWant: \n\r\"%t\" \n\rGot: \n\r\"%t\" \n\n", dir, false, valid)
		}
		if valid != false {
			t.Errorf("ParseDir(%q) failed test. \n\n\rWant: \n\r\"%t\" \n\rGot: \n\r\"%t\" \n\n", dir, false, valid)
		}
	}

	tearDown()
}

// Testing if the host user has the dpkg command.
func TestDpkgCommandExists(t *testing.T) {
	err := CheckDpkg()
	if err != nil {
		t.Errorf("CheckDpkg() failed test. \n\n\rWant: \n\r\"%v\" \n\rGot: \n\r\"%v\" \n\n", nil, err)
	}
}

// Testing .deb file regex
func TestIsDeb(t *testing.T) {
	var paramTests = []struct {
		params string
	}{
		{"test.deb"},
		{"test.xy.deb"},
		{"com.test.deb"},
		{"com.test.xy.deb"},
		{"com..example.test.deb"},
		{"com..example.test.xy.deb"},
		{"com..example.test..deb"},
		{"com..example.test..xy.deb"},
		{"com..example.test...de.deb"},
		{"com..example.test...xy.de.deb"},
	}

	for _, deb := range paramTests {
		valid := IsDeb(deb.params)
		if valid != true {
			t.Errorf("IsDeb(%q) failed test. \n\n\rWant: \n\r\"%t\" \n\rGot: \n\r\"%t\" \n\n", deb.params, true, valid)
		}
	}
}

// Testing if the host user has the bzip2 command.
func TestBzip2Exists(t *testing.T) {
	err := CheckBzip2()
	if err != nil {
		t.Errorf("CheckBzip2() failed test. \n\n\rWant: \n\r\"%v\" \n\rGot: \n\r\"%v\" \n\n", nil, err)
	}
}

// Testing bzip2 packages using exec.
func TestBzipPackages(t *testing.T) {
	os.Chdir("../test_data/packages/")
	err := BzipPackages()
	if err != nil {
		t.Errorf("BzipPackages() failed test. \n\n\rWant: \n\r\"%v\" \n\rGot: \n\r\"%v\" \n\n", nil, err)
	}
	_, direrr := os.Open("Packages.bz2")
	if direrr != nil {
		t.Errorf("BzipPackages() failed test. Packages.bz2 was not found.")
	}
	tearDown()
	os.Chdir("../")
}

// Testing the dir for deb file existence.
func TestCheckDeb(t *testing.T) {
	os.Chdir("../test_data/deb/")
	debs, err := CheckDeb()
	if err != nil {
		t.Errorf("CheckDeb() failed test. \n\n\rWant: \n\r\"%v\" \n\rGot: \n\r\"%v\" \n\n", nil, err)
	}
	if len(debs) != 1 {
		t.Errorf("CheckDeb() failed test. deb file length mismatch.")
	}
	os.Chdir("../")
}

// Testing the dir for deb file existence with a path name.
func TestCheckDebWithPath(t *testing.T) {
	os.Chdir("../test_data")
	debs, err := CheckDebWithPath("deb")
	if err != nil {
		t.Errorf("CheckDebWithPath() failed test. \n\n\rWant: \n\r\"%v\" \n\rGot: \n\r\"%v\" \n\n", nil, err)
	}
	if len(debs) != 1 {
		t.Errorf("CheckDebWithPath() failed test. deb file length mismatch.")
	}
	os.Chdir("../")
}

func tearDown() {
	os.Remove("tests")
	os.Remove("Packages.bz2")
}
