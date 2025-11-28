package cpu

var cbOpcodes = [256]Opcode{
	// RLC (Rotate Left Circular)
	0x00: {Fn: (*CPU).cbRLC_B, Mnemonic: "RLC B"},
	0x01: {Fn: (*CPU).cbRLC_C, Mnemonic: "RLC C"},
	0x02: {Fn: (*CPU).cbRLC_D, Mnemonic: "RLC D"},
	0x03: {Fn: (*CPU).cbRLC_E, Mnemonic: "RLC E"},
	0x04: {Fn: (*CPU).cbRLC_H, Mnemonic: "RLC H"},
	0x05: {Fn: (*CPU).cbRLC_L, Mnemonic: "RLC L"},
	0x06: {Fn: (*CPU).cbRLC_HL, Mnemonic: "RLC (HL)"},
	0x07: {Fn: (*CPU).cbRLC_A, Mnemonic: "RLC A"},

	// RRC (Rotate Right Circular)
	0x08: {Fn: (*CPU).cbRRC_B, Mnemonic: "RRC B"},
	0x09: {Fn: (*CPU).cbRRC_C, Mnemonic: "RRC C"},
	0x0A: {Fn: (*CPU).cbRRC_D, Mnemonic: "RRC D"},
	0x0B: {Fn: (*CPU).cbRRC_E, Mnemonic: "RRC E"},
	0x0C: {Fn: (*CPU).cbRRC_H, Mnemonic: "RRC H"},
	0x0D: {Fn: (*CPU).cbRRC_L, Mnemonic: "RRC L"},
	0x0E: {Fn: (*CPU).cbRRC_HL, Mnemonic: "RRC (HL)"},
	0x0F: {Fn: (*CPU).cbRRC_A, Mnemonic: "RRC A"},

	// RL (Rotate Left through Carry)
	0x10: {Fn: (*CPU).cbRL_B, Mnemonic: "RL B"},
	0x11: {Fn: (*CPU).cbRL_C, Mnemonic: "RL C"},
	0x12: {Fn: (*CPU).cbRL_D, Mnemonic: "RL D"},
	0x13: {Fn: (*CPU).cbRL_E, Mnemonic: "RL E"},
	0x14: {Fn: (*CPU).cbRL_H, Mnemonic: "RL H"},
	0x15: {Fn: (*CPU).cbRL_L, Mnemonic: "RL L"},
	0x16: {Fn: (*CPU).cbRL_HL, Mnemonic: "RL (HL)"},
	0x17: {Fn: (*CPU).cbRL_A, Mnemonic: "RL A"},

	// RR (Rotate Right through Carry)
	0x18: {Fn: (*CPU).cbRR_B, Mnemonic: "RR B"},
	0x19: {Fn: (*CPU).cbRR_C, Mnemonic: "RR C"},
	0x1A: {Fn: (*CPU).cbRR_D, Mnemonic: "RR D"},
	0x1B: {Fn: (*CPU).cbRR_E, Mnemonic: "RR E"},
	0x1C: {Fn: (*CPU).cbRR_H, Mnemonic: "RR H"},
	0x1D: {Fn: (*CPU).cbRR_L, Mnemonic: "RR L"},
	0x1E: {Fn: (*CPU).cbRR_HL, Mnemonic: "RR (HL)"},
	0x1F: {Fn: (*CPU).cbRR_A, Mnemonic: "RR A"},

	// SLA (Shift Left Arithmetic)
	0x20: {Fn: (*CPU).cbSLA_B, Mnemonic: "SLA B"},
	0x21: {Fn: (*CPU).cbSLA_C, Mnemonic: "SLA C"},
	0x22: {Fn: (*CPU).cbSLA_D, Mnemonic: "SLA D"},
	0x23: {Fn: (*CPU).cbSLA_E, Mnemonic: "SLA E"},
	0x24: {Fn: (*CPU).cbSLA_H, Mnemonic: "SLA H"},
	0x25: {Fn: (*CPU).cbSLA_L, Mnemonic: "SLA L"},
	0x26: {Fn: (*CPU).cbSLA_HL, Mnemonic: "SLA (HL)"},
	0x27: {Fn: (*CPU).cbSLA_A, Mnemonic: "SLA A"},

	// SRA (Shift Right Arithmetic)
	0x28: {Fn: (*CPU).cbSRA_B, Mnemonic: "SRA B"},
	0x29: {Fn: (*CPU).cbSRA_C, Mnemonic: "SRA C"},
	0x2A: {Fn: (*CPU).cbSRA_D, Mnemonic: "SRA D"},
	0x2B: {Fn: (*CPU).cbSRA_E, Mnemonic: "SRA E"},
	0x2C: {Fn: (*CPU).cbSRA_H, Mnemonic: "SRA H"},
	0x2D: {Fn: (*CPU).cbSRA_L, Mnemonic: "SRA L"},
	0x2E: {Fn: (*CPU).cbSRA_HL, Mnemonic: "SRA (HL)"},
	0x2F: {Fn: (*CPU).cbSRA_A, Mnemonic: "SRA A"},

	// SWAP (Swap nibbles)
	0x30: {Fn: (*CPU).cbSWAP_B, Mnemonic: "SWAP B"},
	0x31: {Fn: (*CPU).cbSWAP_C, Mnemonic: "SWAP C"},
	0x32: {Fn: (*CPU).cbSWAP_D, Mnemonic: "SWAP D"},
	0x33: {Fn: (*CPU).cbSWAP_E, Mnemonic: "SWAP E"},
	0x34: {Fn: (*CPU).cbSWAP_H, Mnemonic: "SWAP H"},
	0x35: {Fn: (*CPU).cbSWAP_L, Mnemonic: "SWAP L"},
	0x36: {Fn: (*CPU).cbSWAP_HL, Mnemonic: "SWAP (HL)"},
	0x37: {Fn: (*CPU).cbSWAP_A, Mnemonic: "SWAP A"},

	// SRL (Shift Right Logical)
	0x38: {Fn: (*CPU).cbSRL_B, Mnemonic: "SRL B"},
	0x39: {Fn: (*CPU).cbSRL_C, Mnemonic: "SRL C"},
	0x3A: {Fn: (*CPU).cbSRL_D, Mnemonic: "SRL D"},
	0x3B: {Fn: (*CPU).cbSRL_E, Mnemonic: "SRL E"},
	0x3C: {Fn: (*CPU).cbSRL_H, Mnemonic: "SRL H"},
	0x3D: {Fn: (*CPU).cbSRL_L, Mnemonic: "SRL L"},
	0x3E: {Fn: (*CPU).cbSRL_HL, Mnemonic: "SRL (HL)"},
	0x3F: {Fn: (*CPU).cbSRL_A, Mnemonic: "SRL A"},

	// BIT 0
	0x40: {Fn: (*CPU).cbBIT_0_B, Mnemonic: "BIT 0, B"},
	0x41: {Fn: (*CPU).cbBIT_0_C, Mnemonic: "BIT 0, C"},
	0x42: {Fn: (*CPU).cbBIT_0_D, Mnemonic: "BIT 0, D"},
	0x43: {Fn: (*CPU).cbBIT_0_E, Mnemonic: "BIT 0, E"},
	0x44: {Fn: (*CPU).cbBIT_0_H, Mnemonic: "BIT 0, H"},
	0x45: {Fn: (*CPU).cbBIT_0_L, Mnemonic: "BIT 0, L"},
	0x46: {Fn: (*CPU).cbBIT_0_HL, Mnemonic: "BIT 0, (HL)"},
	0x47: {Fn: (*CPU).cbBIT_0_A, Mnemonic: "BIT 0, A"},

	// BIT 1
	0x48: {Fn: (*CPU).cbBIT_1_B, Mnemonic: "BIT 1, B"},
	0x49: {Fn: (*CPU).cbBIT_1_C, Mnemonic: "BIT 1, C"},
	0x4A: {Fn: (*CPU).cbBIT_1_D, Mnemonic: "BIT 1, D"},
	0x4B: {Fn: (*CPU).cbBIT_1_E, Mnemonic: "BIT 1, E"},
	0x4C: {Fn: (*CPU).cbBIT_1_H, Mnemonic: "BIT 1, H"},
	0x4D: {Fn: (*CPU).cbBIT_1_L, Mnemonic: "BIT 1, L"},
	0x4E: {Fn: (*CPU).cbBIT_1_HL, Mnemonic: "BIT 1, (HL)"},
	0x4F: {Fn: (*CPU).cbBIT_1_A, Mnemonic: "BIT 1, A"},

	// BIT 2
	0x50: {Fn: (*CPU).cbBIT_2_B, Mnemonic: "BIT 2, B"},
	0x51: {Fn: (*CPU).cbBIT_2_C, Mnemonic: "BIT 2, C"},
	0x52: {Fn: (*CPU).cbBIT_2_D, Mnemonic: "BIT 2, D"},
	0x53: {Fn: (*CPU).cbBIT_2_E, Mnemonic: "BIT 2, E"},
	0x54: {Fn: (*CPU).cbBIT_2_H, Mnemonic: "BIT 2, H"},
	0x55: {Fn: (*CPU).cbBIT_2_L, Mnemonic: "BIT 2, L"},
	0x56: {Fn: (*CPU).cbBIT_2_HL, Mnemonic: "BIT 2, (HL)"},
	0x57: {Fn: (*CPU).cbBIT_2_A, Mnemonic: "BIT 2, A"},

	// BIT 3
	0x58: {Fn: (*CPU).cbBIT_3_B, Mnemonic: "BIT 3, B"},
	0x59: {Fn: (*CPU).cbBIT_3_C, Mnemonic: "BIT 3, C"},
	0x5A: {Fn: (*CPU).cbBIT_3_D, Mnemonic: "BIT 3, D"},
	0x5B: {Fn: (*CPU).cbBIT_3_E, Mnemonic: "BIT 3, E"},
	0x5C: {Fn: (*CPU).cbBIT_3_H, Mnemonic: "BIT 3, H"},
	0x5D: {Fn: (*CPU).cbBIT_3_L, Mnemonic: "BIT 3, L"},
	0x5E: {Fn: (*CPU).cbBIT_3_HL, Mnemonic: "BIT 3, (HL)"},
	0x5F: {Fn: (*CPU).cbBIT_3_A, Mnemonic: "BIT 3, A"},

	// BIT 4
	0x60: {Fn: (*CPU).cbBIT_4_B, Mnemonic: "BIT 4, B"},
	0x61: {Fn: (*CPU).cbBIT_4_C, Mnemonic: "BIT 4, C"},
	0x62: {Fn: (*CPU).cbBIT_4_D, Mnemonic: "BIT 4, D"},
	0x63: {Fn: (*CPU).cbBIT_4_E, Mnemonic: "BIT 4, E"},
	0x64: {Fn: (*CPU).cbBIT_4_H, Mnemonic: "BIT 4, H"},
	0x65: {Fn: (*CPU).cbBIT_4_L, Mnemonic: "BIT 4, L"},
	0x66: {Fn: (*CPU).cbBIT_4_HL, Mnemonic: "BIT 4, (HL)"},
	0x67: {Fn: (*CPU).cbBIT_4_A, Mnemonic: "BIT 4, A"},

	// BIT 5
	0x68: {Fn: (*CPU).cbBIT_5_B, Mnemonic: "BIT 5, B"},
	0x69: {Fn: (*CPU).cbBIT_5_C, Mnemonic: "BIT 5, C"},
	0x6A: {Fn: (*CPU).cbBIT_5_D, Mnemonic: "BIT 5, D"},
	0x6B: {Fn: (*CPU).cbBIT_5_E, Mnemonic: "BIT 5, E"},
	0x6C: {Fn: (*CPU).cbBIT_5_H, Mnemonic: "BIT 5, H"},
	0x6D: {Fn: (*CPU).cbBIT_5_L, Mnemonic: "BIT 5, L"},
	0x6E: {Fn: (*CPU).cbBIT_5_HL, Mnemonic: "BIT 5, (HL)"},
	0x6F: {Fn: (*CPU).cbBIT_5_A, Mnemonic: "BIT 5, A"},

	// BIT 6
	0x70: {Fn: (*CPU).cbBIT_6_B, Mnemonic: "BIT 6, B"},
	0x71: {Fn: (*CPU).cbBIT_6_C, Mnemonic: "BIT 6, C"},
	0x72: {Fn: (*CPU).cbBIT_6_D, Mnemonic: "BIT 6, D"},
	0x73: {Fn: (*CPU).cbBIT_6_E, Mnemonic: "BIT 6, E"},
	0x74: {Fn: (*CPU).cbBIT_6_H, Mnemonic: "BIT 6, H"},
	0x75: {Fn: (*CPU).cbBIT_6_L, Mnemonic: "BIT 6, L"},
	0x76: {Fn: (*CPU).cbBIT_6_HL, Mnemonic: "BIT 6, (HL)"},
	0x77: {Fn: (*CPU).cbBIT_6_A, Mnemonic: "BIT 6, A"},

	// BIT 7
	0x78: {Fn: (*CPU).cbBIT_7_B, Mnemonic: "BIT 7, B"},
	0x79: {Fn: (*CPU).cbBIT_7_C, Mnemonic: "BIT 7, C"},
	0x7A: {Fn: (*CPU).cbBIT_7_D, Mnemonic: "BIT 7, D"},
	0x7B: {Fn: (*CPU).cbBIT_7_E, Mnemonic: "BIT 7, E"},
	0x7C: {Fn: (*CPU).cbBIT_7_H, Mnemonic: "BIT 7, H"},
	0x7D: {Fn: (*CPU).cbBIT_7_L, Mnemonic: "BIT 7, L"},
	0x7E: {Fn: (*CPU).cbBIT_7_HL, Mnemonic: "BIT 7, (HL)"},
	0x7F: {Fn: (*CPU).cbBIT_7_A, Mnemonic: "BIT 7, A"},

	// RES 0 (Reset bit 0)
	0x80: {Fn: (*CPU).cbRES_0_B, Mnemonic: "RES 0, B"},
	0x81: {Fn: (*CPU).cbRES_0_C, Mnemonic: "RES 0, C"},
	0x82: {Fn: (*CPU).cbRES_0_D, Mnemonic: "RES 0, D"},
	0x83: {Fn: (*CPU).cbRES_0_E, Mnemonic: "RES 0, E"},
	0x84: {Fn: (*CPU).cbRES_0_H, Mnemonic: "RES 0, H"},
	0x85: {Fn: (*CPU).cbRES_0_L, Mnemonic: "RES 0, L"},
	0x86: {Fn: (*CPU).cbRES_0_HL, Mnemonic: "RES 0, (HL)"},
	0x87: {Fn: (*CPU).cbRES_0_A, Mnemonic: "RES 0, A"},

	// RES 1 (Reset bit 1)
	0x88: {Fn: (*CPU).cbRES_1_B, Mnemonic: "RES 1, B"},
	0x89: {Fn: (*CPU).cbRES_1_C, Mnemonic: "RES 1, C"},
	0x8A: {Fn: (*CPU).cbRES_1_D, Mnemonic: "RES 1, D"},
	0x8B: {Fn: (*CPU).cbRES_1_E, Mnemonic: "RES 1, E"},
	0x8C: {Fn: (*CPU).cbRES_1_H, Mnemonic: "RES 1, H"},
	0x8D: {Fn: (*CPU).cbRES_1_L, Mnemonic: "RES 1, L"},
	0x8E: {Fn: (*CPU).cbRES_1_HL, Mnemonic: "RES 1, (HL)"},
	0x8F: {Fn: (*CPU).cbRES_1_A, Mnemonic: "RES 1, A"},

	// RES 2 (Reset bit 2)
	0x90: {Fn: (*CPU).cbRES_2_B, Mnemonic: "RES 2, B"},
	0x91: {Fn: (*CPU).cbRES_2_C, Mnemonic: "RES 2, C"},
	0x92: {Fn: (*CPU).cbRES_2_D, Mnemonic: "RES 2, D"},
	0x93: {Fn: (*CPU).cbRES_2_E, Mnemonic: "RES 2, E"},
	0x94: {Fn: (*CPU).cbRES_2_H, Mnemonic: "RES 2, H"},
	0x95: {Fn: (*CPU).cbRES_2_L, Mnemonic: "RES 2, L"},
	0x96: {Fn: (*CPU).cbRES_2_HL, Mnemonic: "RES 2, (HL)"},
	0x97: {Fn: (*CPU).cbRES_2_A, Mnemonic: "RES 2, A"},

	// RES 3 (Reset bit 3)
	0x98: {Fn: (*CPU).cbRES_3_B, Mnemonic: "RES 3, B"},
	0x99: {Fn: (*CPU).cbRES_3_C, Mnemonic: "RES 3, C"},
	0x9A: {Fn: (*CPU).cbRES_3_D, Mnemonic: "RES 3, D"},
	0x9B: {Fn: (*CPU).cbRES_3_E, Mnemonic: "RES 3, E"},
	0x9C: {Fn: (*CPU).cbRES_3_H, Mnemonic: "RES 3, H"},
	0x9D: {Fn: (*CPU).cbRES_3_L, Mnemonic: "RES 3, L"},
	0x9E: {Fn: (*CPU).cbRES_3_HL, Mnemonic: "RES 3, (HL)"},
	0x9F: {Fn: (*CPU).cbRES_3_A, Mnemonic: "RES 3, A"},

	// RES 4 (Reset bit 4)
	0xA0: {Fn: (*CPU).cbRES_4_B, Mnemonic: "RES 4, B"},
	0xA1: {Fn: (*CPU).cbRES_4_C, Mnemonic: "RES 4, C"},
	0xA2: {Fn: (*CPU).cbRES_4_D, Mnemonic: "RES 4, D"},
	0xA3: {Fn: (*CPU).cbRES_4_E, Mnemonic: "RES 4, E"},
	0xA4: {Fn: (*CPU).cbRES_4_H, Mnemonic: "RES 4, H"},
	0xA5: {Fn: (*CPU).cbRES_4_L, Mnemonic: "RES 4, L"},
	0xA6: {Fn: (*CPU).cbRES_4_HL, Mnemonic: "RES 4, (HL)"},
	0xA7: {Fn: (*CPU).cbRES_4_A, Mnemonic: "RES 4, A"},

	// RES 5 (Reset bit 5)
	0xA8: {Fn: (*CPU).cbRES_5_B, Mnemonic: "RES 5, B"},
	0xA9: {Fn: (*CPU).cbRES_5_C, Mnemonic: "RES 5, C"},
	0xAA: {Fn: (*CPU).cbRES_5_D, Mnemonic: "RES 5, D"},
	0xAB: {Fn: (*CPU).cbRES_5_E, Mnemonic: "RES 5, E"},
	0xAC: {Fn: (*CPU).cbRES_5_H, Mnemonic: "RES 5, H"},
	0xAD: {Fn: (*CPU).cbRES_5_L, Mnemonic: "RES 5, L"},
	0xAE: {Fn: (*CPU).cbRES_5_HL, Mnemonic: "RES 5, (HL)"},
	0xAF: {Fn: (*CPU).cbRES_5_A, Mnemonic: "RES 5, A"},

	// RES 6 (Reset bit 6)
	0xB0: {Fn: (*CPU).cbRES_6_B, Mnemonic: "RES 6, B"},
	0xB1: {Fn: (*CPU).cbRES_6_C, Mnemonic: "RES 6, C"},
	0xB2: {Fn: (*CPU).cbRES_6_D, Mnemonic: "RES 6, D"},
	0xB3: {Fn: (*CPU).cbRES_6_E, Mnemonic: "RES 6, E"},
	0xB4: {Fn: (*CPU).cbRES_6_H, Mnemonic: "RES 6, H"},
	0xB5: {Fn: (*CPU).cbRES_6_L, Mnemonic: "RES 6, L"},
	0xB6: {Fn: (*CPU).cbRES_6_HL, Mnemonic: "RES 6, (HL)"},
	0xB7: {Fn: (*CPU).cbRES_6_A, Mnemonic: "RES 6, A"},

	// RES 7 (Reset bit 7)
	0xB8: {Fn: (*CPU).cbRES_7_B, Mnemonic: "RES 7, B"},
	0xB9: {Fn: (*CPU).cbRES_7_C, Mnemonic: "RES 7, C"},
	0xBA: {Fn: (*CPU).cbRES_7_D, Mnemonic: "RES 7, D"},
	0xBB: {Fn: (*CPU).cbRES_7_E, Mnemonic: "RES 7, E"},
	0xBC: {Fn: (*CPU).cbRES_7_H, Mnemonic: "RES 7, H"},
	0xBD: {Fn: (*CPU).cbRES_7_L, Mnemonic: "RES 7, L"},
	0xBE: {Fn: (*CPU).cbRES_7_HL, Mnemonic: "RES 7, (HL)"},
	0xBF: {Fn: (*CPU).cbRES_7_A, Mnemonic: "RES 7, A"},

	// SET 0 (Set bit 0)
	0xC0: {Fn: (*CPU).cbSET_0_B, Mnemonic: "SET 0, B"},
	0xC1: {Fn: (*CPU).cbSET_0_C, Mnemonic: "SET 0, C"},
	0xC2: {Fn: (*CPU).cbSET_0_D, Mnemonic: "SET 0, D"},
	0xC3: {Fn: (*CPU).cbSET_0_E, Mnemonic: "SET 0, E"},
	0xC4: {Fn: (*CPU).cbSET_0_H, Mnemonic: "SET 0, H"},
	0xC5: {Fn: (*CPU).cbSET_0_L, Mnemonic: "SET 0, L"},
	0xC6: {Fn: (*CPU).cbSET_0_HL, Mnemonic: "SET 0, (HL)"},
	0xC7: {Fn: (*CPU).cbSET_0_A, Mnemonic: "SET 0, A"},

	// SET 1 (Set bit 1)
	0xC8: {Fn: (*CPU).cbSET_1_B, Mnemonic: "SET 1, B"},
	0xC9: {Fn: (*CPU).cbSET_1_C, Mnemonic: "SET 1, C"},
	0xCA: {Fn: (*CPU).cbSET_1_D, Mnemonic: "SET 1, D"},
	0xCB: {Fn: (*CPU).cbSET_1_E, Mnemonic: "SET 1, E"},
	0xCC: {Fn: (*CPU).cbSET_1_H, Mnemonic: "SET 1, H"},
	0xCD: {Fn: (*CPU).cbSET_1_L, Mnemonic: "SET 1, L"},
	0xCE: {Fn: (*CPU).cbSET_1_HL, Mnemonic: "SET 1, (HL)"},
	0xCF: {Fn: (*CPU).cbSET_1_A, Mnemonic: "SET 1, A"},

	// SET 2 (Set bit 2)
	0xD0: {Fn: (*CPU).cbSET_2_B, Mnemonic: "SET 2, B"},
	0xD1: {Fn: (*CPU).cbSET_2_C, Mnemonic: "SET 2, C"},
	0xD2: {Fn: (*CPU).cbSET_2_D, Mnemonic: "SET 2, D"},
	0xD3: {Fn: (*CPU).cbSET_2_E, Mnemonic: "SET 2, E"},
	0xD4: {Fn: (*CPU).cbSET_2_H, Mnemonic: "SET 2, H"},
	0xD5: {Fn: (*CPU).cbSET_2_L, Mnemonic: "SET 2, L"},
	0xD6: {Fn: (*CPU).cbSET_2_HL, Mnemonic: "SET 2, (HL)"},
	0xD7: {Fn: (*CPU).cbSET_2_A, Mnemonic: "SET 2, A"},

	// SET 3 (Set bit 3)
	0xD8: {Fn: (*CPU).cbSET_3_B, Mnemonic: "SET 3, B"},
	0xD9: {Fn: (*CPU).cbSET_3_C, Mnemonic: "SET 3, C"},
	0xDA: {Fn: (*CPU).cbSET_3_D, Mnemonic: "SET 3, D"},
	0xDB: {Fn: (*CPU).cbSET_3_E, Mnemonic: "SET 3, E"},
	0xDC: {Fn: (*CPU).cbSET_3_H, Mnemonic: "SET 3, H"},
	0xDD: {Fn: (*CPU).cbSET_3_L, Mnemonic: "SET 3, L"},
	0xDE: {Fn: (*CPU).cbSET_3_HL, Mnemonic: "SET 3, (HL)"},
	0xDF: {Fn: (*CPU).cbSET_3_A, Mnemonic: "SET 3, A"},

	// SET 4 (Set bit 4)
	0xE0: {Fn: (*CPU).cbSET_4_B, Mnemonic: "SET 4, B"},
	0xE1: {Fn: (*CPU).cbSET_4_C, Mnemonic: "SET 4, C"},
	0xE2: {Fn: (*CPU).cbSET_4_D, Mnemonic: "SET 4, D"},
	0xE3: {Fn: (*CPU).cbSET_4_E, Mnemonic: "SET 4, E"},
	0xE4: {Fn: (*CPU).cbSET_4_H, Mnemonic: "SET 4, H"},
	0xE5: {Fn: (*CPU).cbSET_4_L, Mnemonic: "SET 4, L"},
	0xE6: {Fn: (*CPU).cbSET_4_HL, Mnemonic: "SET 4, (HL)"},
	0xE7: {Fn: (*CPU).cbSET_4_A, Mnemonic: "SET 4, A"},

	// SET 5 (Set bit 5)
	0xE8: {Fn: (*CPU).cbSET_5_B, Mnemonic: "SET 5, B"},
	0xE9: {Fn: (*CPU).cbSET_5_C, Mnemonic: "SET 5, C"},
	0xEA: {Fn: (*CPU).cbSET_5_D, Mnemonic: "SET 5, D"},
	0xEB: {Fn: (*CPU).cbSET_5_E, Mnemonic: "SET 5, E"},
	0xEC: {Fn: (*CPU).cbSET_5_H, Mnemonic: "SET 5, H"},
	0xED: {Fn: (*CPU).cbSET_5_L, Mnemonic: "SET 5, L"},
	0xEE: {Fn: (*CPU).cbSET_5_HL, Mnemonic: "SET 5, (HL)"},
	0xEF: {Fn: (*CPU).cbSET_5_A, Mnemonic: "SET 5, A"},

	// SET 6 (Set bit 6)
	0xF0: {Fn: (*CPU).cbSET_6_B, Mnemonic: "SET 6, B"},
	0xF1: {Fn: (*CPU).cbSET_6_C, Mnemonic: "SET 6, C"},
	0xF2: {Fn: (*CPU).cbSET_6_D, Mnemonic: "SET 6, D"},
	0xF3: {Fn: (*CPU).cbSET_6_E, Mnemonic: "SET 6, E"},
	0xF4: {Fn: (*CPU).cbSET_6_H, Mnemonic: "SET 6, H"},
	0xF5: {Fn: (*CPU).cbSET_6_L, Mnemonic: "SET 6, L"},
	0xF6: {Fn: (*CPU).cbSET_6_HL, Mnemonic: "SET 6, (HL)"},
	0xF7: {Fn: (*CPU).cbSET_6_A, Mnemonic: "SET 6, A"},

	// SET 7 (Set bit 7)
	0xF8: {Fn: (*CPU).cbSET_7_B, Mnemonic: "SET 7, B"},
	0xF9: {Fn: (*CPU).cbSET_7_C, Mnemonic: "SET 7, C"},
	0xFA: {Fn: (*CPU).cbSET_7_D, Mnemonic: "SET 7, D"},
	0xFB: {Fn: (*CPU).cbSET_7_E, Mnemonic: "SET 7, E"},
	0xFC: {Fn: (*CPU).cbSET_7_H, Mnemonic: "SET 7, H"},
	0xFD: {Fn: (*CPU).cbSET_7_L, Mnemonic: "SET 7, L"},
	0xFE: {Fn: (*CPU).cbSET_7_HL, Mnemonic: "SET 7, (HL)"},
	0xFF: {Fn: (*CPU).cbSET_7_A, Mnemonic: "SET 7, A"},
}

