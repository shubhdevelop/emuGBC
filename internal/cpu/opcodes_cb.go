package cpu

var cbOpcodes = [256]Opcode{
	// 0x00: {Fn: (*CPU).InstrNop, Cycles: 4},
	0x37: {Fn: (*CPU).SWAP_A, Mnemonic: "SWAP A"},
}

func (cpu *CPU) SWAP_A() int {
	val := cpu.Registers.A

	// Perform the Swap
	result := (val << 4) | (val >> 4)

	cpu.Registers.A = result

	// Update Flags
	cpu.Registers.SetFlag(FlagZ, result == 0) // Only true if A was 0x00
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, false)

	// Return 8 Cycles (4 fetch CB + 4 fetch 37)
	return 8
}
