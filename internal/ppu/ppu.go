package ppu

const (
	// Register Addresses (for reference)
	LCDC_AD = 0xFF40
	STAT_AD = 0xFF41
	SCY_AD  = 0xFF42
	SCX_AD  = 0xFF43
	LYC_AD  = 0xFF45
	BGP_AD  = 0xFF47

	// PPU Modes
	MODE_HBLANK = 0
	MODE_VBLANK = 1
	MODE_OAM    = 2
	MODE_XFER   = 3

	// LCDC Bitmasks
	LCDC_BG_ENABLE   = 0x01
	LCDC_OBJ_ENABLE  = 0x02
	LCDC_OBJ_SIZE    = 0x04
	LCDC_BG_MAP      = 0x08
	LCDC_BG_WIN_DATA = 0x10
	LCDC_WIN_ENABLE  = 0x20
	LCDC_WIN_MAP     = 0x40
	LCDC_ENABLE      = 0x80

	// STAT Bitmasks
	STAT_LYC_INT    = 0x40 // Bit 6
	STAT_OAM_INT    = 0x20 // Bit 5
	STAT_VBLANK_INT = 0x10 // Bit 4
	STAT_HBLANK_INT = 0x08 // Bit 3
	STAT_LYC_EQUAL  = 0x04 // Bit 2
)

const (
	LINE_CYCLE = 456
	H_LEN      = 143
	LY_AD      = 0xFF44
	V_BLANK    = 144
)

var Palette = [4][4]uint8{
	{0xE0, 0xF8, 0xD0, 0xFF}, // Color 0 (White-ish)
	{0x88, 0xC0, 0x70, 0xFF}, // Color 1 (Light Gray)
	{0x34, 0x68, 0x56, 0xFF}, // Color 2 (Dark Gray)
	{0x08, 0x18, 0x20, 0xFF}, // Color 3 (Black-ish)
}

type MemoryBus interface {
	Read(addr uint16) uint8
}

type PPU struct {
	// Connection to the rest of the system
	Bus MemoryBus

	// Registers
	Lcdc uint8
	Stat uint8
	Scx  uint8
	Scy  uint8
	Wx   uint8
	Wy   uint8
	Ly   uint8
	Lyc  uint8
	Bgp  uint8 // Background Palette (0xFF47)

	// Internal State
	Mode          uint8 // 0, 1, 2, 3
	ModeClock     int   // Cycle accumulator
	StatIrqSignal bool  // Edge detection for STAT interrupts

	// Output Buffer (160x144 pixels, 4 bytes each: RGBA)
	VideoBuffer [160 * 144 * 4]byte

	// Interrupt Callback (Linked in main.go)
	RequestInterrupt func(bit uint8)
}

func NewPPU() *PPU {
	return &PPU{
		Mode: MODE_OAM, // Start in Mode 2 usually
	}
}

// ---------------------------------------------------------
// REGISTER ACCESS
// ---------------------------------------------------------

func (p *PPU) WriteLCDC(val uint8) {
	p.Lcdc = val

	// If LCD is turned off, reset state immediately
	if (p.Lcdc & LCDC_ENABLE) == 0 {
		p.Mode = MODE_HBLANK
		p.Ly = 0
		p.ModeClock = 0
		p.Stat &= 0xFC // Clear mode bits
	}
}

// ---------------------------------------------------------
// CORE TIMING LOGIC (The Tick)
// ---------------------------------------------------------

func (p *PPU) Tick(cycles int) {
	if (p.Lcdc & LCDC_ENABLE) == 0 {
		return // LCD is off, do nothing
	}

	p.ModeClock += cycles

	switch p.Mode {

	// MODE 2: OAM SEARCH (80 Cycles)
	case MODE_OAM:
		if p.ModeClock >= 80 {
			p.ModeClock -= 80
			p.Mode = MODE_XFER
		}

	// MODE 3: PIXEL TRANSFER (172 Cycles)
	case MODE_XFER:
		if p.ModeClock >= 172 {
			p.ModeClock -= 172
			p.Mode = MODE_HBLANK

			// End of Mode 3: Draw the line!
			p.RenderScanline()
		}

	// MODE 0: H-BLANK (204 Cycles)
	case MODE_HBLANK:
		if p.ModeClock >= 204 {
			p.ModeClock -= 204
			p.Ly++

			// Check Coincidence Flag
			p.compareLYC()

			if p.Ly == 144 {
				p.Mode = MODE_VBLANK
				// Fire V-Blank Interrupt (Bit 0)
				if p.RequestInterrupt != nil {
					p.RequestInterrupt(0)
				}
			} else {
				p.Mode = MODE_OAM
			}
		}

	// MODE 1: V-BLANK (456 Cycles per line)
	case MODE_VBLANK:
		if p.ModeClock >= 456 {
			p.ModeClock -= 456
			p.Ly++
			p.compareLYC()

			if p.Ly > 153 {
				// Restart Frame
				p.Mode = MODE_OAM
				p.Ly = 0
				p.compareLYC()
			}
		}
	}

	// Update STAT register bits and check for STAT interrupts
	p.updateStat()
}

