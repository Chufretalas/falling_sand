package main

import "math/rand"

// TODO: optimize this with goroutines if possible
func updateblocks() {

	blocksCopy.clear()

	for iy := len(blocks) - 1; iy >= 0; iy-- {
		for ix, block := range blocks[iy] {
			if block == BTSAND {
				if iy+1 != len(blocks) {
					if blocks[iy+1][ix] == BTAIR {
						if blocksCopy[iy+1][ix] == BTAIR {
							blocksCopy[iy+1][ix] = block
						} else {
							blocksCopy[iy][ix] = block
						}
					} else {
						// start checking all sand blocks bellow
						current_y := iy + 2
						for {
							if current_y >= len(blocks)-1 {
								// its in a column touching the ground
								leftEmpty := false
								rightEmpty := false
								if ix != 0 {
									if blocks[iy][ix-1] == BTAIR && blocksCopy[iy][ix-1] == BTAIR {
										leftEmpty = true
									}
								}
								if ix != len(blocks[iy])-1 {
									if blocks[iy][ix+1] == BTAIR && blocksCopy[iy][ix+1] == BTAIR {
										rightEmpty = true
									}
								}

								if leftEmpty || rightEmpty {
									canFallLeft := false
									canFallRight := false
									if leftEmpty {
										if blocks[iy+1][ix-1] == BTAIR && blocksCopy[iy+1][ix-1] == BTAIR {
											canFallLeft = true
										}
									}
									if rightEmpty {
										if blocks[iy+1][ix+1] == BTAIR && blocksCopy[iy+1][ix+1] == BTAIR {
											canFallRight = true
										}
									}

									if canFallLeft && canFallRight {
										switch n := rand.Intn(2); n {
										case 0:
											blocksCopy[iy][ix-1] = block
										case 1:
											blocksCopy[iy][ix+1] = block
										}
										break
									}

									if canFallLeft {
										blocksCopy[iy][ix-1] = block
										break
									}

									if canFallRight {
										blocksCopy[iy][ix+1] = block
										break
									}
								}

								blocksCopy[iy][ix] = block
								break
							}
							// TODO: if sand blocks start eating other blocks, this might be the problem, but it seens fine now
							if blocksCopy[current_y][ix] == BTAIR {
								if blocksCopy[iy+1][ix] == BTAIR {
									blocksCopy[iy+1][ix] = block
								}
								break
							}
							current_y++
						}
					}
				} else {
					blocksCopy[iy][ix] = block
				}
			}
		}
	}

	for idx1 := range blocks {
		copy(blocks[idx1], blocksCopy[idx1])
	}

}
