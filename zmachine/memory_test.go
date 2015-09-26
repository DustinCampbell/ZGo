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

func TestVersion_Zork1(t *testing.T) {
	memory := createZork1Memory()
	Convey("Zork 1: version should be 3", t, func() {
		So(memory.Version(), ShouldEqual, 3)
	})
}
