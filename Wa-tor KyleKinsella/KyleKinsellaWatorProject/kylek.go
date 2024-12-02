package main

import (
	"fmt"
	"math/rand"
	"time"
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"	
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
	GameOver = "Game Over"

	rows         = 30
	cols         = 30
	screenWidth  = 600
	screenHeight = 600
	cellSize     = screenWidth / cols
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

func emptyGrid(grid [][]string, rows, cols int) bool {
	for i := 0; i<rows; i++ {
		for j := 0; j<cols; j++ {
			if grid[i][j] == Fish {
				return false
			}
		}
	}
	return true
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

// Game implements the ebiten.Game interface
type Game struct{}

// Update is called 60 times per second
func (g *Game) Update() error {

	if gameOver {
		return nil
	}

	if emptyGrid(grid, rows, cols) {
		gameOver = true
		fmt.Println(GameOver)
	}

	moveShark(grid, rows, cols)
	moveFish(grid, rows, cols)
	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			rectX := y * cellSize
			rectY := x * cellSize

			var col color.Color
			switch grid[x][y] {
			case Shark:
				col = color.RGBA{255, 0, 0, 255} // Red 
			case Fish:
				col = color.RGBA{0, 0, 255, 255} // Blue
			default:
				col = color.RGBA{0, 0, 0, 0} // Black
			}

			// Draw cell as a filled rectangle
			ebitenutil.DrawRect(screen, float64(rectX), float64(rectY), float64(cellSize), float64(cellSize), col)

			if gameOver {
				ebitenutil.DebugPrint(screen, "Game Over")
				return
			}		
		}
	}
	// Display the sharks energy level
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Sharks energy is: %d", energy))
}

// Layout specifies the screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	initGrid()
	
	// Create a new game instance
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Kyle Kinsella | C00273146 Wa-Tor Simulation Project")
	game := &Game{}

	// run my wa-tor project
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}