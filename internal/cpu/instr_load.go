package cpu

/*
OPCODE: 0x01
DESCRIPTION: Load immediate 16-bit value into BC
CYCLE: 12
*/
func (cpu *CPU) InstrLD_BC_d16() int {
	w := cpu.FetchWord()
	cpu.Registers.SetBC(w)
	return 12
}

/*
OPCODE: 0x02
DESCRIPTION: Load A into address [BC]
CYCLE: 8
*/
func (cpu *CPU) InstrLD_BC_ad_A() int {
	addr := cpu.Registers.GetBC()
	cpu.Bus.Write(addr, cpu.Registers.A)
	return 8
}

/*
OPCODE: 0x06
DESCRIPTION: Load immediate 8-bits into B
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
func (cpu *CPU) InstrLD_A_BC_ad() int {
	addr := cpu.Registers.GetBC() // 16-bit address
	val := cpu.Bus.Read(addr)     // read 8-bit from memory
	cpu.Registers.A = val         // load into A
	return 8
}

/*
OPCODE: 0x0E
DESCRIPTION: Load immediate 8-bit value into C
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
DESCRIPTION: Store A into address (DE)
CYCLE: 8
*/
func (cpu *CPU) InstrLD_DE_ad_A() int {
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
func (cpu *CPU) InstrLD_A_DE_ad() int {
	addr := cpu.Registers.GetDE()
	b := cpu.Bus.Read(addr)
	cpu.Registers.A = b
	return 8
}

/*
OPCODE: 0x1E
DESCRIPTION:  Load immediate 8-bit value into C
CYCLE: 8
*/
func (cpu *CPU) InstrLD_E_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.E = b
	return 8
}

/*
OPCODE: 0x21
DESCRIPTION: SET HL with 16 bits
CYCLE: 12
*/
func (cpu *CPU) InstrLD_HL_d16() int {
	w := cpu.FetchWord()
	cpu.Registers.SetHL(w)
	return 12
}

/*
OPCODE: 0x31
DESCRIPTION: Load immediate 16-bit value into SP
CYCLE: 12
*/
func (cpu *CPU) InstrLD_SP_d16() int {
	w := cpu.FetchWord()
	cpu.Registers.SP = w
	return 12
}

/*
OPCODE: 0x22
DESCRIPTION: Store A at (HL), then increment HL
CYCLE: 8
*/
func (cpu *CPU) InstrLD_HL_ad_INC_A() int {
	addr := cpu.Registers.GetHL()
	cpu.Bus.Write(addr, cpu.Registers.A)
	cpu.Registers.SetHL(addr + 1)
	return 8
}

/*
OPCODE: 0x32
DESCRIPTION: Store A at (HL), then decrement HL
CYCLE: 8
*/
func (cpu *CPU) InstrLD_HL_ad_DEC_A() int {
	addr := cpu.Registers.GetHL()
	cpu.Bus.Write(addr, cpu.Registers.A)
	cpu.Registers.SetHL(addr - 1)
	return 8
}

/*
OPCODE: 0x3E
DESCRIPTION: SET A with 8 bits
CYCLE: 8
*/
func (cpu *CPU) InstrLD_A_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.A = b
	return 8
}

/*
OPCODE: 0xE2
DESCRIPTION: Store A into (0xFF00 + C)
CYCLE: 8
*/
func (cpu *CPU) InstrLD_C_ad_A() int {
	addr := 0xFF00 + uint16(cpu.Registers.C)
	cpu.Bus.Write(addr, cpu.Registers.A)
	return 8
}

/*
OPCODE: 0xF2
DESCRIPTION: Load A from (0xFF00 + C)
CYCLE: 8
*/
func (cpu *CPU) InstrLD_A_C_Ad() int {
	addr := 0xFF00 + uint16(cpu.Registers.C)
	cpu.Registers.A = cpu.Bus.Read(uint16(addr))
	return 8
}

/*
OPCODE: 0xE0
DESCRIPTION: LDH (a8), A -> Load A into address 0xFF00 + n
CYCLE: 12
*/
func (cpu *CPU) InstrLDH_a8_ad_A() int {
	addr := 0xFF00 + uint16(cpu.FetchByte())
	cpu.Bus.Write(addr, cpu.Registers.A)
	return 12
}

/*
OPCODE: 0xF0
DESCRIPTION: LDH A, (a8) -> Load value at 0xFF00 + (a8) into A
CYCLE: 12
*/
func (cpu *CPU) InstrLDH_A_a8() int {
	addr := 0xFF00 + uint16(cpu.FetchByte())
	value := cpu.Bus.Read(addr)
	cpu.Registers.A = value
	return 12
}

/*
OPCODE: 0x36
DESCRIPTION: laad immediate 8-bits at address [HL]
CYCLE: 12
*/
func (cpu *CPU) InstrLD_HL_ad_d8() int {
	addr := cpu.Registers.GetHL()
	b := cpu.FetchByte()
	cpu.Bus.Write(addr, b)
	return 12
}

/*
OPCODE: 0xEA
DESCRIPTION: laad A into address a16
CYCLE: 16
*/
func (cpu *CPU) InstrLD_a16_ad_A() int {
	addr := cpu.FetchWord()
	cpu.Bus.Write(addr, cpu.Registers.A)
	return 16
}

/*
OPCODE: 0x2A
DESCRIPTION: laad address [HL] into A, increment HL
CYCLE: 8
*/
func (cpu *CPU) InstrLD_A_HL_ad_INC() int {
	addr := cpu.Registers.GetHL()
	b := cpu.Bus.Read(addr)
	cpu.Registers.A = b
	cpu.Registers.SetHL(addr + 1)
	return 8
}

/*
OPCODE: 0x2A
DESCRIPTION: laad address [HL] into A, decrement HL
CYCLE: 8
*/
func (cpu *CPU) InstrLD_A_HL_ad_DEC() int {
	addr := cpu.Registers.GetHL()
	b := cpu.Bus.Read(addr)
	cpu.Registers.A = b
	cpu.Registers.SetHL(addr - 1)
	return 8
}

/*
OPCODE: 0x26
DESCRIPTION: laad immediate 8-bits into H
CYCLE: 8
*/
func (cpu *CPU) InstrLD_H_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.H = b
	return 8
}

/*
OPCODE: 0x2E
DESCRIPTION: llaad immediate 8-bits into L
CYCLE: 8
*/
func (cpu *CPU) InstrLD_L_d8() int {
	b := cpu.FetchByte()
	cpu.Registers.L = b
	return 8
}

/*
OPCODE: 0x78
DESCRIPTION: laad A with B
CYCLE: 4
*/
func (cpu *CPU) InstrLD_A_B() int {
	cpu.Registers.A = cpu.Registers.B
	return 8
}

/*
OPCODE: 0x47
DESCRIPTION: laad B with A
CYCLE: 4
*/
func (cpu *CPU) InstrLD_B_A() int {
	cpu.Registers.B = cpu.Registers.A
	return 8
}

/*
OPCODE: 0x4F
DESCRIPTION: laad C with A
CYCLE: 4
*/
func (cpu *CPU) InstrLD_C_A() int {
	cpu.Registers.C = cpu.Registers.A
	return 8
}

/*
OPCODE: 0x79
DESCRIPTION: laad A with C
CYCLE: 4
*/
func (cpu *CPU) InstrLD_A_C() int {
	cpu.Registers.A = cpu.Registers.C
	return 8
}
