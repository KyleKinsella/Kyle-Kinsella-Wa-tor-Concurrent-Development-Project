package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Reset = "\033[0m"
	Red = "\033[31m"
	Blue = "\033[34m"
	Black = "\033[30m"
	Yellow = "\033[33m"
)

var (
	Fish = Blue + "F" + Reset
	Shark string = Red + "S" + Reset
	EmptyWater string = Black + "_" + Reset

	grid  [][]string
	energy int
	gameOver bool

	rows         = 30
	cols         = 30
	GameOver = "Game Over"
)

func makeGrid(rows, cols int) [][]string {
		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols) 
			for j := range grid[i] {
				random := rand.Float64()
				if random < 0.2 { 
					grid[i][j] = Blue + "F" + Reset // fish
				} else if random < 0.4 {
					grid[i][j] = Red + "S" + Reset // shark
				} else { 
					grid[i][j] = Black + "_" + Reset // empty water
				}
			}
		}
		return grid
}

// Initialize the grid with default values and place my shark, fish and empty water
func initGrid() {
	grid = makeGrid(rows, cols)
	grid[0][0] = Shark // Place my shark
}

// Move the shark based on logic
func moveShark(grid [][]string, rows, cols int) {
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == Shark {
				// Randomize shark movement
				dx := rand.Intn(3) -1 // -1, 0, or 1
				time.Sleep(200)
				dy := rand.Intn(3) -1

				newX := x+dx
				newY := y+dy
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
					if grid[newX][newY] == Fish {
						// The shark eats the fish
						energy++
						energy = 100
						fmt.Println("Shark ate a fish! Energy level is:", energy)
					} 							
					// Move shark to new position
					grid[x][y] = EmptyWater
					grid[newX][newY] = Shark
				}
				// fish is not found deprive the shark of a unit of energy.
				if haveEmptyWater(grid, x, y) {
					if grid[x][y] != EmptyWater {
						energy = energy - 1

						// end the game because shark cannot have negative energy
						if energy < 0 {
							fmt.Println("shark is dead, due to having no energy")
							fmt.Println("shark cannot have negative energy")
							gameOver = true
							break
						}
						fmt.Println("Sharks energy has dropped to:", energy)
						break
					} 
				} 
				// return
			}
		}
	}
}

func haveEmptyWater(grid [][]string, rows, cols int) bool {
	for i := 0; i<rows; i++ {
		for j := 0; j<cols; j++ {
			if grid[i][j] == EmptyWater {
				return true
			}
		}
	}
	return false
}

// Move the shark based on logic
func moveFish(grid [][]string, rows, cols int) {
	for x := 1; x < rows-1; x++ {
		for y := 1; y < cols-1; y++ {
			if grid[x][y] == Fish {
				// Randomize shark movement
				dx := rand.Intn(3) - 1 // -1, 0, or 1
				time.Sleep(500)
				dy := rand.Intn(3) - 1

				newX := x+dx
				newY := y+dy
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
					if grid[newX][newY] == Fish {
						// The shark eats the fish
						energy++
						fmt.Println("fish has moved...")
					} 
					// Move shark to new position
					grid[x][y] = EmptyWater
					grid[newX][newY] = Fish
				}
				// return
			}
		}
	}
}

func main() {
	initGrid()
	moveShark(grid, rows, cols)
	moveFish(grid, rows, cols)
}