package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shubhdevelop/emuGBC/internal/cpu"
	"github.com/shubhdevelop/emuGBC/internal/display"
	"github.com/shubhdevelop/emuGBC/internal/joypad"
	"github.com/shubhdevelop/emuGBC/internal/mmu"
	"github.com/shubhdevelop/emuGBC/internal/ppu"
	"github.com/shubhdevelop/emuGBC/internal/timer"
)

func main() {
	args := os.Args
	fmt.Println("Powering on Game Boy...")
	MMU := mmu.NewMMU()
	PPU := ppu.NewPPU()
	PPU.RequestInterrupt = func(bit uint8) {
		// Read current Interrupt Flag (IF) register (0xFF0F)
		currentIF := MMU.Read(0xFF0F)

		// Set the specific bit (Bit 0 for V-Blank, Bit 1 for STAT)
		newIF := currentIF | (1 << bit)

		// Write it back
		MMU.Write(0xFF0F, newIF)
	}

	JOY := joypad.NewJoypad()
	JOY.RequestInterrupt = func(bit uint8) {
		val := MMU.Read(0xFF0F)
		MMU.Write(0xFF0F, val|(1<<bit))
	}

	Timer := timer.NewTimer()
	// Do the same for Timer if you haven't yet!
	Timer.RequestInterrupt = func(bit uint8) {
		currentIF := MMU.Read(0xFF0F)
		MMU.Write(0xFF0F, currentIF|(1<<bit))
	}
	MMU.PPU = PPU
	PPU.Bus = MMU
	MMU.Timer = Timer
	MMU.Joypad = JOY

	cpu := cpu.NewCPU(MMU)
	var romPath string

	if len(args) > 2 && args[2] == "bin" {
		cpu.LogMode = "bin"
	}
	if len(args) > 1 {
		romPath = args[1]
	}

	romData, err := os.ReadFile(romPath)
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
