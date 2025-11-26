package timer

// timer register addresses
const (
	DIV  = 0xFF04
	TIMA = 0xFF05
	TMA  = 0xFF06
	TAC  = 0xFF07
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
