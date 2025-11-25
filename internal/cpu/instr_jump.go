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

func (cpu *CPU) RET() int {
	cpu.Registers.PC = cpu.Pop()
	return 16
}

func (cpu *CPU) RETI() int {
	cpu.Registers.PC = cpu.Pop()
	cpu.IME = true
	return 16
}

func (cpu *CPU) RET_NZ() int {
	if !cpu.Registers.GetFlag(FlagZ) {
		cpu.Registers.PC = cpu.Pop()
		return 20
	} else {
		return 8
	}
}
func (cpu *CPU) RET_NC() int {
	if !cpu.Registers.GetFlag(FlagC) {
		cpu.Registers.PC = cpu.Pop()
		return 20
	} else {
		return 8
	}
}

func (cpu *CPU) RET_Z() int {
	if cpu.Registers.GetFlag(FlagZ) {
		cpu.Registers.PC = cpu.Pop()
		return 20
	} else {
		return 8
	}
}

func (cpu *CPU) RET_C() int {
	if cpu.Registers.GetFlag(FlagC) {
		cpu.Registers.PC = cpu.Pop()
		return 20
	} else {
		return 8
	}
}

func (cpu *CPU) CALL() int {
	w := cpu.FetchWord()
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = w
	return 24
}

func (cpu *CPU) CALL_C() int {
	w := cpu.FetchWord()
	if cpu.Registers.GetFlag(FlagC) {
		cpu.Push(cpu.Registers.PC)
		cpu.Registers.PC = w
		return 24
	} else {
		return 12
	}
}

func (cpu *CPU) CALL_Z() int {
	w := cpu.FetchWord()
	if cpu.Registers.GetFlag(FlagZ) {
		cpu.Push(cpu.Registers.PC)
		cpu.Registers.PC = w
		return 24
	} else {
		return 12
	}
}

func (cpu *CPU) CALL_NC() int {
	w := cpu.FetchWord()
	if !cpu.Registers.GetFlag(FlagC) {
		cpu.Push(cpu.Registers.PC)
		cpu.Registers.PC = w
		return 24
	} else {
		return 12
	}
}

func (cpu *CPU) CALL_NZ() int {
	w := cpu.FetchWord()
	if !cpu.Registers.GetFlag(FlagZ) {
		cpu.Push(cpu.Registers.PC)
		cpu.Registers.PC = w
		return 24
	} else {
		return 12
	}
}

func (cpu *CPU) RST_00H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0000
	return 16
}
func (cpu *CPU) RST_10H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0010
	return 16
}
func (cpu *CPU) RST_20H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0020
	return 16
}
func (cpu *CPU) RST_30H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0030
	return 16
}
func (cpu *CPU) RST_08H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0008
	return 16
}
func (cpu *CPU) RST_18H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0018
	return 16
}
func (cpu *CPU) RST_28H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0028
	return 16
}
func (cpu *CPU) RST_38H() int {
	cpu.Push(cpu.Registers.PC)
	cpu.Registers.PC = 0x0038
	return 16
}
