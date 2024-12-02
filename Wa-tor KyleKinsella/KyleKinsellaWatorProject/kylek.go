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
	
	rows         = 30
	cols         = 30
	screenWidth  = 600
	screenHeight = 600
	threads = 0
	cellSize     = screenWidth / cols
)

// this function will make a 2d grid of fish, sharks and empty water
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

// Initialize the grid with default values and place my shark
func initGrid() {
	grid = makeGrid(rows, cols)
	grid[0][0] = Shark // Place my shark
	// give my shark an energy level
	energy = 300
}

// this function moves my shark within my 2d grid
func moveShark(grid [][]string, rows, cols int) {
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == Shark {
				// Randomize shark movement
				dx := rand.Intn(3) -1 // -1, 0, or 1
				time.Sleep(200)
				dy := rand.Intn(3) -1

				// here I add the value of x to the random spot to where the shark has moved(dx)
				newX := x+dx
				// here I add the value of y to the random spot to where the shark has moved(dy)
				newY := y+dy

				// Perform a boundary check to ensure the coordinates are within the 2D grid
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
					// this if statement check to see if a shark has found a fish within the grid
					if grid[newX][newY] == Fish {
						// The shark eats the fish, give him so energy
						energy++
						fmt.Println("Shark ate a fish! Energy level is:", energy)
					} 							
					// Move shark to new position
					grid[x][y] = EmptyWater
					grid[newX][newY] = Shark
				}
				// fish is not found deprive the shark of a unit of energy.
				if haveEmptyWater(grid, x, y) { // if this function gives us true then I do the following
					// if we are at a spot in the grid and it is not empty water, then we have not found a fish so remove some 
					// energy from the shark
					if grid[x][y] != EmptyWater {
						// remove some energy from the shark
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
			}
		}
	}
}

// this function checks to see if my grid has empty water
func haveEmptyWater(grid [][]string, rows, cols int) bool {
	for i := 0; i<rows; i++ {
		for j := 0; j<cols; j++ {
			// I check to see if there is empty water at an index within the 2d grid
			if grid[i][j] == EmptyWater {
				// yes there is empty water
				return true
			}
		}
	}
	// no there is not empty water 
	return false
}

// this function checks to see if my grid has a fish
func emptyGrid(grid [][]string, rows, cols int) bool {
	for i := 0; i<rows; i++ {
		for j := 0; j<cols; j++ {
			// I check to see if there is a fish in the 2d grid
			if grid[i][j] == Fish {
				// grid is not empty
				return false
			}
		}
	}
	// grid is empty
	return true
}

// this function moves my fish within my 2d grid
func moveFish(grid [][]string, rows, cols int) {
	// here I am doing a diffrent loop compared to my move shark, this is because if I have my move fish the same as my move shark,
	// my fish were eating my sharks, I tried to fix this for hours but could not fix this, this is why my loops are diffrent. As a result this means 
	// my fish on the edges do not move.
	for x := 1; x < rows-1; x++ {
		for y := 1; y < cols-1; y++ {
			if grid[x][y] == Fish {
				// Randomize fish movement
				dx := rand.Intn(3) - 1 // -1, 0, or 1
				time.Sleep(500)
				dy := rand.Intn(3) - 1
				
				// here I add the value of x to the random spot to where the shark has moved(dx)
				newX := x+dx
				// here I add the value of y to the random spot to where the shark has moved(dy)
				newY := y+dy

				// Perform a boundary check to ensure the coordinates are within the 2D grid
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
					// this if statement check to see if a fish is in the grid
					if grid[newX][newY] == Fish {
						// fish has moved
						fmt.Println("fish has moved...")
					} 
					// Move fish to new position
					grid[x][y] = EmptyWater
					grid[newX][newY] = Fish
				}
			}
		}
	}
}

// Game implements the ebiten.Game interface
type Game struct{}

// Update is called 60 times per second
func (g *Game) Update() error {

	// if my gameOver variable is true then end the game 
	if gameOver {
		return nil
	}

	// I call my emptyGrid function, if this returns a true then we end the game
	if emptyGrid(grid, rows, cols) {
		gameOver = true
	}

	// call my shark function
	moveShark(grid, rows, cols)
	// call my fish function
	moveFish(grid, rows, cols)
	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			// here I calculate the size of the cell for each y value
			rectX := y * cellSize
			// here I calculate the size of the cell for each x value
			rectY := x * cellSize

			// this is how I set up the color for my shark, fish and empty water
			var col color.Color

			switch grid[x][y] {
			case Shark:
				col = color.RGBA{255, 0, 0, 255} // Red: shark
			case Fish:
				col = color.RGBA{0, 0, 255, 255} // Blue: fish
			default:
				col = color.RGBA{0, 0, 0, 0} // Black: empty water
			}

			// Draw cell as a filled rectangle
			ebitenutil.DrawRect(screen, float64(rectX), float64(rectY), float64(cellSize), float64(cellSize), col)

			// if this is true show game over screen to user
			if gameOver {
				ebitenutil.DebugPrint(screen, "Game Over")
				return
			}		
		}
	}
	// Display the sharks energy level
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Sharks energy is: %d", energy))
}

// this function returns the size of the screen width and height
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// call the initGrid method to set up my grid 
	initGrid()

	fmt.Printf("How big do you want your screen width to be?\n")
	fmt.Scan(&screenWidth)
	userInputWidth := screenWidth

	fmt.Println("How big do you want your screen height to be?")
	fmt.Scan(&screenHeight)
	userInputHeight := screenHeight

	var userInputRows int
	fmt.Println("How many rows do you want in your grid?")
	fmt.Scan(&userInputRows)
	userInputRows = rows

	var userInputCols int
	fmt.Println("How many cols do you want in your grid?")
	fmt.Scan(&userInputCols)
	userInputCols = cols

	var userInputThreads int
	fmt.Println("How many threads do you want to use?")
	fmt.Scan(&threads)
	userInputThreads = threads

	fmt.Println("--------------------------------------------")
	fmt.Println("The width of the screen is:", userInputWidth)
	fmt.Println("The height of the screen is:", userInputHeight)
	fmt.Println("Amount of rows is:", userInputRows)
	fmt.Println("Amount of cols is:", userInputCols)
	fmt.Println("Amount of threads in use is:", userInputThreads)
	fmt.Println("--------------------------------------------")

	for runtime := 0; runtime < userInputThreads; runtime++ {
		// here i get the current local time 
		now := time.Now()

		// I show how long it takes to run the program at the end of the main function
		defer func() {
			fmt.Println("time it took to run", userInputThreads, "thread(s) was", time.Since(now))
		}()
	}

	// Create a new game instance
	ebiten.SetWindowSize(userInputWidth, userInputHeight)
	ebiten.SetWindowTitle("Kyle Kinsella | C00273146 Wa-Tor Simulation Project")
	
	// make a game object, this uses the ebiten game interface
	game := &Game{}

	// run my wa-tor project
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}