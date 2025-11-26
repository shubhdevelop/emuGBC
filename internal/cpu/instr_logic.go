package cpu

func (cpu *CPU) AND(val uint8) int {
	cpu.Registers.A &= val
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, true)
	cpu.Registers.SetFlag(FlagC, false)
	return 4
}

func (cpu *CPU) XOR(val uint8) int {
	cpu.Registers.A ^= val
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
	return 4
}

func (cpu *CPU) OR(val uint8) int {
	cpu.Registers.A |= val
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
	return 4
}

func (cpu *CPU) CP(value byte) int {
	a := cpu.Registers.A
	result := a - value
	cpu.Registers.SetFlag(FlagZ, (result == 0))
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, halfCarrySub(a, value))
	cpu.Registers.SetFlag(FlagH, carrySub(a, value))
	return 4
}

func (cpu *CPU) InstrLOGIC_CP_d8() int {
	return cpu.CP(cpu.FetchByte()) + 4
}

func (cpu *CPU) InstrLOGIC_CP_B() int {
	return cpu.CP(cpu.Registers.B)
}
func (cpu *CPU) InstrLOGIC_CP_C() int {
	return cpu.CP(cpu.Registers.C)
}

func (cpu *CPU) InstrLOGIC_CP_D() int {
	return cpu.CP(cpu.Registers.D)
}
func (cpu *CPU) InstrLOGIC_CP_E() int {
	return cpu.CP(cpu.Registers.E)
}
func (cpu *CPU) InstrLOGIC_CP_H() int {
	return cpu.CP(cpu.Registers.H)
}
func (cpu *CPU) InstrLOGIC_CP_L() int {
	return cpu.CP(cpu.Registers.L)
}
func (cpu *CPU) InstrLOGIC_CP_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	return cpu.CP(val) + 4
}
func (cpu *CPU) InstrLOGIC_CP_A() int {
	return cpu.CP(cpu.Registers.A)
}

func (cpu *CPU) InstrLOGIC_OR_B() int {
	return cpu.OR(cpu.Registers.B)
}
func (cpu *CPU) InstrLOGIC_OR_C() int {
	return cpu.OR(cpu.Registers.C)
}
func (cpu *CPU) InstrLOGIC_OR_D() int {
	return cpu.OR(cpu.Registers.D)
}
func (cpu *CPU) InstrLOGIC_OR_E() int {
	return cpu.OR(cpu.Registers.E)
}
func (cpu *CPU) InstrLOGIC_OR_H() int {
	return cpu.OR(cpu.Registers.H)
}
func (cpu *CPU) InstrLOGIC_OR_L() int {
	return cpu.OR(cpu.Registers.L)
}
func (cpu *CPU) InstrLOGIC_OR_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	return cpu.OR(val) + 4
}
func (cpu *CPU) InstrLOGIC_OR_A() int {
	return cpu.OR(cpu.Registers.A)
}
func (cpu *CPU) InstrLOGIC_OR_d8() int {
	val := cpu.FetchByte()
	return cpu.OR(val) + 4
}

func (cpu *CPU) InstrLOGIC_AND_B() int {
	return cpu.AND(cpu.Registers.B)
}
func (cpu *CPU) InstrLOGIC_AND_C() int {
	return cpu.AND(cpu.Registers.C)
}
func (cpu *CPU) InstrLOGIC_AND_D() int {
	return cpu.AND(cpu.Registers.D)
}
func (cpu *CPU) InstrLOGIC_AND_E() int {
	return cpu.AND(cpu.Registers.E)
}
func (cpu *CPU) InstrLOGIC_AND_H() int {
	return cpu.AND(cpu.Registers.H)
}
func (cpu *CPU) InstrLOGIC_AND_L() int {
	return cpu.AND(cpu.Registers.L)
}
func (cpu *CPU) InstrLOGIC_AND_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	return cpu.AND(val) + 4
}
func (cpu *CPU) InstrLOGIC_AND_A() int {
	return cpu.AND(cpu.Registers.A)
}
func (cpu *CPU) InstrLOGIC_AND_d8() int {
	val := cpu.FetchByte()
	return cpu.AND(val) + 4
}

func (cpu *CPU) InstrLOGIC_XOR_B() int {
	return cpu.XOR(cpu.Registers.B)
}
func (cpu *CPU) InstrLOGIC_XOR_C() int {
	return cpu.XOR(cpu.Registers.C)
}

func (cpu *CPU) InstrLOGIC_XOR_D() int {
	return cpu.XOR(cpu.Registers.D)
}
func (cpu *CPU) InstrLOGIC_XOR_E() int {
	return cpu.XOR(cpu.Registers.E)
}
func (cpu *CPU) InstrLOGIC_XOR_H() int {
	return cpu.XOR(cpu.Registers.H)
}
func (cpu *CPU) InstrLOGIC_XOR_L() int {
	return cpu.XOR(cpu.Registers.L)
}
func (cpu *CPU) InstrLOGIC_XOR_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)

	return cpu.XOR(val) + 4
}
func (cpu *CPU) InstrLOGIC_XOR_A() int {
	return cpu.XOR(cpu.Registers.A)
}
func (cpu *CPU) InstrLOGIC_XOR_d8() int {
	val := cpu.FetchByte()
	return cpu.XOR(val) + 4
}
func (cpu *CPU) InstrLOGIC_CPL() int {
	cpu.Registers.A = ^cpu.Registers.A

	// Flags: N=1, H=1. Z and C are preserved.
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, true)
	return 4
}
