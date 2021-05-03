package pceplib

import (
	"fmt"
)

const (
	PCEPLIB_NAME = "pceplib"
	PCEPLIB_VERSION = 1
)

type Header struct {
	Version       uint8
	Flags         uint8
	MessageType   uint8
	MessageLenght uint16
}

func FooPceplib() int {
	head := Header{Version: PCEPLIB_VERSION}
	var ptr *Header
	if ptr == nil {
		ptr = &head
	} else {
	}
	fmt.Println(PCEPLIB_NAME, "running version", ptr.Version)
	return 1
}

func GetVersion() int {

	return PCEPLIB_VERSION
}

