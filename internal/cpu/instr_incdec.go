package cpu

func (cpu *CPU) INC_B() int {
	v := cpu.Registers.B
	cpu.Registers.B = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}
func (cpu *CPU) INC_D() int {
	v := cpu.Registers.D
	cpu.Registers.D = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}
func (cpu *CPU) INC_H() int {
	v := cpu.Registers.H
	cpu.Registers.H = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}

/*
OPCODE: 0x34
DESCRIPTION: INC (HL) -> Increment the BYTE pointed to by HL
CYCLES: 12
*/
func (cpu *CPU) INC_HL_ad() int {
	addr := cpu.Registers.GetHL()
	v := cpu.Bus.Read(addr)
	result := v + 1
	cpu.Bus.Write(addr, result)
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 12
}

/*
OPCODE: 0x34
DESCRIPTION: DEV (HL) -> Decrement the BYTE pointed to by HL
CYCLES: 12
*/
func (cpu *CPU) DEC_HL_ad() int {
	addr := cpu.Registers.GetHL()
	v := cpu.Bus.Read(addr)
	result := v - 1
	cpu.Bus.Write(addr, result)
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 12
}

func (cpu *CPU) INC_C() int {
	v := cpu.Registers.C
	cpu.Registers.C = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}
func (cpu *CPU) INC_E() int {
	v := cpu.Registers.E
	cpu.Registers.E = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}
func (cpu *CPU) INC_L() int {
	v := cpu.Registers.L
	cpu.Registers.L = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}
func (cpu *CPU) INC_A() int {
	v := cpu.Registers.A
	cpu.Registers.A = v + 1
	cpu.Registers.SetFlag(FlagZ, v+1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryInc(v))
	cpu.Registers.SetFlag(FlagN, false)
	return 4
}

func (cpu *CPU) DEC_B() int {
	v := cpu.Registers.B
	cpu.Registers.B = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}
func (cpu *CPU) DEC_D() int {
	v := cpu.Registers.D
	cpu.Registers.D = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}
func (cpu *CPU) DEC_H() int {
	v := cpu.Registers.H
	cpu.Registers.H = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}

func (cpu *CPU) DEC_C() int {
	v := cpu.Registers.C
	cpu.Registers.C = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}
func (cpu *CPU) DEC_E() int {
	v := cpu.Registers.E
	cpu.Registers.E = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}
func (cpu *CPU) DEC_L() int {
	v := cpu.Registers.L
	cpu.Registers.L = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}
func (cpu *CPU) DEC_A() int {
	v := cpu.Registers.A
	cpu.Registers.A = v - 1
	cpu.Registers.SetFlag(FlagZ, v-1 == 0)
	cpu.Registers.SetFlag(FlagH, halfCarryDec(v))
	cpu.Registers.SetFlag(FlagN, true)
	return 4
}

func (cpu *CPU) INC_BC() int {
	v := cpu.Registers.GetBC()
	cpu.Registers.SetBC(v + 1)
	return 8
}
func (cpu *CPU) INC_DE() int {
	v := cpu.Registers.GetDE()
	cpu.Registers.SetDE(v + 1)
	return 8
}
func (cpu *CPU) INC_HL() int {
	v := cpu.Registers.GetHL()
	cpu.Registers.SetHL(v + 1)
	return 8
}
func (cpu *CPU) INC_SP() int {
	v := cpu.Registers.SP
	cpu.Registers.SP = v + 1
	return 8
}

func (cpu *CPU) DEC_BC() int {
	v := cpu.Registers.GetBC()
	cpu.Registers.SetBC(v - 1)
	return 8
}
func (cpu *CPU) DEC_DE() int {
	v := cpu.Registers.GetDE()
	cpu.Registers.SetDE(v - 1)
	return 8
}
func (cpu *CPU) DEC_HL() int {
	v := cpu.Registers.GetHL()
	cpu.Registers.SetHL(v - 1)
	return 8
}
func (cpu *CPU) DEC_SP() int {
	v := cpu.Registers.SP
	cpu.Registers.SP = v - 1
	return 8
}
