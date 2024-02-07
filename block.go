package main

type BlockType int

const (
	BTAIR BlockType = iota
	BTSAND
)

type BlockGrid [][]BlockType

func (b *BlockGrid) init() {
	(*b) = make(BlockGrid, SCREENHEIGHT/squareSide)
	for idx := range blocks {
		(*b)[idx] = make([]BlockType, SCREENWIDTH/squareSide)
	}
}

func (b BlockGrid) clear() {
	for idx1 := range b {
		for idx2 := range b[idx1] {
			b[idx1][idx2] = BTAIR
		}
	}
}
