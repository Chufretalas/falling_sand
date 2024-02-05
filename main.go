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

type Block struct {
	bType BlockType
	pos   Position
}

type BlockGrid [][]Block

var (
	blocks      BlockGrid
	blocksCopy  BlockGrid
	activeBlock Position
)

const (
	SCREENWIDTH  = 1920
	SCREENHEIGHT = 1080
	SQUARESIDE   = 20
)

func init() {
	blocks = make([][]Block, SCREENHEIGHT/SQUARESIDE)
	blocksCopy = make([][]Block, SCREENHEIGHT/SQUARESIDE)
	for idx := range blocks {
		blocks[idx] = make([]Block, SCREENWIDTH/SQUARESIDE)
		blocksCopy[idx] = make([]Block, SCREENWIDTH/SQUARESIDE)
	}

	activeBlock = Position{0, 0}
}

func blockPos2BlocksIdx(p Position) (int, int) {
	return int(p.y / SQUARESIDE), int(p.x / SQUARESIDE)
}

func clearBlocks() {
	for idx1 := range blocks {
		for idx2 := range blocks[idx1] {
			blocks[idx1][idx2] = Block{}
		}
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		go ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("ahahaha")
	}

	mouseX, mouseY := ebiten.CursorPosition()
	activeBlock.x = float32(mouseX) - float32(mouseX%SQUARESIDE)
	activeBlock.y = float32(mouseY) - float32(mouseY%SQUARESIDE)

	blocksIdx1, blocksIdx2 := blockPos2BlocksIdx(activeBlock)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		blocks[blocksIdx1][blocksIdx2].bType = BTSAND
		blocks[blocksIdx1][blocksIdx2].pos.x = activeBlock.x
		blocks[blocksIdx1][blocksIdx2].pos.y = activeBlock.y
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		clearBlocks()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\n", ebiten.ActualFPS()))

	for y := 0; y <= SCREENHEIGHT; y += SQUARESIDE {
		vector.StrokeLine(screen, 0, float32(y), SCREENWIDTH, float32(y), 1, color.RGBA{20, 20, 20, 10}, true)
	}
	for x := 0; x <= SCREENWIDTH; x += SQUARESIDE {
		vector.StrokeLine(screen, float32(x), 0, float32(x), SCREENHEIGHT, 1, color.RGBA{20, 20, 20, 10}, true)
	}

	for idx1 := range blocks {
		for _, block := range blocks[idx1] {
			if block.bType != BTAIR {
				vector.DrawFilledRect(screen, block.pos.x, block.pos.y, SQUARESIDE, SQUARESIDE, color.RGBA{200, 100, 100, 255}, true)
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
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