// Generic helper functions for CB prefix opcodes

// cbOpR8 performs an operation on an 8-bit register
// op is a function that takes a value and returns (result, carry)
// cycles is the number of cycles for register operations
func (cpu *CPU) cbOpR8(reg *uint8, op func(uint8) (uint8, uint8), cycles int) int {
	val := *reg
	result, carry := op(val)
	*reg = result
	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, carry == 1)
	return cycles
}

// cbOpHL performs an operation on memory at (HL)
// op is a function that takes a value and returns (result, carry)
// cycles is the number of cycles for memory operations
func (cpu *CPU) cbOpHL(op func(uint8) (uint8, uint8), cycles int) int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	result, carry := op(val)
	cpu.Bus.Write(addr, result)
	cpu.Registers.SetFlag(FlagZ, result == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, false)
	cpu.Registers.SetFlag(FlagC, carry == 1)
	return cycles
}

// Operation functions
func rlcOp(val uint8) (uint8, uint8) {
	bit7 := (val >> 7) & 0x01
	result := (val << 1) | bit7
	return result, bit7
}

func rrcOp(val uint8) (uint8, uint8) {
	bit0 := val & 0x01
	result := (val >> 1) | (bit0 << 7)
	return result, bit0
}

func rlOp(carryIn uint8) func(uint8) (uint8, uint8) {
	return func(v uint8) (uint8, uint8) {
		newCarry := (v >> 7) & 0x01
		result := (v << 1) | carryIn
		return result, newCarry
	}
}

