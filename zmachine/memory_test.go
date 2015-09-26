package zmachine

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	adventFilePath = "..\\ZCode\\advent\\Advent.z5"
	zork1FilePath  = "..\\ZCode\\Zork1\\zork1.z3"
	rotaFilePath   = "..\\ZCode\\rota\\Rota.z8"
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

func createMemory(filePath string) Memory {
	var buffer []byte
	loadStory(filePath, &buffer)
	return NewMemory(buffer)
}

func TestHeader_Zork1(t *testing.T) {
	memory := createMemory(zork1FilePath)

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

	Convey("Zork 1: file size should be 0x14b8c", t, func() {
		So(memory.FileSize(), ShouldEqual, 0x14b8c)
	})

	Convey("Zork 1: file checksum should be 0xa129", t, func() {
		So(memory.Checksum(), ShouldEqual, 0xa129)
	})
}

func TestHeader_Advent(t *testing.T) {
	memory := createMemory(adventFilePath)

	Convey("Advent: version should be 5", t, func() {
		So(memory.Version(), ShouldEqual, 5)
	})

	Convey("Advent: static memory base should be 0x45c8", t, func() {
		So(memory.StaticMemoryBase(), ShouldEqual, 0x45c8)
	})

	Convey("Advent: high memory base should be 0x6a30", t, func() {
		So(memory.HighMemoryBase(), ShouldEqual, 0x6a30)
	})

	Convey("Advent: initial program counter should be 0x6a31", t, func() {
		So(memory.InitialProgramCounter(), ShouldEqual, 0x6a31)
	})

	Convey("Advent: dictionary address should be 0x4e87", t, func() {
		So(memory.DictionaryAddress(), ShouldEqual, 0x4e87)
	})

	Convey("Advent: object table address should be 0x010a", t, func() {
		So(memory.ObjectTableAddress(), ShouldEqual, 0x010a)
	})

	Convey("Advent: global variables table address should be 0x3be9", t, func() {
		So(memory.GlobalVariablesTableAddress(), ShouldEqual, 0x3be9)
	})

	Convey("Advent: abbreviations table address should be 0x0042", t, func() {
		So(memory.AbbreviationsTableAddress(), ShouldEqual, 0x0042)
	})

	Convey("Advent: file size should be 0x21a18", t, func() {
		So(memory.FileSize(), ShouldEqual, 0x21a18)
	})

	Convey("Advent: file checksum should be 0x76bd", t, func() {
		So(memory.Checksum(), ShouldEqual, 0x76bd)
	})
}

func TestHeader_Rota(t *testing.T) {
	memory := createMemory(rotaFilePath)

	Convey("Rota: version should be 8", t, func() {
		So(memory.Version(), ShouldEqual, 8)
	})

	Convey("Rota: static memory base should be 0xd041", t, func() {
		So(memory.StaticMemoryBase(), ShouldEqual, 0xd041)
	})

	Convey("Rota: high memory base should be 0xfbf0", t, func() {
		So(memory.HighMemoryBase(), ShouldEqual, 0xfbf0)
	})

	Convey("Rota: initial program counter should be 0xfbf1", t, func() {
		So(memory.InitialProgramCounter(), ShouldEqual, 0xfbf1)
	})

	Convey("Rota: dictionary address should be 0xd88a", t, func() {
		So(memory.DictionaryAddress(), ShouldEqual, 0xd88a)
	})

	Convey("Rota: object table address should be 0x0138", t, func() {
		So(memory.ObjectTableAddress(), ShouldEqual, 0x0138)
	})

	Convey("Rota: global variables table address should be 0x9b60", t, func() {
		So(memory.GlobalVariablesTableAddress(), ShouldEqual, 0x9b60)
	})

	Convey("Rota: abbreviations table address should be 0x0070", t, func() {
		So(memory.AbbreviationsTableAddress(), ShouldEqual, 0x0070)
	})

	Convey("Rota: file size should be 0x47288", t, func() {
		So(memory.FileSize(), ShouldEqual, 0x47288)
	})

	Convey("Rota: file checksum should be 0x04bf", t, func() {
		So(memory.Checksum(), ShouldEqual, 0x04bf)
	})
}
