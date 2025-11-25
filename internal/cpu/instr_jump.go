package cpu

func (cpu *CPU) JumpAbsolute() int {
	target := cpu.Bus.ReadWord(cpu.Registers.PC)
	cpu.Registers.PC = target
	return 16
}

func (cpu *CPU) JumpAbsoluteNotZero() int {
	target := cpu.Bus.ReadWord(cpu.Registers.PC)
	if !cpu.Registers.GetFlag(FlagZ) {
		cpu.Registers.PC = target
		return 16
	} else {
		return 12
	}
}

func (cpu *CPU) JumpAbsoluteZero() int {
	target := cpu.Bus.ReadWord(cpu.Registers.PC)
	if cpu.Registers.GetFlag(FlagZ) {
		cpu.Registers.PC = target
		return 16
	} else {
		return 12
	}
}

func (cpu *CPU) JumpAbsoluteNotCarry() int {
	target := cpu.Bus.ReadWord(cpu.Registers.PC)
	if !cpu.Registers.GetFlag(FlagC) {
		cpu.Registers.PC = target
		return 16
	} else {
		return 12
	}
}

func (cpu *CPU) JumpAbsoluteCarry() int {
	target := cpu.Bus.ReadWord(cpu.Registers.PC)
	if cpu.Registers.GetFlag(FlagC) {
		cpu.Registers.PC = target
		return 16
	} else {
		return 12
	}
}

func (cpu *CPU) JumpAbsoluteHL_ad() int {
	target := cpu.Registers.GetHL()
	cpu.Registers.PC = target
	return 4
}

func (cpu *CPU) JumpRelative() int {

	offset := cpu.Bus.Read(cpu.Registers.PC)
	cpu.Registers.PC++
	currentPC := int32(cpu.Registers.PC)
	jump := int32(int8(offset)) // 0xFE becomes -2

	cpu.Registers.PC = uint16(currentPC + jump)
	return 12
}

func (cpu *CPU) JumpRelativeNotZero() int {
	offset := int8(cpu.Bus.Read(cpu.Registers.PC))
	cpu.Registers.PC++

	// 2. Check the condition
	if !cpu.Registers.GetFlag(FlagZ) {
		cpu.Registers.PC = uint16(int16(offset))
		return 12
	} else {
		return 8
	}
}

func (cpu *CPU) JumpRelativeNotCarry() int {
	offset := int8(cpu.Bus.Read(cpu.Registers.PC))
	cpu.Registers.PC++

	// 2. Check the condition
	if !cpu.Registers.GetFlag(FlagC) {
		cpu.Registers.PC = uint16(int16(offset))
		return 12
	} else {
		return 8
	}
}

func (cpu *CPU) JumpRelativeZero() int {
	offset := int8(cpu.Bus.Read(cpu.Registers.PC))
	cpu.Registers.PC++

	// 2. Check the condition
	if cpu.Registers.GetFlag(FlagZ) {
		cpu.Registers.PC = uint16(int16(offset))
		return 12
	} else {
		return 8
	}
}

func (cpu *CPU) JumpRelativeCarry() int {
	offset := int8(cpu.Bus.Read(cpu.Registers.PC))
	cpu.Registers.PC++

	// 2. Check the condition
	if cpu.Registers.GetFlag(FlagC) {
		cpu.Registers.PC = uint16(int16(offset))
		return 12
	} else {
		return 8
	}
}
