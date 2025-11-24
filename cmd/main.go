package main

import (
	"fmt"

	"github.com/shubhdevelop/emuGBC/internal/cpu"
	"github.com/shubhdevelop/emuGBC/internal/mmu"
	"github.com/shubhdevelop/emuGBC/internal/ppu"
)

func main() {
	fmt.Println("Powering on Game Boy...")
	memoryBus := mmu.NewMMU()
	cpu := cpu.NewCPU(memoryBus)

	romData := []byte{
		0x00,
		0x01,
		0x02,
		0x06,
		0x08,
		0x0A,
		0x0E,
		0x11,
		0x12,
		0x16,
		0x1A,
		0x1E,
	}

	cpu.Bus.LoadROM(romData)
	fmt.Println("Starting execution...")

	for range [10]int{} {
		cpu.Step()
	}

	tileData := []byte{
		0x3C, 0x7E, // Row 0
		0x42, 0x42, // Row 1
		0x42, 0x42, // Row 2
		0x42, 0x42, // Row 3
		0x7E, 0x7E, // Row 4
		0x42, 0x42, // Row 5
		0x42, 0x42, // Row 6
		0x00, 0x00, // Row 7
		0x3C, 0x7E, // Row 0
		0x42, 0x42, // Row 1
		0x42, 0x42, // Row 2
		0x42, 0x42, // Row 3
		0x7E, 0x7E, // Row 4
		0x42, 0x42, // Row 5
		0x42, 0x42, // Row 6
		0x00, 0x00, // Row 7
	}

	// Decode it!
	tile := ppu.DecodeTile(tileData)
	fmt.Println("Decoded Tile (ASCII View):")
	for y := range [8]int{} {
		for x := range [8]int{} {
			id := tile[y][x]

			// Map IDs to characters so we can "see" the image
			switch id {
			case 0:
				fmt.Print(". ") // White
			case 1:
				fmt.Print("- ") // Light Gray
			case 2:
				fmt.Print("+ ") // Dark Gray
			case 3:
				fmt.Print("# ") // Black
			}
		}
		fmt.Println() // New line at end of row
	}
}
