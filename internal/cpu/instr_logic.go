package cpu

func (cpu *CPU) InstructLOGIC_XOR_A() int {
	cpu.Registers.A ^= cpu.Registers.A
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
	return 4
}