func rrOp(carryIn uint8) func(uint8) (uint8, uint8) {
	return func(v uint8) (uint8, uint8) {
		newCarry := v & 0x01
		result := (v >> 1) | (carryIn << 7)
		return result, newCarry
	}
}

func slaOp(val uint8) (uint8, uint8) {
	carry := (val >> 7) & 0x01
	result := val << 1
	return result, carry
}

func sraOp(val uint8) (uint8, uint8) {
	carry := val & 0x01
	msb := val & 0x80
	result := (val >> 1) | msb
	return result, carry
}

func swapOp(val uint8) (uint8, uint8) {
	result := (val << 4) | (val >> 4)
	return result, 0
}

func srlOp(val uint8) (uint8, uint8) {
	carry := val & 0x01
	result := val >> 1
	return result, carry
}

func resOp(bit uint8) func(uint8) (uint8, uint8) {
	return func(val uint8) (uint8, uint8) {
		mask := ^(uint8(1 << bit))
		result := val & mask
		return result, 0
	}
}

func setOp(bit uint8) func(uint8) (uint8, uint8) {
	return func(val uint8) (uint8, uint8) {
		mask := uint8(1 << bit)
		result := val | mask
		return result, 0
	}
}

// RLC (Rotate Left Circular) Functions
func (cpu *CPU) cbRLC_B() int {
	return cpu.cbOpR8(&cpu.Registers.B, rlcOp, 8)
}

