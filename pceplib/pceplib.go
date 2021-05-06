package pceplib

import (
	"fmt"
)

const (
	PCEPLIB_NAME        = "pceplib"
	PCEPLIB_TYPE_HEADER = "Header"
	PCEPLIB_TYPE_DETAIL = "Detail"
	PCEPLIB_VERSION     = 1
)

type I interface {
	getType() string
}

type Header struct {
	Version       uint8
	Flags         uint8
	MessageType   uint8
	MessageLenght uint16
}

type Detail struct {
	Version       uint8
	Flags         uint8
	MessageType   uint8
	MessageLenght uint16
}

func (d Detail) getType() string {
	return PCEPLIB_TYPE_DETAIL
}

func (d Header) getType() string {
	return PCEPLIB_TYPE_HEADER
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

func GetType(i I) string {
	var dataType string
	switch v := i.(type) {
	case Header:
		fmt.Printf("Header Type %T\n", v)
		dataType = v.getType()
	case Detail:
		fmt.Printf("Detail Type %T\n", v)
		dataType = v.getType()
	default:
		fmt.Printf("Default don't know Type %T\n", v)
		dataType = v.getType()
	}
	return dataType
}
