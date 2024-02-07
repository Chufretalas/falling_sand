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

type Game struct{}

type Position struct {
	x, y float32
}

func blockPos2BlocksIdx(p Position) (int, int) {
	return int(p.y / float32(squareSide)), int(p.x / float32(squareSide))
}

var (
	blocks             BlockGrid
	blocksCopy         BlockGrid
	activeBlock        Position
	updateDelayCounter int
	updateDelayMax     int
	squareSide         int // possible values: 1, 2, 3, 4, 5, 6, 10, 20
	cSize              int // cursorSize, default is 0
)

const (
	SCREENWIDTH  = 1920
	SCREENHEIGHT = 1080
)

// ---------------------------------------- START EBITENGINE FUNCTIONS ---------------------------------------- //

func init() {
	activeBlock = Position{0, 0}
	updateDelayMax = 8
	updateDelayCounter = 0
	squareSide = 20
	cSize = 0
	blocks.init()
	blocksCopy.init()
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		go ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("ahahaha")
	}

	// Putting and removing sand

	mouseX, mouseY := ebiten.CursorPosition()

	activeBlock.x = float32(mouseX) - float32(mouseX%squareSide)
	activeBlock.y = float32(mouseY) - float32(mouseY%squareSide)

	blocksIdx1, blocksIdx2 := blockPos2BlocksIdx(activeBlock)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		for iy := blocksIdx1 - cSize; iy <= blocksIdx1+cSize; iy++ {
			for ix := blocksIdx2 - cSize; ix <= blocksIdx2+cSize; ix++ {
				if ix >= 0 && ix < len(blocks[0]) && iy >= 0 && iy < len(blocks) {
					blocks[iy][ix] = BTSAND
				}
			}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		for iy := blocksIdx1 - cSize; iy <= blocksIdx1+cSize; iy++ {
			for ix := blocksIdx2 - cSize; ix <= blocksIdx2+cSize; ix++ {
				if ix >= 0 && ix < len(blocks[0]) && iy >= 0 && iy < len(blocks) {
					blocks[iy][ix] = BTAIR
				}
			}
		}
	}

	// End - Putting and removing sand

	// Chance cursor size
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && cSize > 0 {
		cSize--
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && cSize < 15 {
		cSize++
	}
	// End Chance cursor size

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		blocks.clear()
	}

	updateDelayCounter++
	if updateDelayCounter >= updateDelayMax {
		updateblocks()
		updateDelayCounter = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y <= SCREENHEIGHT; y += squareSide {
		vector.StrokeLine(screen, 0, float32(y), SCREENWIDTH, float32(y), 1, color.RGBA{20, 20, 20, 10}, true)
	}
	for x := 0; x <= SCREENWIDTH; x += squareSide {
		vector.StrokeLine(screen, float32(x), 0, float32(x), SCREENHEIGHT, 1, color.RGBA{20, 20, 20, 10}, true)
	}

	for idx1 := range blocks {
		for idx2, block := range blocks[idx1] {
			if block != BTAIR {
				vector.DrawFilledRect(screen, float32(idx2*squareSide), float32(idx1*squareSide), float32(squareSide), float32(squareSide), color.RGBA{200, 100, 100, 255}, true)
			}

		}
	}

	vector.StrokeRect(screen, activeBlock.x-float32(cSize*squareSide), activeBlock.y-float32(cSize*squareSide), float32(squareSide*(cSize*2+1)), float32(squareSide*(cSize*2+1)), 2, color.RGBA{255, 255, 255, 255}, true)

	mouseX, mouseY := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nTPS: %v\nMouse: x%v y%v", ebiten.ActualFPS(), ebiten.ActualTPS(), mouseX, mouseY))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Falling Sand")
	ebiten.SetTPS(240)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
