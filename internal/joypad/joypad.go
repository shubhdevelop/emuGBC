package joypad

const (
	// Memory Address
	P1 = 0xFF00

	JOYPAD_INT = 4
)

type Joypad struct {
	ActionState    uint8
	DirectionState uint8

	SelectAction    bool
	SelectDirection bool

	RequestInterrupt func(bit uint8)
	
	// Track if this is the first update to prevent false interrupts
	firstUpdate bool
}

func NewJoypad() *Joypad {
	return &Joypad{
		ActionState:     0x0F,
		DirectionState:  0x0F,
		SelectAction:    true,
		SelectDirection: true,
		firstUpdate:     true,
	}
}

// Write updates the selection bits (Bits 4 and 5)
func (j *Joypad) Write(val uint8) {
	// Bits 4 and 5 are the selectors.
	// 0 means "Select this row".
	j.SelectDirection = (val & 0x10) != 0 // Check Bit 4
	j.SelectAction = (val & 0x20) != 0    // Check Bit 5
}

func (j *Joypad) Read() uint8 {
	// Start with 0xFF (All 1s)
	// Bits 6 and 7 are always 1 (Unused)
	// Bit 5 and 4 come from our selection state
	res := uint8(0xCF)

	if !j.SelectAction {
		res &= ^uint8(0x20)  // Clear Bit 5 (Set to 0) to show it's active
		res &= j.ActionState // Combine with Action Buttons (A, B, etc)
	}

	if !j.SelectDirection {
		res &= ^uint8(0x10)     // Clear Bit 4
		res &= j.DirectionState // Combine with D-Pad
	}

	return res
}

// UpdateState sets the button status based on keyboard input
// inputs: [Right, Left, Up, Down, A, B, Select, Start] (bools)
func (j *Joypad) UpdateState(right, left, up, down, a, b, sel, start bool) {

	// Remember old state to detect changes (Edges)
	oldState := j.Read()

	// Reset internal state to 0x0F (Released)
	j.DirectionState = 0x0F
	j.ActionState = 0x0F

	// Directions (Bit 0=Right, 1=Left, 2=Up, 3=Down)
	if right {
		j.DirectionState &= ^uint8(0x01)
	}
	if left {
		j.DirectionState &= ^uint8(0x02)
	}
	if up {
		j.DirectionState &= ^uint8(0x04)
	}
	if down {
		j.DirectionState &= ^uint8(0x08)
	}

	// Actions (Bit 0=A, 1=B, 2=Select, 3=Start)
	if a {
		j.ActionState &= ^uint8(0x01)
	}
	if b {
		j.ActionState &= ^uint8(0x02)
	}
	if sel {
		j.ActionState &= ^uint8(0x04)
	}
	if start {
		j.ActionState &= ^uint8(0x08)
	}

	// Interrupt Logic: Did any line go from 1 -> 0?
	newState := j.Read()

	// Skip interrupt on first update to prevent false triggers
	if j.firstUpdate {
		j.firstUpdate = false
		return
	}

	// Check if any bit turned to 0 (Falling Edge)
	// (old & ^new) finds bits that were 1 and are now 0
	if (oldState & ^newState) != 0 {
		if j.RequestInterrupt != nil {
			j.RequestInterrupt(JOYPAD_INT)
		}
	}
}
