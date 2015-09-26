package zmachine

// Memory represents the memory map of a Z-Machine.
type Memory struct {
	bytes []byte
}

// NewMemory creates an instance of Memory.
func NewMemory(bytes []byte) Memory {
	return Memory{
		bytes: bytes,
	}
}

func readByte(bytes []byte, address uint) byte {
	return bytes[address]
}

func readWord(bytes []byte, address uint) uint16 {
	return uint16(bytes[address])<<8 | uint16(bytes[address+1])
}

// ByteAddress represents an address to a byte in memory in the range 0
// up to the last byte of static memory.
type ByteAddress uint16

// WordAddress represents an even address in the bottom 128K of memory (by
// giving the address divided by 2).
type WordAddress uint16

// PackedAddress specifies where a routine or string begins in high memory.
type PackedAddress uint16

func (m *Memory) unpackRoutineAddress(address PackedAddress) uint {
	switch m.Version() {
	case 1, 2, 3:
		return uint(address) * 2
	case 4, 5:
		return uint(address) * 4
	case 6, 7:
		return uint(address) + (uint(m.RoutinesOffset()) * 8)
	case 8:
		return uint(address) * 8
	default:
		panic("Invalid Z-Machine version")
	}
}

func (m *Memory) unpackStringAddress(address PackedAddress) uint {
	switch m.Version() {
	case 1, 2, 3:
		return uint(address) * 2
	case 4, 5:
		return uint(address) * 4
	case 6, 7:
		return uint(address) + (uint(m.StringsOffset()) * 8)
	case 8:
		return uint(address) * 8
	default:
		panic("Invalid Z-Machine version")
	}
}

func (m *Memory) readByte(address uint) byte {
	return readByte(m.bytes, address)
}

func (m *Memory) readWord(address uint) uint16 {
	return readWord(m.bytes, address)
}

const (
	versionOffset                     = 0x00
	highMemoryBaseOffset              = 0x04
	initialProgramCounterOffset       = 0x06
	dictionaryAddressOffset           = 0x08
	objectTableAddressOffset          = 0x0a
	globalVariablesTableAddressOffset = 0x0c
	staticMemoryBaseOffset            = 0x0e
	abbreviationsTableAddressOffset   = 0x18
	fileSizeOffset                    = 0x1a
	checksumOffset                    = 0x1c
	routineOffset                     = 0x28
	stringsOffset                     = 0x2a
)

// Version returns the version number from Z-Machine memory.
func (m *Memory) Version() byte {
	return m.readByte(versionOffset)
}

// HighMemoryBase returns the byte address of the base of high memory.
func (m *Memory) HighMemoryBase() ByteAddress {
	w := m.readWord(highMemoryBaseOffset)
	return ByteAddress(w)
}

// StaticMemoryBase returns the byte address of the base of high memory.
func (m *Memory) StaticMemoryBase() ByteAddress {
	w := m.readWord(staticMemoryBaseOffset)
	return ByteAddress(w)
}

// InitialProgramCounter returns the byte address of initial program counter.
func (m *Memory) InitialProgramCounter() ByteAddress {
	w := m.readWord(initialProgramCounterOffset)
	return ByteAddress(w)
}

// DictionaryAddress returns the byte address of dictionary.
func (m *Memory) DictionaryAddress() ByteAddress {
	w := m.readWord(dictionaryAddressOffset)
	return ByteAddress(w)
}

// ObjectTableAddress returns the byte address of object table.
func (m *Memory) ObjectTableAddress() ByteAddress {
	w := m.readWord(objectTableAddressOffset)
	return ByteAddress(w)
}

// GlobalVariablesTableAddress returns the byte address of global variables table.
func (m *Memory) GlobalVariablesTableAddress() ByteAddress {
	w := m.readWord(globalVariablesTableAddressOffset)
	return ByteAddress(w)
}

// AbbreviationsTableAddress returns the byte address of abbreviations table.
func (m *Memory) AbbreviationsTableAddress() ByteAddress {
	w := m.readWord(abbreviationsTableAddressOffset)
	return ByteAddress(w)
}

// FileSize returns the length of the file.
func (m *Memory) FileSize() uint {
	w := m.readWord(fileSizeOffset)

	switch m.Version() {
	case 1, 2, 3:
		return uint(w) * 2
	case 4, 5:
		return uint(w) * 4
	case 6, 7, 8:
		return uint(w) * 8
	default:
		panic("Invalid Z-Machine version")
	}
}

// Checksum returns the checksum value.
func (m *Memory) Checksum() uint16 {
	return m.readWord(checksumOffset)
}

// RoutinesOffset returns the offset (divided by 8) used to unpack routine addresses.
func (m *Memory) RoutinesOffset() uint16 {
	return m.readWord(routineOffset)
}

// StringsOffset returns the offset (divided by 8) used to unpack string addresses.
func (m *Memory) StringsOffset() uint16 {
	return m.readWord(stringsOffset)
}