func (cpu *CPU) cbRLC_C() int {
	return cpu.cbOpR8(&cpu.Registers.C, rlcOp, 8)
}

func (cpu *CPU) cbRLC_D() int {
	return cpu.cbOpR8(&cpu.Registers.D, rlcOp, 8)
}

func (cpu *CPU) cbRLC_E() int {
	return cpu.cbOpR8(&cpu.Registers.E, rlcOp, 8)
}

func (cpu *CPU) cbRLC_H() int {
	return cpu.cbOpR8(&cpu.Registers.H, rlcOp, 8)
}

func (cpu *CPU) cbRLC_L() int {
	return cpu.cbOpR8(&cpu.Registers.L, rlcOp, 8)
}

func (cpu *CPU) cbRLC_HL() int {
	return cpu.cbOpHL(rlcOp, 16)
}

func (cpu *CPU) cbRLC_A() int {
	return cpu.cbOpR8(&cpu.Registers.A, rlcOp, 8)
}

// RRC (Rotate Right Circular) Functions
func (cpu *CPU) cbRRC_B() int {
	return cpu.cbOpR8(&cpu.Registers.B, rrcOp, 8)
}

func (cpu *CPU) cbRRC_C() int {
	return cpu.cbOpR8(&cpu.Registers.C, rrcOp, 8)
}

