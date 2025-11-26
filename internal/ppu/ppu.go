package ppu

import "fmt"

// LCDC Flags (Bit positions)
const (
	LCDC_BG_ENABLE   = 0x01 // Bit 0
	LCDC_OBJ_ENABLE  = 0x02 // Bit 1
	LCDC_OBJ_SIZE    = 0x04 // Bit 2
	LCDC_BG_MAP      = 0x08 // Bit 3
	LCDC_BG_WIN_DATA = 0x10 // Bit 4
	LCDC_WIN_ENABLE  = 0x20 // Bit 5
	LCDC_WIN_MAP     = 0x40 // Bit 6
	LCDC_ENABLE      = 0x80 // Bit 7
)

type PPU struct {

	// The Control Register
	Lcdc       uint8
	dotCounter int
	LY         uint8
}

const (
	LINE_CYCLE = 456
	H_LEN      = 143
	LY_AD      = 0xFF44
	V_BLANK    = 144
)

func NewPPU() *PPU {
	return &PPU{}
}

func (p *PPU) ReadLCDC() uint8 {
	return p.Lcdc
}

func (p *PPU) WriteLCDC(val uint8) {
	p.Lcdc = val
}

type Tile [8][8]uint8

func DecodeTile(data []byte) Tile {
	var tile Tile

	// A tile has 8 rows (height)
	for y := range [8]int{} {
		lowByte := data[y*2]
		highByte := data[y*2+1]

		// A row has 8 pixels (width)
		for x := range [8]int{} {
			bitIndex := 7 - x

			// 1. Extract the bit from the Low Byte
			// Shift right by bitIndex to put the target bit at position 0, then AND with 1.
			lowBit := (lowByte >> bitIndex) & 0x01

			// 2. Extract the bit from the High Byte
			highBit := (highByte >> bitIndex) & 0x01

			// 3. Combine them to get the Color ID (0-3)
			// The High Bit is worth 2, the Low Bit is worth 1.
			colorID := (highBit << 1) | lowBit

			tile[y][x] = colorID
		}
	}

	return tile
}

func (ppu *PPU) Tick(cycles uint8) {
	ppu.dotCounter += int(cycles)
	if ppu.dotCounter > LINE_CYCLE {
		ppu.dotCounter -= LINE_CYCLE
		ppu.LY++
		if ppu.LY == V_BLANK {
			fmt.Println("---VBLANK----")
		}
		if ppu.LY > H_LEN {
			ppu.LY = 0
		}
	}
}
