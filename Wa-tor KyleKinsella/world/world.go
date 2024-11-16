package main

import (
	"fmt"
	"math/rand"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Blue    = "\033[34m"
	Black   = "\033[30m"

	//testing only
	Yellow = "\033[33m"
)

func makeGrid(rows, cols int) [][]string {
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
		for j := range grid[i] {
			random := rand.Float64()
			if random < 0.1 { // 10% chance of a shark
				grid[i][j] = Red + "S" + Reset
			} else if random < 0.4 { // 40% chance of a fish
				grid[i][j] = Blue + "F" + Reset
			} else { // 50% chance of a empty water
				grid[i][j] = Black + "_" + Reset
			}
		}
	}
	return grid
}

func findAFish(grid[][]string) [][]int {
	pos := [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == Blue + "F" + Reset {
				// fmt.Println("Fish found at:", i, j)
				pos = append(pos, []int{i,j})
			}
		}
	}
	return pos
}

func gameMessage() {
	// Adding colors to strings with ANSI codes
	welcome := "\033[32m--------------------------Welcome to my Wa-tor Simulation project--------------------------" + Reset
	shark := "\nA shark is in the color Red and denoted with the letter " + Red + "S" + Reset + "."
	fish := "\nA fish is in the color Blue and denoted with the letter " + Blue + "F" + Reset + "."
	emptyWater := "\nThe empty water is in the color Black and denoted with the letter " + Black + "EW" + Reset + "."
	end := "\033[32m\n-------------------------------------------------------------------------------------------" + Reset
	fmt.Println(welcome, shark, fish, emptyWater, end)
}

func main() {
	rows := 2
	cols := 2
	grid := makeGrid(rows, cols)

	data := []string{Red + "S" + Reset, Blue + "F" + Reset, Black + "_" + Reset}	
	for i:=0; i<rows; i++ {
		for j:=0; j<cols; j++ {
			if grid[i][j] == "" {
				for i:=0; i<30; i++ {
					randomIn := rand.Intn(len(data))
					randomVal := data[randomIn]
					grid[i][j] = randomVal
				}
			}
		}
	}
	
	gameMessage()
	// fmt.Println("below is the inital grid")
	for i := range len(grid) {
		fmt.Println(grid[i])
	}	
	fishPositions := findAFish(grid)
	fmt.Println("Fish found at:", fishPositions)
}