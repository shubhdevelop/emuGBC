package cpu

import (
	"fmt"

	"github.com/shubhdevelop/emuGBC/internal/mmu"
)

type Registers struct {
	// 8-bit registers
	A, F uint8
	B, C uint8
	D, E uint8
	H, L uint8

	// 16-bit registers
	SP uint16 // Stack Pointer
	PC uint16 // Program Counter
}

/*
 * GBC CPU doesn't use the V flag for the overflow
 */
const (
	FlagZ = 1 << 7 // Zero
	FlagN = 1 << 6 // Subtract
	FlagH = 1 << 5 // Half Carry
	FlagC = 1 << 4 // Carry
)

type CPU struct {
	Registers Registers
	Bus       *mmu.MMU
}

func NewCPU(bus *mmu.MMU) *CPU {
	return &CPU{
		Bus: bus,
	}
}

func (r *Registers) GetFlag(flag uint8) bool {
	return (r.F & flag) != 0
}

func (r *Registers) SetFlag(flag uint8, value bool) {
	if value {
		r.F |= flag
	} else {
		r.F &^= flag
	}
}

func (r *Registers) GetBC() uint16 {
	return (uint16(r.B) << 8) | uint16(r.C)
}

func (r *Registers) GetDE() uint16 {
	return (uint16(r.D) << 8) | uint16(r.E)
}

func (r *Registers) GetHL() uint16 {
	return (uint16(r.H) << 8) | uint16(r.L)
}

func (r *Registers) GetAF() uint16 {
	return (uint16(r.A) << 8) | uint16(r.F)
}

func (r *Registers) SetBC(value uint16) {
	r.B = uint8((value & 0xFF00) >> 8) // Grab top 8 bits
	r.C = uint8(value & 0x00FF)        // Grab bottom 8 bits
}

func (r *Registers) SetDE(value uint16) {
	r.D = uint8((value & 0xFF00) >> 8)
	r.E = uint8(value & 0x00FF)
}

func (r *Registers) SetHL(value uint16) {
	r.H = uint8((value & 0xFF00) >> 8)
	r.L = uint8(value & 0x00FF)
}

func (r *Registers) SetAF(value uint16) {
	r.A = uint8((value & 0xFF00) >> 8)
	r.F = uint8(value&0x00FF) & 0xF0
}

// Step executes the next instruction
func (cpu *CPU) Step() {
	opcode := cpu.Bus.Read(cpu.Registers.PC)
	cpu.Registers.PC++

	switch opcode {
	case 0x00:
		return

	case 0x04:
		cpu.Registers.B++
		return

	case 0x05:
		cpu.Registers.B--
		return

	case 0x06:
		// LD B, n (Load 8-bit immediate value into B)
		value := cpu.Bus.Read(cpu.Registers.PC)
		cpu.Registers.B = value

		/*
		 * We need to read the NEXT byte.
		 * This is because the opcode is 2 bytes long.
		 * We consumed a second byte, so increment PC again!
		 */
		cpu.Registers.PC++ //
		return

	default:
		panic(fmt.Sprintf("Unknown Opcode: 0x%X at PC: 0x%X", opcode, cpu.Registers.PC-1))
	}
}
