package cpu

func (cpu *CPU) InstrSTACK_POP_BC() int {
	val := cpu.Pop()
	cpu.Registers.SetBC(val)
	return 12
}
func (cpu *CPU) InstrSTACK_POP_DE() int {
	val := cpu.Pop()
	cpu.Registers.SetDE(val)
	return 12
}
func (cpu *CPU) InstrSTACK_POP_HL() int {
	val := cpu.Pop()
	cpu.Registers.SetHL(val)
	return 12
}
func (cpu *CPU) InstrSTACK_POP_AF() int {
	val := cpu.Pop()
	cpu.Registers.SetAF(val)
	return 12
}

func (cpu *CPU) InstrSTACK_PUSH_BC() int {
	cpu.Push(cpu.Registers.GetBC())
	return 16
}
func (cpu *CPU) InstrSTACK_PUSH_DE() int {
	cpu.Push(cpu.Registers.GetDE())
	return 16
}
func (cpu *CPU) InstrSTACK_PUSH_HL() int {
	cpu.Push(cpu.Registers.GetHL())
	return 16
}
func (cpu *CPU) InstrSTACK_PUSH_AF() int {
	cpu.Push(cpu.Registers.GetAF())
	return 16
}
