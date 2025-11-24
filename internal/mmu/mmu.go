package mmu

import "github.com/shubhdevelop/emuGBC/internal/ppu"

type MMU struct {
	// Represents the entire 64KB address space
	memory [0x10000]uint8

	// GBC VRAM: 2 Banks, each 8KB (0x2000 bytes)
	// Bank 0: Standard Data
	// Bank 1: Attributes & Extra Data
	vram [2][0x2000]uint8
	PPU  *ppu.PPU

	vbk uint8 // VRAM Bank Register (0xFF4F)
}

func NewMMU() *MMU {
	return &MMU{
		memory: [0x10000]uint8{},
	}
}

func (m *MMU) Read(addr uint16) uint8 {
	// 1. Intercept VRAM Range (0x8000 - 0x9FFF)
	if addr >= 0x8000 && addr <= 0x9FFF {
		// Calculate offset relative to start of VRAM
		offset := addr - 0x8000

		// Which bank is currently active?
		bankIndex := m.vbk & 0x01 // Only bit 0 matters

		return m.vram[bankIndex][offset]
	}

	// 2. Intercept VBK Register (0xFF4F)
	if addr == 0xFF4F {
		return m.vbk | 0xFE // bits 1-7 are always 1 on reads
	}

	// Default Memory Read
	return m.memory[addr]
}

func (m *MMU) Write(addr uint16, val uint8) {

	// 1. ROM PROTECTION (0x0000 - 0x7FFF)
	// Read-Only. Ignore writes.
	if addr < 0x8000 {
		return
	}

	// 2. VRAM (0x8000 - 0x9FFF) - Banked
	// Routes to the specific VRAM bank and returns.
	// We DO NOT write to m.memory[] here.
	if addr >= 0x8000 && addr <= 0x9FFF {
		offset := addr - 0x8000
		bankIndex := m.vbk & 0x01
		m.vram[bankIndex][offset] = val
		return
	}

	// 3. WRAM & ECHO RAM (Mirroring Logic)
	// Corrected: We must write to BOTH locations regardless of which one was targeted.

	// Case A: Writing to WRAM (0xC000 - 0xDFFF)
	if addr >= 0xC000 && addr <= 0xDFFF {
		m.memory[addr] = val
		m.memory[addr+0x2000] = val // Mirror to Echo
		return
	}

	// Case B: Writing to Echo RAM (0xE000 - 0xFDFF)
	if addr >= 0xE000 && addr <= 0xFDFF {
		m.memory[addr] = val
		m.memory[addr-0x2000] = val // Mirror to WRAM
		return
	}

	// 4. RESTRICTED (0xFEA0 - 0xFEFF)
	// Unusable memory. Ignore.
	if addr >= 0xFEA0 && addr <= 0xFEFF {
		return
	}

	// 5. IO REGISTERS (0xFF00 - 0xFF7F)
	// We check for specific hardware hooks here.
	// Note: We DO NOT return. We let it fall through to step 6.
	// Why? So that m.memory[addr] gets updated too. This makes implementing
	// Read() easier (you can just read from memory[] for simple registers).

	if addr == 0xFF40 {
		m.PPU.WriteLCDC(val)
	}

	if addr == 0xFF4F {
		m.vbk = val & 0x01
	}

	// 6. GENERIC WRITE
	// Handles OAM (0xFE00-0xFE9F), HRAM (0xFF80-0xFFFE), and updates the IO Cache.
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

func (m *MMU) LoadROM(data []byte) {
	copy(m.memory[:], data)
}
