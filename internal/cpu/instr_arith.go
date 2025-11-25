package cpu

func (cpu *CPU) Add_A(value uint8) {
	isHalfCarry := (cpu.Registers.A&0x0f)+(value&0x0f) > 0x0f
	sum16 := uint16(cpu.Registers.A) + uint16(value)
	isCarry := sum16 > 0xFF
	cpu.Registers.A = uint8(sum16)
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0) // Set Z if result is 0
	cpu.Registers.SetFlag(FlagN, false)                // Reset N (this is add, not sub)
	cpu.Registers.SetFlag(FlagH, isHalfCarry)
	cpu.Registers.SetFlag(FlagC, isCarry)
}

func (cpu *CPU) Sub_A(value uint8) {
	isHalfCarry := (cpu.Registers.A & 0x0f) < (value & 0x0f)
	isCarry := cpu.Registers.A < value
	cpu.Registers.A -= value
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, isHalfCarry)
	cpu.Registers.SetFlag(FlagC, isCarry)
}

func (cpu *CPU) And_A(value uint8) {
	cpu.Registers.A &= value
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)

	/*
	 * Game Boy CPU (Sharp LR35902),
	 * the AND instruction's internal circuitry biases the Half-Carry line high.
	 * It is a famous quirk of this specific CPU.
	 */
	cpu.Registers.SetFlag(FlagH, true)
	cpu.Registers.SetFlag(FlagC, false)
}

func (cpu *CPU) Or_A(value uint8) {
	cpu.Registers.A |= value
	cpu.Registers.SetFlag(FlagZ, cpu.Registers.A == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)
}

func (cpu *CPU) InstrARITH_Dec_B() int {
	cpu.Registers.A -= 1
	return 4
}

func (cpu *CPU) InstrARITH_Dec_C() int {
	cpu.Registers.C -= 1
	return 4
}
