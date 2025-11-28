package cpu

func (cpu *CPU) Add(value uint8) int {
	a := cpu.Registers.A
	result := a + value

	// Flags
	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, halfCarryAdd(a, value))
	cpu.Registers.SetFlag(FlagC, carryAdd(a, value))

	cpu.Registers.A = result
	return 4
}

func (cpu *CPU) add_HL_rr(val uint16) int {
	hl := cpu.Registers.GetHL()

	sum := int32(hl) + int32(val)

	cpu.Registers.SetFlag(FlagN, false)

	isHalfCarry := (hl&0x0FFF)+(val&0x0FFF) > 0x0FFF
	cpu.Registers.SetFlag(FlagH, isHalfCarry)

	isCarry := sum > 0xFFFF
	cpu.Registers.SetFlag(FlagC, isCarry)

	cpu.Registers.SetHL(uint16(sum))

	return 8
}

func (cpu *CPU) AddC(value uint8) int {
	a := cpu.Registers.A
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}

	result := a + value + carryIn

	// Flags
	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, ((a&0x0F)+(value&0x0F)+carryIn) > 0x0F)
	cpu.Registers.SetFlag(
		FlagC,
		uint16(a)+uint16(value)+uint16(carryIn) > 0xFF,
	)

	cpu.Registers.A = result
	return 4
}

func (cpu *CPU) Sub(value uint8) int {
	a := cpu.Registers.A
	result := a - value

	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, true)
	cpu.Registers.SetFlag(FlagH, halfCarrySub(a, value)) // half borrow
	cpu.Registers.SetFlag(FlagC, a < value)              // borrow

	cpu.Registers.A = result
	return 4
}

func (cpu *CPU) SubC(value uint8) int {
	a := cpu.Registers.A
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}

	result := a - value - carryIn

	// Flags
	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, true)

	// Half-borrow: checks bit 3 → bit 4 borrow
	cpu.Registers.SetFlag(
		FlagH,
		(a&0x0F) < ((value&0x0F)+carryIn),
	)

	// Full borrow
	cpu.Registers.SetFlag(
		FlagC,
		uint16(a) < uint16(value)+uint16(carryIn),
	)

	cpu.Registers.A = result
	// fmt.Printf("DEBUG: SUB 0x%02X from A=0x%02X\n", value, a)
	return 4
}

// ADD
func (cpu *CPU) ADD_B() int {
	return cpu.Add(cpu.Registers.B)
}
func (cpu *CPU) ADD_C() int {
	return cpu.Add(cpu.Registers.C)
}
func (cpu *CPU) ADD_D() int {
	return cpu.Add(cpu.Registers.D)
}
func (cpu *CPU) ADD_E() int {
	return cpu.Add(cpu.Registers.E)
}
func (cpu *CPU) ADD_H() int {
	return cpu.Add(cpu.Registers.H)
}
func (cpu *CPU) ADD_L() int {
	return cpu.Add(cpu.Registers.L)
}
func (cpu *CPU) ADD_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	return cpu.Add(val) + 4
}
func (cpu *CPU) ADD_A() int {
	return cpu.Add(cpu.Registers.A)
}

func (cpu *CPU) ADD_d8() int {
	val := cpu.FetchByte()
	return cpu.Add(val) + 4
}

/*
OPCODE: 0xE8
DESCRIPTION: ADD SP, r8 (Add signed immediate to SP)
CYCLES: 16
*/
func (cpu *CPU) ADD_SP_r8() int {

	raw := cpu.FetchByte()
	offset := int8(raw)

	sp := cpu.Registers.SP

	// We cast to int32 to handle the negative addition cleanly
	result := uint16(int32(sp) + int32(offset))

	// 3. Update Flags (Z=0, N=0)
	cpu.Registers.SetFlag(FlagZ, false)
	cpu.Registers.SetFlag(FlagN, false)

	cpu.Registers.SetFlag(
		FlagH,
		(sp&0x0F)+uint16(raw&0x0F) > 0x0F,
	)
	cpu.Registers.SetFlag(
		FlagC,
		(sp&0xFF)+uint16(raw) > 0xFF,
	)

	cpu.Registers.SP = result

	return 16
}