func (cpu *CPU) cbRRC_D() int {
	return cpu.cbOpR8(&cpu.Registers.D, rrcOp, 8)
}

func (cpu *CPU) cbRRC_E() int {
	return cpu.cbOpR8(&cpu.Registers.E, rrcOp, 8)
}

func (cpu *CPU) cbRRC_H() int {
	return cpu.cbOpR8(&cpu.Registers.H, rrcOp, 8)
}

func (cpu *CPU) cbRRC_L() int {
	return cpu.cbOpR8(&cpu.Registers.L, rrcOp, 8)
}

func (cpu *CPU) cbRRC_HL() int {
	return cpu.cbOpHL(rrcOp, 16)
}

func (cpu *CPU) cbRRC_A() int {
	return cpu.cbOpR8(&cpu.Registers.A, rrcOp, 8)
}

// RL (Rotate Left through Carry) Functions
func (cpu *CPU) cbRL_B() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.B, rlOp(carryIn), 8)
}

func (cpu *CPU) cbRL_C() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.C, rlOp(carryIn), 8)
}

func (cpu *CPU) cbRL_D() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.D, rlOp(carryIn), 8)
}

func (cpu *CPU) cbRL_E() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.E, rlOp(carryIn), 8)
}

func (cpu *CPU) cbRL_H() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.H, rlOp(carryIn), 8)
}

func (cpu *CPU) cbRL_L() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.L, rlOp(carryIn), 8)
}

func (cpu *CPU) cbRL_HL() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpHL(rlOp(carryIn), 16)
}

func (cpu *CPU) cbRL_A() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.A, rlOp(carryIn), 8)
}

// RR (Rotate Right through Carry) Functions
func (cpu *CPU) cbRR_B() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.B, rrOp(carryIn), 8)
}

func (cpu *CPU) cbRR_C() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.C, rrOp(carryIn), 8)
}

func (cpu *CPU) cbRR_D() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.D, rrOp(carryIn), 8)
}

func (cpu *CPU) cbRR_E() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.E, rrOp(carryIn), 8)
}

func (cpu *CPU) cbRR_H() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.H, rrOp(carryIn), 8)
}

func (cpu *CPU) cbRR_L() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.L, rrOp(carryIn), 8)
}

func (cpu *CPU) cbRR_HL() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpHL(rrOp(carryIn), 16)
}

