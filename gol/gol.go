package main

func calculateNextState(p golParams, world [][]byte) [][]byte {

	newWorld := make([][]byte, len(world))
	for i := range world {
		newWorld[i] = make([]byte, len(world[0]))
		copy(newWorld[i], world[i])
	}

	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			iBehind := (i - 1 + p.imageHeight) % p.imageHeight
			iAhead := (i + 1 + p.imageHeight) % p.imageHeight
			jBehind := (j - 1 + p.imageWidth) % p.imageWidth
			jAhead := (j + 1 + p.imageWidth) % p.imageWidth

			neighbourSum := int(world[iBehind][jBehind]) + int(world[iBehind][j]) + int(world[iBehind][jAhead]) + int(world[i][jBehind]) + int(world[i][jAhead]) + int(world[iAhead][jBehind]) + int(world[iAhead][j]) + int(world[iAhead][jAhead])
			liveNeighbours := neighbourSum / 255

			if world[i][j] == 255 {
				if (liveNeighbours < 2) || (liveNeighbours > 3) {
					newWorld[i][j] = 0
				}
			} else if world[i][j] == 0 {
				if liveNeighbours == 3 {
					newWorld[i][j] = 255
				}
			}
		}
	}

	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}

	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			if world[i][j] == 255 {
				aliveCells = append(aliveCells, cell{x: j, y: i})
			}
		}
	}

	return aliveCells
}
