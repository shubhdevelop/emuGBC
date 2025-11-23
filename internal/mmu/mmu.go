package mmu

type MMU struct {
	// Represents the entire 64KB address space
	memory [0x10000]uint8
}

func NewMMU() *MMU {
	return &MMU{
		memory: [0x10000]uint8{},
	}
}
