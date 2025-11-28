package display

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shubhdevelop/emuGBC/internal/cpu"
	"github.com/shubhdevelop/emuGBC/internal/mmu"
	"github.com/shubhdevelop/emuGBC/internal/ppu"
	"github.com/shubhdevelop/emuGBC/internal/timer"
)

type Game struct {
	CPU   *cpu.CPU
	PPU   *ppu.PPU
	Timer *timer.Timer
	MMU   *mmu.MMU
}

// Update is called 60 times a second by Ebiten
func (g *Game) Update() error {
	// We need to run enough CPU cycles to finish ONE FRAME (approx 70224 cycles)
	// before we return, so Ebiten can draw the result.

	cyclesPerFrame := 70224
	cyclesRun := 0

	right := ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	left := ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	up := ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	down := ebiten.IsKeyPressed(ebiten.KeyArrowDown)

	a := ebiten.IsKeyPressed(ebiten.KeyZ)         // Z = A
	b := ebiten.IsKeyPressed(ebiten.KeyX)         // X = B
	sel := ebiten.IsKeyPressed(ebiten.KeySpace)   // Space = Select
	start := ebiten.IsKeyPressed(ebiten.KeyEnter) // Enter = Start
	if right || left || up || down || a || b || sel || start {
		fmt.Printf("right: %t, left: %t, up: %t, down: %t, a: %t , b: %t , sel: %t, start: %t \n", right, left, up, down, a, b, sel, start)
	}

	// Push to Emulator
	g.MMU.Joypad.UpdateState(right, left, up, down, a, b, sel, start)
	for cyclesRun < cyclesPerFrame {
		// 1. Step CPU
		cycles := g.CPU.Step()
		// 2. Step PPU & Timer
		g.PPU.Tick(int(cycles))
		g.Timer.Tick(cycles)

		cyclesRun += int(cycles)

		// Safety: If CPU halts (returns 0 or 4 constantly), loop continues properly
	}

	return nil
}

// Draw is called every frame to put pixels on screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Write the PPU's buffer to the Ebiten screen
	screen.WritePixels(g.PPU.VideoBuffer[:])
}

// Layout defines the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}
