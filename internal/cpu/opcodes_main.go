package cpu

import "fmt"

var mainOpcodes = [256]Opcode{
	// MOVE OPCODES
	0x00: {Fn: (*CPU).InstrNop, Mnemonic: "NOP"},
	0xCB: {Fn: (*CPU).InstrPrefixCB, Mnemonic: "PREFIX CB"},
	0x01: {Fn: (*CPU).InstrLD_BC_d16, Mnemonic: "LD BC, d16"},
	0x02: {Fn: (*CPU).InstrLD_BC_A, Mnemonic: "LD (BC), A"},
	0x06: {Fn: (*CPU).InstrLD_B_d8, Mnemonic: "LD B, d8"},
	0x08: {Fn: (*CPU).InstrLD_a16_SP, Mnemonic: "LD (a16), SP"},
	0x0A: {Fn: (*CPU).InstrLD_A_BC, Mnemonic: "LD A, (BC)"},
	0x0E: {Fn: (*CPU).InstrLD_C_d8, Mnemonic: "LD C, d8"},
	0x11: {Fn: (*CPU).InstrLD_DE_d16, Mnemonic: "LD DE, n16"},
	0x12: {Fn: (*CPU).InstrLD_DE_A, Mnemonic: "LD [DE], A"},
	0x16: {Fn: (*CPU).InstrLD_D_d8, Mnemonic: "LD D, n8"},
	0x1A: {Fn: (*CPU).InstrLD_A_DE, Mnemonic: "LD A, [DE]"},
	0x1E: {Fn: (*CPU).InstrLD_E_d8, Mnemonic: "LD E, n8"},
	0x21: {},
	0x22: {},
	0x26: {},
	0x2A: {},
	0x2E: {},
	0x31: {},
	0x32: {},
	0x36: {},
	0x3A: {},
	0x3E: {},
	0x40: {},
	0x41: {},
	0x42: {},
	0x44: {},
	0x45: {},
	0x46: {},
	0x47: {},
	0x48: {},
	0x49: {},
	0x4A: {},
	0x4B: {},
	0x4C: {},
	0x4D: {},
	0x4E: {},
	0x4F: {},
	0x50: {},
	0x51: {},
	0x52: {},
	0x54: {},
	0x55: {},
	0x56: {},
	0x57: {},
	0x58: {},
	0x59: {},
	0x5A: {},
	0x5B: {},
	0x5C: {},
	0x5D: {},
	0x5E: {},
	0x5F: {},
	0x60: {},
	0x61: {},
	0x62: {},
	0x64: {},
	0x65: {},
	0x66: {},
	0x67: {},
	0x68: {},
	0x69: {},
	0x6A: {},
	0x6B: {},
	0x6C: {},
	0x6D: {},
	0x6E: {},
	0x6F: {},
	0x70: {},
	0x71: {},
	0x72: {},
	0x74: {},
	0x75: {},
	// 0x76: Not Part of the Load instructions
	0x77: {},
	0x78: {},
	0x79: {},
	0x7A: {},
	0x7B: {},
	0x7C: {},
	0x7D: {},
	0x7E: {},
	0x7F: {},
	0xE0: {},
	0xE2: {},
	0xEA: {},
	0xF0: {},
	0xF2: {},
	0xFA: {},
	// Move Instructions STACK:Pop
	0xC1: {},
	0xD1: {},
	0xE1: {},
	0xF1: {},
	//Jumps
	0x20: {Fn: (*CPU).JumpRelativeNotZero, Mnemonic: "JR NZ, r8"},
	0x30: {Fn: (*CPU).JumpRelativeNotCarry, Mnemonic: "JR NC, r8"},
	0x28: {Fn: (*CPU).JumpRelativeZero, Mnemonic: "JR Z, r8"},
	0x38: {Fn: (*CPU).JumpRelativeCarry, Mnemonic: "JR C, r8"},
	0x18: {Fn: (*CPU).JumpRelative, Mnemonic: "JR r8"},
	0xCA: {Fn: (*CPU).JumpAbsoluteNotZero, Mnemonic: "JP NZ a16"},
	0xDA: {Fn: (*CPU).JumpAbsoluteNotCarry, Mnemonic: "JP NC a16"},
	0xC2: {Fn: (*CPU).JumpAbsoluteZero, Mnemonic: "JP 16"},
	0xD2: {Fn: (*CPU).JumpAbsoluteCarry, Mnemonic: "JP C a16"},
	0xC3: {Fn: (*CPU).JumpAbsolute, Mnemonic: "JP a16"},
	0xE9: {Fn: (*CPU).JumpAbsoluteHL, Mnemonic: "JP (HL)"},
	// Move Instructions STACK:Push
	0xC5: {},
	0xD5: {},
	0xE5: {},
	0xF5: {},
	// Misc
	0xF8: {},
	0xF9: {},
}

/*
OPCODE: 0x00
DESCRIPTION: No operation
CYCLES: 4
*/
func (cpu *CPU) InstrNop() int {
	return 4
}

func (cpu *CPU) InstrPrefixCB() int {
	cb := cpu.FetchByte()
	fn := cbOpcodes[cb].Fn
	if fn == nil {
		panic(fmt.Sprintf("Undefined CB opcode %02X at PC: %04X", cb, cpu.Registers.PC-1))
	}
	return fn(cpu)
}