func (cpu *CPU) cbRR_A() int {
	carryIn := uint8(0)
	if cpu.Registers.GetFlag(FlagC) {
		carryIn = 1
	}
	return cpu.cbOpR8(&cpu.Registers.A, rrOp(carryIn), 8)
}

// SLA (Shift Left Arithmetic) Functions
func (cpu *CPU) cbSLA_B() int {
	return cpu.cbOpR8(&cpu.Registers.B, slaOp, 8)
}

func (cpu *CPU) cbSLA_C() int {
	return cpu.cbOpR8(&cpu.Registers.C, slaOp, 8)
}

func (cpu *CPU) cbSLA_D() int {
	return cpu.cbOpR8(&cpu.Registers.D, slaOp, 8)
}

func (cpu *CPU) cbSLA_E() int {
	return cpu.cbOpR8(&cpu.Registers.E, slaOp, 8)
}

func (cpu *CPU) cbSLA_H() int {
	return cpu.cbOpR8(&cpu.Registers.H, slaOp, 8)
}

func (cpu *CPU) cbSLA_L() int {
	return cpu.cbOpR8(&cpu.Registers.L, slaOp, 8)
}

func (cpu *CPU) cbSLA_HL() int {
	return cpu.cbOpHL(slaOp, 16)
}

func (cpu *CPU) cbSLA_A() int {
	return cpu.cbOpR8(&cpu.Registers.A, slaOp, 8)
}

// SRA (Shift Right Arithmetic) Functions
func (cpu *CPU) cbSRA_B() int {
	return cpu.cbOpR8(&cpu.Registers.B, sraOp, 8)
}

func (cpu *CPU) cbSRA_C() int {
	return cpu.cbOpR8(&cpu.Registers.C, sraOp, 8)
}

func (cpu *CPU) cbSRA_D() int {
	return cpu.cbOpR8(&cpu.Registers.D, sraOp, 8)
}

func (cpu *CPU) cbSRA_E() int {
	return cpu.cbOpR8(&cpu.Registers.E, sraOp, 8)
}

func (cpu *CPU) cbSRA_H() int {
	return cpu.cbOpR8(&cpu.Registers.H, sraOp, 8)
}

func (cpu *CPU) cbSRA_L() int {
	return cpu.cbOpR8(&cpu.Registers.L, sraOp, 8)
}

func (cpu *CPU) cbSRA_HL() int {
	return cpu.cbOpHL(sraOp, 16)
}

func (cpu *CPU) cbSRA_A() int {
	return cpu.cbOpR8(&cpu.Registers.A, sraOp, 8)
}

// SWAP Functions
func (cpu *CPU) cbSWAP_B() int {
	return cpu.cbOpR8(&cpu.Registers.B, swapOp, 8)
}

func (cpu *CPU) cbSWAP_C() int {
	return cpu.cbOpR8(&cpu.Registers.C, swapOp, 8)
}

func (cpu *CPU) cbSWAP_D() int {
	return cpu.cbOpR8(&cpu.Registers.D, swapOp, 8)
}

func (cpu *CPU) cbSWAP_E() int {
	return cpu.cbOpR8(&cpu.Registers.E, swapOp, 8)
}

func (cpu *CPU) cbSWAP_H() int {
	return cpu.cbOpR8(&cpu.Registers.H, swapOp, 8)
}

func (cpu *CPU) cbSWAP_L() int {
	return cpu.cbOpR8(&cpu.Registers.L, swapOp, 8)
}

func (cpu *CPU) cbSWAP_HL() int {
	return cpu.cbOpHL(swapOp, 16)
}

func (cpu *CPU) cbSWAP_A() int {
	return cpu.cbOpR8(&cpu.Registers.A, swapOp, 8)
}

// SRL (Shift Right Logical) Functions
func (cpu *CPU) cbSRL_B() int {
	return cpu.cbOpR8(&cpu.Registers.B, srlOp, 8)
}

func (cpu *CPU) cbSRL_C() int {
	return cpu.cbOpR8(&cpu.Registers.C, srlOp, 8)
}

func (cpu *CPU) cbSRL_D() int {
	return cpu.cbOpR8(&cpu.Registers.D, srlOp, 8)
}

func (cpu *CPU) cbSRL_E() int {
	return cpu.cbOpR8(&cpu.Registers.E, srlOp, 8)
}

func (cpu *CPU) cbSRL_H() int {
	return cpu.cbOpR8(&cpu.Registers.H, srlOp, 8)
}

func (cpu *CPU) cbSRL_L() int {
	return cpu.cbOpR8(&cpu.Registers.L, srlOp, 8)
}

func (cpu *CPU) cbSRL_HL() int {
	return cpu.cbOpHL(srlOp, 16)
}

func (cpu *CPU) cbSRL_A() int {
	return cpu.cbOpR8(&cpu.Registers.A, srlOp, 8)
}

// BIT Functions - special handling for BIT operations
func (cpu *CPU) cbOpBIT_R8(reg *uint8, bit uint8, cycles int) int {
	val := *reg
	mask := uint8(1 << bit)
	bitValue := (val & mask) >> bit
	cpu.Registers.SetFlag(FlagZ, bitValue == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, true)
	// FlagC is not affected by BIT
	return cycles
}

func (cpu *CPU) cbOpBIT_HL(bit uint8, cycles int) int {
	addr := cpu.Registers.GetHL()
	val := cpu.Bus.Read(addr)
	mask := uint8(1 << bit)
	bitValue := (val & mask) >> bit
	cpu.Registers.SetFlag(FlagZ, bitValue == 0)
	cpu.Registers.SetFlag(FlagN, false)
	cpu.Registers.SetFlag(FlagH, true)
	// FlagC is not affected by BIT
	return cycles
}

func (cpu *CPU) cbBIT_0_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 0, 8) }
func (cpu *CPU) cbBIT_0_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 0, 8) }
func (cpu *CPU) cbBIT_0_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 0, 8) }
func (cpu *CPU) cbBIT_0_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 0, 8) }
func (cpu *CPU) cbBIT_0_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 0, 8) }
func (cpu *CPU) cbBIT_0_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 0, 8) }
func (cpu *CPU) cbBIT_0_HL() int { return cpu.cbOpBIT_HL(0, 16) }
func (cpu *CPU) cbBIT_0_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 0, 8) }

func (cpu *CPU) cbBIT_1_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 1, 8) }
func (cpu *CPU) cbBIT_1_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 1, 8) }
func (cpu *CPU) cbBIT_1_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 1, 8) }
func (cpu *CPU) cbBIT_1_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 1, 8) }
func (cpu *CPU) cbBIT_1_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 1, 8) }
func (cpu *CPU) cbBIT_1_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 1, 8) }
func (cpu *CPU) cbBIT_1_HL() int { return cpu.cbOpBIT_HL(1, 16) }
func (cpu *CPU) cbBIT_1_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 1, 8) }

