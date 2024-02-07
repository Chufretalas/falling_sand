package main

type BlockType int

const (
	BTAIR BlockType = iota
	BTSAND
)

type BlockGrid [][]BlockType

func (b *BlockGrid) init() {
	(*b) = make(BlockGrid, SCREENHEIGHT/squareSize[squareSizeIdx])
	for idx := range blocks {
		(*b)[idx] = make([]BlockType, SCREENWIDTH/squareSize[squareSizeIdx])
	}
}

func (b BlockGrid) clear() {
	for idx1 := range b {
		for idx2 := range b[idx1] {
			b[idx1][idx2] = BTAIR
		}
	}
}
