package zmachine

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	zork1 = "..\\ZCode\\Zork1\\zork1.z3"
)

func loadStory(filePath string, buffer *[]byte) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err == nil {
		*buffer = bytes
	}

	return err
}

func createZork1Memory() Memory {
	var buffer []byte
	loadStory(zork1, &buffer)
	return NewMemory(buffer)
}

func TestHeader_Zork1(t *testing.T) {
	memory := createZork1Memory()
	Convey("Zork 1: version should be 3", t, func() {
		So(memory.Version(), ShouldEqual, 3)
	})

	Convey("Zork 1: static memory base should be 0x2e53", t, func() {
		So(memory.StaticMemoryBase(), ShouldEqual, 0x2e53)
	})

	Convey("Zork 1: high memory base should be 0x4e37", t, func() {
		So(memory.HighMemoryBase(), ShouldEqual, 0x4e37)
	})

	Convey("Zork 1: initial program counter should be 0x4f05", t, func() {
		So(memory.InitialProgramCounter(), ShouldEqual, 0x4f05)
	})

	Convey("Zork 1: dictionary address should be 0x3b21", t, func() {
		So(memory.DictionaryAddress(), ShouldEqual, 0x3b21)
	})

	Convey("Zork 1: object table address should be 0x02b0", t, func() {
		So(memory.ObjectTableAddress(), ShouldEqual, 0x02b0)
	})

	Convey("Zork 1: global variables table address should be 0x2271", t, func() {
		So(memory.GlobalVariablesTableAddress(), ShouldEqual, 0x2271)
	})

	Convey("Zork 1: abbreviations table address should be 0x01f0", t, func() {
		So(memory.AbbreviationsTableAddress(), ShouldEqual, 0x01f0)
	})
}
