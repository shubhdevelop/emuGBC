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

func (m *MMU) Read(addr uint16) uint8 {
	return m.memory[addr]
}

func (m *MMU) Write(addr uint16, val uint8) {
	// ROM PROTECTION (0x0000 - 0x7FFF)
	// If the CPU tries to write here, we ignore it (for now).
	// Later, this is how "Memory Bank Controllers" (MBC) switch banks.
	if addr < 0x8000 {
		return
	}

	// ECHO RAM (0xE000 - 0xFDFF)
	// This area mirrors WRAM (0xC000 - 0xDFFF).
	// Nintendo says "Don't use this", but developers did anyway.
	if addr >= 0xE000 && addr < 0xFE00 {
		m.memory[addr] = val
		m.memory[addr-0x2000] = val // Write to the mirror too
		return
	}

	// RESTRICTED (0xFEA0 - 0xFEFF)
	// This area is unusable on real hardware.
	if addr >= 0xFEA0 && addr < 0xFEFF {
		return
	}

	// Normal Write
	m.memory[addr] = val
}

func (m *MMU) ReadWord(addr uint16) uint16 {
	low := uint16(m.Read(addr))
	high := uint16(m.Read(addr + 1))

	return (high << 8) | low
}

func (m *MMU) WriteWord(addr uint16, val uint16) {
	low := uint8(val & 0x00FF)
	high := uint8((val & 0xFF00) >> 8)

	m.Write(addr, low)
	m.Write(addr+1, high)
}
