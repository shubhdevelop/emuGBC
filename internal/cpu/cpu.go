package cpu

import (
	"fmt"

	"github.com/shubhdevelop/emuGBC/internal/mmu"
)

type CPU struct {
	Registers Registers
	Bus       *mmu.MMU

	// internal state
	isHalted bool
	IME      bool

	LogMode string
}

type Opcode struct {
	Fn       func(*CPU) int
	Mnemonic string
}

func NewCPU(bus *mmu.MMU) *CPU {
	return &CPU{
		Bus:     bus,
		LogMode: "hex",
	}
}

func (cpu *CPU) FetchByte() uint8 {
	b := cpu.Bus.Read(cpu.Registers.PC)
	cpu.Registers.PC++
	return b
}

func (cpu *CPU) FetchWord() uint16 {
	w := cpu.Bus.ReadWord(cpu.Registers.PC)
	cpu.Registers.PC += 2
	return w
}

func (cpu *CPU) Log() {
	// Print current state BEFORE executing the instruction
	switch cpu.LogMode {
	case "hex":
		fmt.Printf("PC:0x%04X AF:0x%04X BC:0x%04X DE:0x%04X HL:0x%04X SP:0x%04X\n",
			cpu.Registers.PC,
			cpu.Registers.GetAF(),
			cpu.Registers.GetBC(),
			cpu.Registers.GetDE(),
			cpu.Registers.GetHL(),
			cpu.Registers.SP,
		)
	case "bin":
		fmt.Printf("PC:%b AF:%b BC:%b DE:%b HL:%b SP:%b\n",
			cpu.Registers.PC,
			cpu.Registers.GetAF(),
			cpu.Registers.GetBC(),
			cpu.Registers.GetDE(),
			cpu.Registers.GetHL(),
			cpu.Registers.SP,
		)
	}
}

// Step executes the next instruction
func (cpu *CPU) Step() int {
	cpu.Log()

	if cpu.isHalted {
		return 4
	}

	opcode := cpu.FetchByte() // advances PC by 1

	// CB prefix
	if opcode == 0xCB {
		cb := cpu.FetchByte()
		entry := cbOpcodes[cb]
		fmt.Printf("Current CB Opcode: 0x%02X Mnemonic: %s \n", opcode, entry.Mnemonic)
		if entry.Fn == nil {
			panic(fmt.Sprintf("Undefined CB opcode 0x%02X at PC: 0x%04X", cb, cpu.Registers.PC-1))
		}
		return entry.Fn(cpu)
	}

	entry := mainOpcodes[opcode]
	fmt.Printf("Current Opcode: 0x%02X Mnemonic: %s \n", opcode, entry.Mnemonic)
	if entry.Fn == nil {
		panic(fmt.Sprintf("Undefined opcode 0x%02X at PC: 0x%04X", opcode, cpu.Registers.PC))
	}
	return entry.Fn(cpu)
}

func (cpu *CPU) Push(val uint16) {
	high := uint8(val >> 8)
	low := uint8(val & 0xFF)

	cpu.Registers.SP--
	cpu.Bus.Write(cpu.Registers.SP, high)

	cpu.Registers.SP--
	cpu.Bus.Write(cpu.Registers.SP, low)
}

func (cpu *CPU) Pop() uint16 {
	low := cpu.Bus.Read(cpu.Registers.SP)
	cpu.Registers.SP++

	high := cpu.Bus.Read(cpu.Registers.SP)
	cpu.Registers.SP++

	return (uint16(high) << 8) | uint16(low)
}
