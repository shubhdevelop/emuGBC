package cpu

/*
OPCODE: 0x01
DESCRIPTION: Load 16 bits into BC
CYCLE: 12
*/
func (cpu *CPU) InstrLD_BC_d16() int {
	w := cpu.FetchWord()
	cpu.Registers.SetBC(w)
	return 12
}

/*
OPCODE: 0x02
DESCRIPTION: Load A into BC
CYCLE: 8
*/
func (cpu *CPU) InstrLD_BC_A() int {
	addr := cpu.Registers.GetBC()
	cpu.Bus.Write(addr, cpu.Registers.A)
	return 8
}

/*
OPCODE: 0x06
DESCRIPTION: Load 8 bits into B
CYCLE: 8
*/
func (cpu *CPU) InstrLD_B_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.B = b
	return 8
}

/*
OPCODE: 0x08
DESCRIPTION: store the 16-bit SP value into two bytes of memory at address a16
CYCLE: 20
*/
func (cpu *CPU) InstrLD_a16_SP() int {
	addr := cpu.FetchWord() // read 16-bit address (little-endian)
	sp := cpu.Registers.SP  // 16-bit stack pointer
	cpu.Bus.WriteWord(addr, sp)
	return 20
}

/*
OPCODE: 0x0A
DESCRIPTION: SET A with [BC]
CYCLE: 8
*/
func (cpu *CPU) InstrLD_A_BC() int {
	addr := cpu.Registers.GetBC() // 16-bit address
	val := cpu.Bus.Read(addr)     // read 8-bit from memory
	cpu.Registers.A = val         // load into A
	return 8
}

/*
OPCODE: 0x0E
DESCRIPTION: SET C with 8 bits
CYCLE: 8
*/
func (cpu *CPU) InstrLD_C_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.C = b
	return 8
}

/*
OPCODE: 0x11
DESCRIPTION: SET DE with 16 bits
CYCLE: 12
*/
func (cpu *CPU) InstrLD_DE_d16() int {
	w := cpu.FetchWord()
	cpu.Registers.SetDE(w)
	return 12
}

/*
OPCODE: 0x12
DESCRIPTION: SET [DE] with A
CYCLE: 8
*/
func (cpu *CPU) InstrLD_DE_A() int {
	addr := cpu.Registers.GetDE()
	val := cpu.Registers.A
	cpu.Bus.Write(addr, val)
	return 8
}

/*
OPCODE: 0x16
DESCRIPTION: SET D with 8 bits
CYCLE: 8
*/
func (cpu *CPU) InstrLD_D_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.D = b
	return 8
}

/*
OPCODE: 0x1A
DESCRIPTION: SET A with [DE]
CYCLE: 8
*/
func (cpu *CPU) InstrLD_A_DE() int {
	addr := cpu.Registers.GetDE()
	b := cpu.Bus.Read(addr)
	cpu.Registers.A = b
	return 8
}

/*
OPCODE: 0x1E
DESCRIPTION: SET E with 8 bits
CYCLE: 8
*/
func (cpu *CPU) InstrLD_E_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.E = b
	return 8
}