func (cpu *CPU) cbBIT_2_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 2, 8) }
func (cpu *CPU) cbBIT_2_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 2, 8) }
func (cpu *CPU) cbBIT_2_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 2, 8) }
func (cpu *CPU) cbBIT_2_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 2, 8) }
func (cpu *CPU) cbBIT_2_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 2, 8) }
func (cpu *CPU) cbBIT_2_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 2, 8) }
func (cpu *CPU) cbBIT_2_HL() int { return cpu.cbOpBIT_HL(2, 16) }
func (cpu *CPU) cbBIT_2_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 2, 8) }

func (cpu *CPU) cbBIT_3_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 3, 8) }
func (cpu *CPU) cbBIT_3_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 3, 8) }
func (cpu *CPU) cbBIT_3_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 3, 8) }
func (cpu *CPU) cbBIT_3_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 3, 8) }
func (cpu *CPU) cbBIT_3_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 3, 8) }
func (cpu *CPU) cbBIT_3_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 3, 8) }
func (cpu *CPU) cbBIT_3_HL() int { return cpu.cbOpBIT_HL(3, 16) }
func (cpu *CPU) cbBIT_3_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 3, 8) }

func (cpu *CPU) cbBIT_4_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 4, 8) }
func (cpu *CPU) cbBIT_4_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 4, 8) }
func (cpu *CPU) cbBIT_4_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 4, 8) }
func (cpu *CPU) cbBIT_4_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 4, 8) }
func (cpu *CPU) cbBIT_4_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 4, 8) }
func (cpu *CPU) cbBIT_4_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 4, 8) }
func (cpu *CPU) cbBIT_4_HL() int { return cpu.cbOpBIT_HL(4, 16) }
func (cpu *CPU) cbBIT_4_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 4, 8) }

func (cpu *CPU) cbBIT_5_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 5, 8) }
func (cpu *CPU) cbBIT_5_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 5, 8) }
func (cpu *CPU) cbBIT_5_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 5, 8) }
func (cpu *CPU) cbBIT_5_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 5, 8) }
func (cpu *CPU) cbBIT_5_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 5, 8) }
func (cpu *CPU) cbBIT_5_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 5, 8) }
func (cpu *CPU) cbBIT_5_HL() int { return cpu.cbOpBIT_HL(5, 16) }
func (cpu *CPU) cbBIT_5_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 5, 8) }

func (cpu *CPU) cbBIT_6_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 6, 8) }
func (cpu *CPU) cbBIT_6_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 6, 8) }
func (cpu *CPU) cbBIT_6_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 6, 8) }
func (cpu *CPU) cbBIT_6_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 6, 8) }
func (cpu *CPU) cbBIT_6_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 6, 8) }
func (cpu *CPU) cbBIT_6_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 6, 8) }
func (cpu *CPU) cbBIT_6_HL() int { return cpu.cbOpBIT_HL(6, 16) }
func (cpu *CPU) cbBIT_6_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 6, 8) }

func (cpu *CPU) cbBIT_7_B() int { return cpu.cbOpBIT_R8(&cpu.Registers.B, 7, 8) }
func (cpu *CPU) cbBIT_7_C() int { return cpu.cbOpBIT_R8(&cpu.Registers.C, 7, 8) }
func (cpu *CPU) cbBIT_7_D() int { return cpu.cbOpBIT_R8(&cpu.Registers.D, 7, 8) }
func (cpu *CPU) cbBIT_7_E() int { return cpu.cbOpBIT_R8(&cpu.Registers.E, 7, 8) }
func (cpu *CPU) cbBIT_7_H() int { return cpu.cbOpBIT_R8(&cpu.Registers.H, 7, 8) }
func (cpu *CPU) cbBIT_7_L() int { return cpu.cbOpBIT_R8(&cpu.Registers.L, 7, 8) }
func (cpu *CPU) cbBIT_7_HL() int { return cpu.cbOpBIT_HL(7, 16) }
func (cpu *CPU) cbBIT_7_A() int { return cpu.cbOpBIT_R8(&cpu.Registers.A, 7, 8) }

// RES Functions
func (cpu *CPU) cbRES_0_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(0), 8) }
func (cpu *CPU) cbRES_0_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(0), 8) }
func (cpu *CPU) cbRES_0_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(0), 8) }
func (cpu *CPU) cbRES_0_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(0), 8) }
func (cpu *CPU) cbRES_0_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(0), 8) }
func (cpu *CPU) cbRES_0_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(0), 8) }
func (cpu *CPU) cbRES_0_HL() int { return cpu.cbOpHL(resOp(0), 16) }
func (cpu *CPU) cbRES_0_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(0), 8) }

func (cpu *CPU) cbRES_1_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(1), 8) }
func (cpu *CPU) cbRES_1_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(1), 8) }
func (cpu *CPU) cbRES_1_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(1), 8) }
func (cpu *CPU) cbRES_1_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(1), 8) }
func (cpu *CPU) cbRES_1_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(1), 8) }
func (cpu *CPU) cbRES_1_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(1), 8) }
func (cpu *CPU) cbRES_1_HL() int { return cpu.cbOpHL(resOp(1), 16) }
func (cpu *CPU) cbRES_1_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(1), 8) }

func (cpu *CPU) cbRES_2_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(2), 8) }
func (cpu *CPU) cbRES_2_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(2), 8) }
func (cpu *CPU) cbRES_2_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(2), 8) }
func (cpu *CPU) cbRES_2_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(2), 8) }
func (cpu *CPU) cbRES_2_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(2), 8) }
func (cpu *CPU) cbRES_2_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(2), 8) }
func (cpu *CPU) cbRES_2_HL() int { return cpu.cbOpHL(resOp(2), 16) }
func (cpu *CPU) cbRES_2_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(2), 8) }

func (cpu *CPU) cbRES_3_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(3), 8) }
func (cpu *CPU) cbRES_3_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(3), 8) }
func (cpu *CPU) cbRES_3_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(3), 8) }
func (cpu *CPU) cbRES_3_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(3), 8) }
func (cpu *CPU) cbRES_3_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(3), 8) }
func (cpu *CPU) cbRES_3_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(3), 8) }
func (cpu *CPU) cbRES_3_HL() int { return cpu.cbOpHL(resOp(3), 16) }
func (cpu *CPU) cbRES_3_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(3), 8) }

func (cpu *CPU) cbRES_4_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(4), 8) }
func (cpu *CPU) cbRES_4_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(4), 8) }
func (cpu *CPU) cbRES_4_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(4), 8) }
func (cpu *CPU) cbRES_4_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(4), 8) }
func (cpu *CPU) cbRES_4_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(4), 8) }
func (cpu *CPU) cbRES_4_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(4), 8) }
func (cpu *CPU) cbRES_4_HL() int { return cpu.cbOpHL(resOp(4), 16) }
func (cpu *CPU) cbRES_4_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(4), 8) }

