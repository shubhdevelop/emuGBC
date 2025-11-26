package cpu

/*
 * GBC CPU doesn't use the V flag for the overflow
 */
const (
	FlagZ = 1 << 7 // Zero
	FlagN = 1 << 6 // Subtract
	FlagH = 1 << 5 // Half Carry
	FlagC = 1 << 4 // Carry
)

func (r *Registers) GetFlag(flag uint8) bool {
	return (r.F & flag) != 0
}

func (r *Registers) SetFlag(flag uint8, value bool) {
	if value {
		r.F |= flag
	} else {
		r.F &^= flag
	}
}

func halfCarrySub(a, b byte) bool {
	return (a & 0x0F) < (b & 0x0F)
}

func halfCarryInc(a byte) bool {
	// If the bottom 4 bits are 1111 (15), adding 1 causes a carry to bit 4
	return (a & 0x0F) == 0x0F
}

func halfCarryDec(a byte) bool {
	// If the bottom 4 bits are 0000 (0), subtracting 1 causes a borrow from bit 4
	return (a & 0x0F) == 0x00
}

func carrySub(a, b byte) bool {
	return a < b
}

func halfCarryAdd(a, b byte) bool {
	return (a&0x0f)+(b&0x0f) > 0x0f
}

func carryAdd(a, b byte) bool {
	return uint16(a)+uint16(b) > 0xFF
}
