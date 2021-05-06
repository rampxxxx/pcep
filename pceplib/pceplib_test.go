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
func TestType(t *testing.T) {
	var i I
	var h Header
	//var d Detail
	i = h
	dataType := GetType(i)

	if dataType != PCEPLIB_TYPE_HEADER {
		t.Error("Type expected ", PCEPLIB_TYPE_HEADER, " got ", dataType)
	}
}