func (cpu *CPU) cbRES_5_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(5), 8) }
func (cpu *CPU) cbRES_5_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(5), 8) }
func (cpu *CPU) cbRES_5_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(5), 8) }
func (cpu *CPU) cbRES_5_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(5), 8) }
func (cpu *CPU) cbRES_5_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(5), 8) }
func (cpu *CPU) cbRES_5_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(5), 8) }
func (cpu *CPU) cbRES_5_HL() int { return cpu.cbOpHL(resOp(5), 16) }
func (cpu *CPU) cbRES_5_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(5), 8) }

func (cpu *CPU) cbRES_6_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(6), 8) }
func (cpu *CPU) cbRES_6_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(6), 8) }
func (cpu *CPU) cbRES_6_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(6), 8) }
func (cpu *CPU) cbRES_6_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(6), 8) }
func (cpu *CPU) cbRES_6_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(6), 8) }
func (cpu *CPU) cbRES_6_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(6), 8) }
func (cpu *CPU) cbRES_6_HL() int { return cpu.cbOpHL(resOp(6), 16) }
func (cpu *CPU) cbRES_6_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(6), 8) }

func (cpu *CPU) cbRES_7_B() int { return cpu.cbOpR8(&cpu.Registers.B, resOp(7), 8) }
func (cpu *CPU) cbRES_7_C() int { return cpu.cbOpR8(&cpu.Registers.C, resOp(7), 8) }
func (cpu *CPU) cbRES_7_D() int { return cpu.cbOpR8(&cpu.Registers.D, resOp(7), 8) }
func (cpu *CPU) cbRES_7_E() int { return cpu.cbOpR8(&cpu.Registers.E, resOp(7), 8) }
func (cpu *CPU) cbRES_7_H() int { return cpu.cbOpR8(&cpu.Registers.H, resOp(7), 8) }
func (cpu *CPU) cbRES_7_L() int { return cpu.cbOpR8(&cpu.Registers.L, resOp(7), 8) }
func (cpu *CPU) cbRES_7_HL() int { return cpu.cbOpHL(resOp(7), 16) }
func (cpu *CPU) cbRES_7_A() int { return cpu.cbOpR8(&cpu.Registers.A, resOp(7), 8) }

// SET Functions
func (cpu *CPU) cbSET_0_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(0), 8) }
func (cpu *CPU) cbSET_0_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(0), 8) }
func (cpu *CPU) cbSET_0_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(0), 8) }
func (cpu *CPU) cbSET_0_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(0), 8) }
func (cpu *CPU) cbSET_0_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(0), 8) }
func (cpu *CPU) cbSET_0_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(0), 8) }
func (cpu *CPU) cbSET_0_HL() int { return cpu.cbOpHL(setOp(0), 16) }
func (cpu *CPU) cbSET_0_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(0), 8) }

func (cpu *CPU) cbSET_1_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(1), 8) }
func (cpu *CPU) cbSET_1_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(1), 8) }
func (cpu *CPU) cbSET_1_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(1), 8) }
func (cpu *CPU) cbSET_1_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(1), 8) }
func (cpu *CPU) cbSET_1_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(1), 8) }
func (cpu *CPU) cbSET_1_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(1), 8) }
func (cpu *CPU) cbSET_1_HL() int { return cpu.cbOpHL(setOp(1), 16) }
func (cpu *CPU) cbSET_1_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(1), 8) }

func (cpu *CPU) cbSET_2_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(2), 8) }
func (cpu *CPU) cbSET_2_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(2), 8) }
func (cpu *CPU) cbSET_2_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(2), 8) }
func (cpu *CPU) cbSET_2_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(2), 8) }
func (cpu *CPU) cbSET_2_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(2), 8) }
func (cpu *CPU) cbSET_2_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(2), 8) }
func (cpu *CPU) cbSET_2_HL() int { return cpu.cbOpHL(setOp(2), 16) }
func (cpu *CPU) cbSET_2_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(2), 8) }

func (cpu *CPU) cbSET_3_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(3), 8) }
func (cpu *CPU) cbSET_3_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(3), 8) }
func (cpu *CPU) cbSET_3_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(3), 8) }
func (cpu *CPU) cbSET_3_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(3), 8) }
func (cpu *CPU) cbSET_3_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(3), 8) }
func (cpu *CPU) cbSET_3_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(3), 8) }
func (cpu *CPU) cbSET_3_HL() int { return cpu.cbOpHL(setOp(3), 16) }
func (cpu *CPU) cbSET_3_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(3), 8) }

func (cpu *CPU) cbSET_4_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(4), 8) }
func (cpu *CPU) cbSET_4_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(4), 8) }
func (cpu *CPU) cbSET_4_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(4), 8) }
func (cpu *CPU) cbSET_4_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(4), 8) }
func (cpu *CPU) cbSET_4_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(4), 8) }
func (cpu *CPU) cbSET_4_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(4), 8) }
func (cpu *CPU) cbSET_4_HL() int { return cpu.cbOpHL(setOp(4), 16) }
func (cpu *CPU) cbSET_4_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(4), 8) }

func (cpu *CPU) cbSET_5_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(5), 8) }
func (cpu *CPU) cbSET_5_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(5), 8) }
func (cpu *CPU) cbSET_5_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(5), 8) }
func (cpu *CPU) cbSET_5_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(5), 8) }
func (cpu *CPU) cbSET_5_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(5), 8) }
func (cpu *CPU) cbSET_5_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(5), 8) }
func (cpu *CPU) cbSET_5_HL() int { return cpu.cbOpHL(setOp(5), 16) }
func (cpu *CPU) cbSET_5_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(5), 8) }

func (cpu *CPU) cbSET_6_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(6), 8) }
func (cpu *CPU) cbSET_6_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(6), 8) }
func (cpu *CPU) cbSET_6_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(6), 8) }
func (cpu *CPU) cbSET_6_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(6), 8) }
func (cpu *CPU) cbSET_6_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(6), 8) }
func (cpu *CPU) cbSET_6_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(6), 8) }
func (cpu *CPU) cbSET_6_HL() int { return cpu.cbOpHL(setOp(6), 16) }
func (cpu *CPU) cbSET_6_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(6), 8) }

func (cpu *CPU) cbSET_7_B() int { return cpu.cbOpR8(&cpu.Registers.B, setOp(7), 8) }
func (cpu *CPU) cbSET_7_C() int { return cpu.cbOpR8(&cpu.Registers.C, setOp(7), 8) }
func (cpu *CPU) cbSET_7_D() int { return cpu.cbOpR8(&cpu.Registers.D, setOp(7), 8) }
func (cpu *CPU) cbSET_7_E() int { return cpu.cbOpR8(&cpu.Registers.E, setOp(7), 8) }
func (cpu *CPU) cbSET_7_H() int { return cpu.cbOpR8(&cpu.Registers.H, setOp(7), 8) }
func (cpu *CPU) cbSET_7_L() int { return cpu.cbOpR8(&cpu.Registers.L, setOp(7), 8) }
func (cpu *CPU) cbSET_7_HL() int { return cpu.cbOpHL(setOp(7), 16) }
func (cpu *CPU) cbSET_7_A() int { return cpu.cbOpR8(&cpu.Registers.A, setOp(7), 8) }