// ---------------------------------------------------------
// HELPER LOGIC
// ---------------------------------------------------------

// compareLYC checks if LY == LYC and updates the STAT register
func (p *PPU) compareLYC() {
	if p.Ly == p.Lyc {
		p.Stat |= STAT_LYC_EQUAL
	} else {
		p.Stat &= ^uint8(STAT_LYC_EQUAL)
	}
}

// updateStat updates the Mode bits in STAT and fires interrupts
func (p *PPU) updateStat() {
	// 1. Update Mode Bits (0-1)
	p.Stat &= 0xFC
	p.Stat |= (p.Mode & 0x03)

	// 2. Check Interrupt Signal
	currentSignal := false

	// LY=LYC Interrupt
	if (p.Stat&STAT_LYC_EQUAL) != 0 && (p.Stat&STAT_LYC_INT) != 0 {
		currentSignal = true
	}
	// Mode 2 Interrupt
	if p.Mode == MODE_OAM && (p.Stat&STAT_OAM_INT) != 0 {
		currentSignal = true
	}
	// Mode 1 Interrupt
	if p.Mode == MODE_VBLANK && (p.Stat&STAT_VBLANK_INT) != 0 {
		currentSignal = true
	}
	// Mode 0 Interrupt
	if p.Mode == MODE_HBLANK && (p.Stat&STAT_HBLANK_INT) != 0 {
		currentSignal = true
	}

	// 3. Fire on Rising Edge
	if currentSignal && !p.StatIrqSignal {
		if p.RequestInterrupt != nil {
			p.RequestInterrupt(1) // STAT Interrupt (Bit 1)
		}
	}
	p.StatIrqSignal = currentSignal
}

// ---------------------------------------------------------
// RENDERING LOGIC
// ---------------------------------------------------------

func (p *PPU) RenderScanline() {
	if p.Bus == nil {
		return // Safety check
	}

	// 1. Read Control Registers via Interface
	scy := p.Bus.Read(SCY_AD)
	scx := p.Bus.Read(SCX_AD)
	// For Week 1, we ignore BGP (Palette Reg) and use the hardcoded one

	// 2. Determine Map & Tile Data Areas
	mapBase := uint16(0x9800)
	if (p.Lcdc & LCDC_BG_MAP) != 0 {
		mapBase = 0x9C00
	}

	unsignedAddressing := (p.Lcdc & LCDC_BG_WIN_DATA) != 0

	// 3. Loop through 160 pixels
	yPos := uint16(p.Ly) + uint16(scy) // Implicit % 256 for 8-bit math, but safe to cast
	mapY := uint8(yPos)                // Wrap 256

	for x := range [160]int{} {
		xPos := uint8(x) + scx // Wrap 256 automatically
		mapX := xPos

		// Which tile is this?
		tileRow := uint16(mapY / 8)
		tileCol := uint16(mapX / 8)

		tileAddr := mapBase + (tileRow * 32) + tileCol
		tileID := p.Bus.Read(tileAddr)

		// Where is the pixel data?
		var dataAddr uint16
		if unsignedAddressing {
			dataAddr = 0x8000 + (uint16(tileID) * 16)
		} else {
			// Signed addressing: 0x9000 + (int8(id) * 16)
			signedID := int16(int8(tileID))
			dataAddr = uint16(0x9000 + (int32(signedID) * 16))
		}

		// Which row of the tile?
		lineInTile := uint16(mapY % 8)
		dataAddr += (lineInTile * 2)

		// Read the two bytes
		byte1 := p.Bus.Read(dataAddr)
		byte2 := p.Bus.Read(dataAddr + 1)

		// Decode bit (7-0)
		bitIndex := 7 - (mapX % 8)
		lo := (byte1 >> bitIndex) & 0x01
		hi := (byte2 >> bitIndex) & 0x01
		colorIndex := (hi << 1) | lo

		// Get Color
		color := Palette[colorIndex]

		// Write to Ebiten Buffer (RGBA)
		pixelIndex := (int(p.Ly)*160 + x) * 4
		p.VideoBuffer[pixelIndex] = color[0]
		p.VideoBuffer[pixelIndex+1] = color[1]
		p.VideoBuffer[pixelIndex+2] = color[2]
		p.VideoBuffer[pixelIndex+3] = color[3]
	}
}
