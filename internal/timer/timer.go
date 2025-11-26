package timer

// timer register addresses
const (
	DIV_AD  = 0xFF04
	TIMA_AD = 0xFF05
	TMA_AD  = 0xFF06
	TAC_AD  = 0xFF07
)

type Timer struct {
	Div  uint16 // We use a 16-bit int to track the internal high-speed counter
	Tima uint8
	Tma  uint8
	Tac  uint8

	RequestInterrupt func(bit uint8)
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Tick(cycles uint8) {
	t.Div += uint16(cycles)
}

// handles CPU reading from 0xFF04-0xFF07
func (t *Timer) Read(addr uint16) uint8 {
	switch addr {
	case DIV_AD:
		// DIV (0xFF04) is just the High Byte of the internal 16-bit counter
		return uint8(t.Div >> 8)

	case TIMA_AD:
		return t.Tima

	case TMA_AD:
		return t.Tma

	case TAC_AD:
		// Bit 0-2 are used. Bits 3-7 are unused and usually return 1s (0xF8)
		return t.Tac | 0xF8
	}

	return 0xFF // Fallback
}

// handles CPU writing to 0xFF04-0xFF07
func (t *Timer) Write(addr uint16, val uint8) {
	switch addr {
	case DIV_AD:

		t.Div = 0 // Writing ANY value to 0xFF04 resets the whole internal counter to 0.

	case TIMA_AD:
		t.Tima = val

	case TMA_AD:
		t.Tma = val

	case TAC_AD:
		t.Tac = val
	}
}
