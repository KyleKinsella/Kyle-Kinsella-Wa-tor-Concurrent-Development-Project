package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rows := 30
	cols := 68
	grid := make([][]string, rows)	
	for i := range grid { // 0-29
		grid[i] = make([]string, cols) // Create a slice for each row
	}

	data := []string{"S", "F", "EW"}
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
				grid[i][j] = "_"
			}
		}
	}
	
	for i := range len(grid) {
		fmt.Println(grid[i])
	}
}