package zmachine

// Memory represents the memory map of a Z-Machine.
type Memory struct {
	bytes []byte
}

func (m *Memory) readByte(index int) byte {
	return m.bytes[index]
}

func (m *Memory) readWord(index int) uint16 {
	return uint16(m.bytes[index])<<8 | uint16(m.bytes[index+1])
}

const (
	versionOffset               = 0x00
	highMemoryBaseOffset        = 0x04
	initialProgramCounterOffset = 0x06
	staticMemoryBaseOffset      = 0x0e
)

// Version returns the version number from Z-Machine memory.
func (m *Memory) Version() byte {
	return m.readByte(versionOffset)
}

// HighMemoryBase returns the byte address of the base of high memory.
func (m *Memory) HighMemoryBase() uint16 {
	return m.readWord(highMemoryBaseOffset)
}

// StaticMemoryBase returns the byte address of the base of high memory.
func (m *Memory) StaticMemoryBase() uint16 {
	return m.readWord(staticMemoryBaseOffset)
}

// InitialProgramCounter returns the byte address of initial program counter.
func (m *Memory) InitialProgramCounter() uint16 {
	return m.readWord(initialProgramCounterOffset)
}

// NewMemory creates an instance of Memory.
func NewMemory(bytes []byte) Memory {
	return Memory{
		bytes: bytes,
	}
}
