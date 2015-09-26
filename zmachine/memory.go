package zmachine

// Memory represents the memory map of a Z-Machine.
type Memory struct {
	bytes []byte
}

// Version returns the version number from Z-Machine memory.
func (m *Memory) Version() byte {
	return m.bytes[0]
}

// NewMemory creates an instance of Memory.
func NewMemory(bytes []byte) Memory {
	return Memory{
		bytes: bytes,
	}
}
