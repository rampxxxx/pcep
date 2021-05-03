package pceplib

import (
	"testing"
)

func TestVersion(t *testing.T) {
	version := GetVersion()

	if version != PCEPLIB_VERSION {
		t.Error("Version expected ", PCEPLIB_VERSION, " got ", version)
	}
}
