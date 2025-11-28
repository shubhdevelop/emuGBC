package cpu

func (cpu *CPU) RLCA() int {
	bit7 := (cpu.Registers.A >> 7) & 0x01

	cpu.Registers.A = (cpu.Registers.A << 1) | bit7

	cpu.Registers.SetFlag(FlagZ, false) // Quirk: Always 0 on DMG!
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, bit7 == 1)

	return 4
}

func (cpu *CPU) RRCA() int {
	bit0 := cpu.Registers.A & 0x01

	cpu.Registers.A = (cpu.Registers.A >> 1) | (bit0 << 7)

	cpu.Registers.SetFlag(FlagZ, false) // Quirk: Always 0 on DMG!
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, bit0 == 1)

	return 4
}

func (cpu *CPU) RLA() int {
	oldCarry := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		oldCarry = 1
	}

	newCarry := (cpu.Registers.A >> 7) & 0x01

	cpu.Registers.A = (cpu.Registers.A << 1) | oldCarry

	cpu.Registers.SetFlag(FlagZ, false) // Quirk: Always 0 on DMG
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, newCarry == 1)

	return 4
}

func (cpu *CPU) RRA() int {
	oldCarry := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		oldCarry = 1
	}

	newCarry := cpu.Registers.A & 0x01

	cpu.Registers.A = (cpu.Registers.A >> 1) | (oldCarry << 7)

	cpu.Registers.SetFlag(FlagZ, false) // Quirk: Always 0 on DMG
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, newCarry == 1)

	return 4
}
