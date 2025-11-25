package cpu

func (cpu *CPU) cp(value byte) {
	a := cpu.Registers.A
	result := a - value

	// Z flag: set if result == 0
	cpu.Registers.SetFlag(FlagZ, (result == 0))

	// N flag: always set (comparison is subtraction)
	cpu.Registers.SetFlag(FlagZ, true)

	// H flag: half-carry/borrow from bit 4
	cpu.Registers.SetFlag(FlagH, halfCarrySub(a, value))

	// C flag: full borrow
	cpu.Registers.SetFlag(FlagH, carrySub(a, value))
}

func (cpu *CPU) InstrLOGIC_XOR_A() int {
	cpu.Registers.A ^= cpu.Registers.A
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
	return 4
}

func (cpu *CPU) InstrLOGIC_CP_d8() int {
	val := cpu.FetchByte()
	cpu.cp(val)
	return 8
}
