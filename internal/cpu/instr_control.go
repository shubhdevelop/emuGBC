package cpu

func (cpu *CPU) InstrCONT_DI() int {
	cpu.IME = false
	return 4
}

func (cpu *CPU) InstrCONT_EI() int {
	cpu.IME = true
	return 4
}

func (cpu *CPU) InstrCONT_HALT() int {
	cpu.isHalted = true
	return 4
}

/*
When the CPU sees the opcode 0x10,
it is hardwired to automatically consume the next byte
(which is almost always 0x00) before going to sleep.
*/
func (cpu *CPU) InstrCONT_STOP() int {
	//It's part of the instruction!
	cpu.Registers.PC++
	cpu.isHalted = true
	return 4
}
