package cpu

func (cpu *CPU) cp(value byte) {
	a := cpu.Registers.A
	result := a - value

	cpu.Registers.SetFlag(FlagZ, (result == 0))

	cpu.Registers.SetFlag(FlagZ, true)

	cpu.Registers.SetFlag(FlagH, halfCarrySub(a, value))

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
func (cpu *CPU) InstrLOGIC_CP_B() int {
	val := cpu.Registers.B
	cpu.cp(val)
	return 4
}
func (cpu *CPU) InstrLOGIC_CP_C() int {
	val := cpu.Registers.C
	cpu.cp(val)
	return 4
}

func (cpu *CPU) InstrLOGIC_CP_D() int {
	val := cpu.Registers.D
	cpu.cp(val)
	return 4
}
func (cpu *CPU) InstrLOGIC_CP_E() int {
	val := cpu.Registers.E
	cpu.cp(val)
	return 4
}
func (cpu *CPU) InstrLOGIC_CP_H() int {
	val := cpu.Registers.H
	cpu.cp(val)
	return 4
}
func (cpu *CPU) InstrLOGIC_CP_L() int {
	val := cpu.Registers.L
	cpu.cp(val)
	return 4
}
func (cpu *CPU) InstrLOGIC_CP_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	cpu.cp(val)
	return 8
}
func (cpu *CPU) InstrLOGIC_CP_A() int {
	val := cpu.Registers.A
	cpu.cp(val)
	return 4
}
