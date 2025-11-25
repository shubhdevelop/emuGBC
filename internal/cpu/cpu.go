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
	if cpu.LogMode == "hex" {
		fmt.Printf("PC:0x%04X AF:0x%04X BC:0x%04X DE:0x%04X HL:0x%04X SP:0x%04X\n",
			cpu.Registers.PC,
			cpu.Registers.GetAF(),
			cpu.Registers.GetBC(),
			cpu.Registers.GetDE(),
			cpu.Registers.GetHL(),
			cpu.Registers.SP,
		)
	} else if cpu.LogMode == "bin" {
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

func (cpu *CPU) Push() {}
func (cpu *CPU) Pop()  {}

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
		fmt.Printf("Current Opcode: 0x%02X\n", cb)
		fn := cbOpcodes[cb].Fn
		if fn == nil {
			panic(fmt.Sprintf("Undefined CB opcode 0x%02X at PC: 0x%04X", cb, cpu.Registers.PC-1))
		}
		return fn(cpu)
	}

	fmt.Printf("Current Opcode: 0x%02X\n", opcode)
	entry := mainOpcodes[opcode]
	if entry.Fn == nil {
		panic(fmt.Sprintf("Undefined opcode 0x%02X at PC: 0x%04X", opcode, cpu.Registers.PC))
	}
	return entry.Fn(cpu)
}
