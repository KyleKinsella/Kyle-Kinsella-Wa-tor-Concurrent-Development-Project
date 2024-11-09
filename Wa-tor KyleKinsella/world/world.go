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
)

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
	rows := 30
	cols := 60
	grid := make([][]string, rows)	
	for i := range grid { // 0-29
		grid[i] = make([]string, cols) // Create a slice for each row
	}

	data := []string{Red + "S" + Reset, Blue + "F" + Reset, Black + "EW" + Reset}	
	// this code puts a random value into the 2d-array
	for _, value := range data {
		place := false
		for !place {
			placeRandomRow := rand.Intn(rows)
			placeRandomCol := rand.Intn(cols)

			if grid[placeRandomRow][placeRandomCol] == "" {
				grid[placeRandomRow][placeRandomCol] = value
				place = true
			} else {
				place = false
			}
		}
	}

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
	for i := range len(grid) {
		fmt.Println(grid[i])
	}
}