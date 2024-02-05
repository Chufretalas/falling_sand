package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

var ()

const (
	SCREENWIDTH  = 1920
	SCREENHEIGHT = 1080
	SQUARESIDE   = 10
)

type Game struct{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		go ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("ahahaha")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mouse_x, mouse_y := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nMouse pos: x%v y%v", ebiten.ActualFPS(), mouse_x, mouse_y))

	for y := 0; y <= SCREENHEIGHT; y += SQUARESIDE {
		vector.StrokeLine(screen, 0, float32(y), SCREENWIDTH, float32(y), 1, color.RGBA{20, 20, 20, 10}, true)
	}
	for x := 0; x <= SCREENWIDTH; x += SQUARESIDE {
		vector.StrokeLine(screen, float32(x), 0, float32(x), SCREENHEIGHT, 1, color.RGBA{20, 20, 20, 10}, true)
	}

	vector.DrawFilledRect(screen, float32(mouse_x)-float32(mouse_x%SQUARESIDE), float32(mouse_y)-float32(mouse_y%SQUARESIDE), SQUARESIDE, SQUARESIDE, color.RGBA{255, 0, 0, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
