package main

func calculateNextState(p golParams, world [][]byte) [][]byte {

	newWorld := make([][]byte, len(world))
	for i := range world {
		newWorld[i] = make([]byte, len(world[0]))
		copy(newWorld[i], world[i])
	}

	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			iBehind := i - 1
			iAhead := i + 1
			jBehind := j - 1
			jAhead := j + 1

			if i == 0 {
				iBehind = p.imageHeight - 1
			} else if i == p.imageHeight-1 {
				iAhead = 0
			}

			if j == 0 {
				jBehind = p.imageWidth - 1
			} else if j == p.imageWidth-1 {
				jAhead = 0

			}

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
