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
	err := CheckDPKG()
	if err != nil {
		t.Errorf("CheckDPKG() failed test. \n\n\rWant: \n\r\"%v\" \n\rGot: \n\r\"%v\" \n\n", nil, err)
	}
	tearDown()
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

func tearDown() {
	os.Remove("tests")
}
