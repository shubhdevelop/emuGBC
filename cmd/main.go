package main

import (
	"fmt"
	"os"

	"github.com/shubhdevelop/emuGBC/internal/cpu"
	"github.com/shubhdevelop/emuGBC/internal/mmu"
)

func main() {
	fmt.Println("Powering on Game Boy...")
	memoryBus := mmu.NewMMU()
	cpu := cpu.NewCPU(memoryBus)

	romData, err := os.ReadFile("roms/tetris.gb")
	if err != nil {
		panic(fmt.Sprintf("Failed to load ROM: %v", err))
	}

	cpu.Bus.LoadROM(romData)
	fmt.Println("Starting execution...")
	fmt.Println("Setting up Registers...")
	cpu.Registers.PC = 0x0100
	cpu.Registers.SP = 0xFFFE
	cpu.Registers.A = 0x01
	cpu.Registers.F = 0xB0
	cpu.Registers.SetBC(0x0013)
	cpu.Registers.SetDE(0x00D8)
	cpu.Registers.SetHL(0x014D)
	fmt.Println("Registers Setup done")
	for {
		cycles := cpu.Step()
		_ = cycles
	}
}