// ADC
func (cpu *CPU) ADC_B() int {
	return cpu.AddC(cpu.Registers.B)
}
func (cpu *CPU) ADC_C() int {
	return cpu.AddC(cpu.Registers.C)
}
func (cpu *CPU) ADC_D() int {
	return cpu.AddC(cpu.Registers.D)
}
func (cpu *CPU) ADC_E() int {
	return cpu.AddC(cpu.Registers.E)
}
func (cpu *CPU) ADC_H() int {
	return cpu.AddC(cpu.Registers.H)
}
func (cpu *CPU) ADC_L() int {
	return cpu.AddC(cpu.Registers.L)
}
func (cpu *CPU) ADC_HL_ad() int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	return cpu.AddC(val) + 4
}
func (cpu *CPU) ADC_A() int {
	return cpu.AddC(cpu.Registers.A)
}
func (cpu *CPU) ADC_d8() int {
	val := cpu.FetchByte()
	return cpu.AddC(val) + 4
}

// SUB
func (cpu *CPU) SUB_B() int {
	return cpu.Sub(cpu.Registers.B)
}
func (cpu *CPU) SUB_C() int {
	return cpu.Sub(cpu.Registers.C)
}
func (cpu *CPU) SUB_D() int {
	return cpu.Sub(cpu.Registers.D)
}
func (cpu *CPU) SUB_E() int {
	return cpu.Sub(cpu.Registers.E)
}
func (cpu *CPU) SUB_H() int {
	return cpu.Sub(cpu.Registers.H)
}
func (cpu *CPU) SUB_L() int {
	return cpu.Sub(cpu.Registers.L)
}
func (cpu *CPU) SUB_HL_ad() int {
	SUBr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(SUBr)
	return cpu.Sub(val) + 4
}
func (cpu *CPU) SUB_A() int {
	return cpu.Sub(cpu.Registers.A)
}

func (cpu *CPU) SUB_d8() int {
	val := cpu.FetchByte()
	return cpu.Sub(val) + 4
}

// SBC
func (cpu *CPU) SBC_B() int {
	return cpu.SubC(cpu.Registers.B)
}
func (cpu *CPU) SBC_C() int {
	return cpu.SubC(cpu.Registers.C)
}
func (cpu *CPU) SBC_D() int {
	return cpu.SubC(cpu.Registers.D)
}
func (cpu *CPU) SBC_E() int {
	return cpu.SubC(cpu.Registers.E)
}
func (cpu *CPU) SBC_H() int {
	return cpu.SubC(cpu.Registers.H)
}
func (cpu *CPU) SBC_L() int {
	return cpu.SubC(cpu.Registers.L)
}
func (cpu *CPU) SBC_HL_ad() int {
	SUBr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(SUBr)
	return cpu.SubC(val) + 4
}
func (cpu *CPU) SBC_A() int {
	return cpu.SubC(cpu.Registers.A)
}
func (cpu *CPU) SBC_d8() int {
	val := cpu.FetchByte()
	return cpu.SubC(val) + 4
}

func (cpu *CPU) ADD_HL_BC() int {
	return cpu.add_HL_rr(cpu.Registers.GetBC())
}

func (cpu *CPU) ADD_HL_DE() int {
	return cpu.add_HL_rr(cpu.Registers.GetDE())
}

func (cpu *CPU) ADD_HL_HL() int {
	return cpu.add_HL_rr(cpu.Registers.GetHL())
}

func (cpu *CPU) ADD_HL_SP() int {
	return cpu.add_HL_rr(cpu.Registers.SP)
}

/*
OPCODE: 0x27
DESCRIPTION: DAA (Decimal Adjust Accumulator)
CYCLES: 4
*/
func (cpu *CPU) DAA() int {
	a := cpu.Registers.A
	adjust := uint8(0)
	carry := false

	if cpu.Registers.GetFlag(FlagH) || (!cpu.Registers.GetFlag(FlagN) && (a&0x0F) > 9) {
		adjust |= 0x06
	}

	if cpu.Registers.GetFlag(FlagC) || (!cpu.Registers.GetFlag(FlagN) && a > 0x99) {
		adjust |= 0x60
		carry = true
	}

	if cpu.Registers.GetFlag(FlagN) {
		a -= adjust
	} else {
		a += adjust
	}

	cpu.Registers.A = a
	cpu.Registers.SetFlag(FlagZ, a == 0)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, carry)
	return 4
}

/*
OPCODE: 0x37
DESCRIPTION: SCF (Set Carry Flag)
CYCLES: 4
*/
func (cpu *CPU) SCF() int {
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, true)
	return 4
}

/*
OPCODE: 0x3F
DESCRIPTION: CCF (Complement Carry Flag)
CYCLES: 4
*/
func (cpu *CPU) CCF() int {
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, !cpu.Registers.GetFlag(FlagC))
	return 4
}
