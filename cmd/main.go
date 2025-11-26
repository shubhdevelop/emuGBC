package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shubhdevelop/emuGBC/internal/cpu"
	"github.com/shubhdevelop/emuGBC/internal/display"
	"github.com/shubhdevelop/emuGBC/internal/mmu"
	"github.com/shubhdevelop/emuGBC/internal/ppu"
	"github.com/shubhdevelop/emuGBC/internal/timer"
)

func main() {
	args := os.Args
	fmt.Println("Powering on Game Boy...")
	MMU := mmu.NewMMU()
	PPU := ppu.NewPPU()
	Timer := timer.NewTimer()
	MMU.PPU = PPU
	PPU.Bus = MMU
	MMU.Timer = Timer

	cpu := cpu.NewCPU(MMU)

	if len(args) > 1 && args[1] == "bin" {
		cpu.LogMode = "bin"
	}

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

	game := &display.Game{
		CPU:   cpu,
		PPU:   PPU,
		Timer: Timer,
		MMU:   MMU,
	}

	// 4. Start the Window!
	ebiten.SetWindowSize(160*3, 144*3) // 3x scale so it's not tiny
	ebiten.SetWindowTitle("EmuGBC - Tetris")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
