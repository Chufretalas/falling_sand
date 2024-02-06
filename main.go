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

type BlockType int

const (
	BTAIR BlockType = iota
	BTSAND
)

type BlockGrid [][]BlockType

func (b *BlockGrid) clear() {
	for idx1 := range *b {
		for idx2 := range (*b)[idx1] {
			(*b)[idx1][idx2] = BTAIR
		}
	}
}

var (
	blocks             BlockGrid
	blocksCopy         BlockGrid
	activeBlock        Position
	updateDelayCounter int
)

const (
	SCREENWIDTH  = 1920
	SCREENHEIGHT = 1080
	SQUARESIDE   = 20
)

func init() {
	blocks = make(BlockGrid, SCREENHEIGHT/SQUARESIDE)
	blocksCopy = make(BlockGrid, SCREENHEIGHT/SQUARESIDE)
	for idx := range blocks {
		blocks[idx] = make([]BlockType, SCREENWIDTH/SQUARESIDE)
		blocksCopy[idx] = make([]BlockType, SCREENWIDTH/SQUARESIDE)
	}

	activeBlock = Position{0, 0}
	updateDelayCounter = 0
}

func blockPos2BlocksIdx(p Position) (int, int) {
	return int(p.y / SQUARESIDE), int(p.x / SQUARESIDE)
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		go ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("ahahaha")
	}

	mouseX, mouseY := ebiten.CursorPosition()

	if mouseX >= 0 && mouseX <= SCREENWIDTH && mouseY >= 0 && mouseY <= SCREENHEIGHT {
		activeBlock.x = float32(mouseX) - float32(mouseX%SQUARESIDE)
		activeBlock.y = float32(mouseY) - float32(mouseY%SQUARESIDE)

		blocksIdx1, blocksIdx2 := blockPos2BlocksIdx(activeBlock)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			blocks[blocksIdx1][blocksIdx2] = BTSAND
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
			blocks[blocksIdx1][blocksIdx2] = BTAIR
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		blocks.clear()
	}

	updateDelayCounter++
	if updateDelayCounter >= 4 {
		updateblocks()
		updateDelayCounter = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nTPS: %v\nMouse: x%v y%v", ebiten.ActualFPS(), ebiten.ActualTPS(), mouseX, mouseY))

	for y := 0; y <= SCREENHEIGHT; y += SQUARESIDE {
		vector.StrokeLine(screen, 0, float32(y), SCREENWIDTH, float32(y), 1, color.RGBA{20, 20, 20, 10}, true)
	}
	for x := 0; x <= SCREENWIDTH; x += SQUARESIDE {
		vector.StrokeLine(screen, float32(x), 0, float32(x), SCREENHEIGHT, 1, color.RGBA{20, 20, 20, 10}, true)
	}

	for idx1 := range blocks {
		for idx2, block := range blocks[idx1] {
			if block != BTAIR {
				vector.DrawFilledRect(screen, float32(idx2*SQUARESIDE), float32(idx1*SQUARESIDE), SQUARESIDE, SQUARESIDE, color.RGBA{200, 100, 100, 255}, true)
			}

		}
	}

	vector.DrawFilledRect(screen, activeBlock.x, activeBlock.y, SQUARESIDE, SQUARESIDE, color.RGBA{255, 0, 0, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Falling Sand")
	ebiten.SetTPS(120)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
