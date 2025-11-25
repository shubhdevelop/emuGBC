package cpu

import (
	"fmt"

	"github.com/shubhdevelop/emuGBC/internal/mmu"
)

type CPU struct {
	Registers Registers
	Bus       *mmu.MMU
}

type Opcode struct {
	Fn       func(*CPU) int
	Mnemonic string
}

func NewCPU(bus *mmu.MMU) *CPU {
	return &CPU{
		Bus: bus,
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
	fmt.Printf("PC:%04X AF:%04X BC:%04X DE:%04X HL:%04X SP:%04X\n",
		cpu.Registers.PC,
		cpu.Registers.GetAF(),
		cpu.Registers.GetBC(),
		cpu.Registers.GetDE(),
		cpu.Registers.GetHL(),
		cpu.Registers.SP,
	)
}

func (cpu *CPU) Push() {}
func (cpu *CPU) Pop()  {}

// Step executes the next instruction
func (cpu *CPU) Step() int {
	cpu.Log()
	pc := cpu.Registers.PC
	opcode := cpu.FetchByte() // advances PC by 1

	// CB prefix
	if opcode == 0xCB {
		cb := cpu.FetchByte() // advances PC by 1
		entry := cbOpcodes[cb]
		if entry.Fn == nil {
			panic(fmt.Sprintf("Undefined CB opcode CB %02X at PC: %04X", cb, pc))
		}

		return entry.Fn(cpu)
	}

	entry := mainOpcodes[opcode]
	if entry.Fn == nil {
		panic(fmt.Sprintf("Undefined opcode %02X at PC: %04X", opcode, pc))
	}
	return entry.Fn(cpu)
}
