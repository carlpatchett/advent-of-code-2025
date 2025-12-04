package main

import (
	"fmt"

	aoc_helpers "github.com/carlpatchett/advent-of-code-2025/inputs"
)

func main() {
	input, _ := aoc_helpers.LoadInputsAsStrings("inputs.txt")
	grid := make([][]string, 0)

	for _, line := range input {
		row := make([]string, 0)
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	rollsThatCanBeAccessed := 0

	for {
		rollsRemoved, updatedGrid := removeRolls(grid)
		if rollsRemoved == 0 {
			break
		}

		grid = updatedGrid
		rollsThatCanBeAccessed += rollsRemoved
	}

	fmt.Println("Rolls that can be accessed:", rollsThatCanBeAccessed)
}

func removeRolls(grid [][]string) (int, [][]string) {
	rollsRemoved := 0
	for rowIndex, row := range grid {
		for colIndex, symbol := range row {
			if symbol == "." {
				continue
			}

			aboveRow := rowIndex
			if rowIndex > 0 {
				aboveRow = rowIndex - 1
			}

			belowRow := rowIndex
			if rowIndex < len(grid)-1 {
				belowRow = rowIndex + 1
			}

			leftCol := colIndex
			if colIndex > 0 {
				leftCol = colIndex - 1
			}

			rightCol := colIndex
			if colIndex < len(row)-1 {
				rightCol = colIndex + 1
			}

			rollsSeen := 0
			for r := aboveRow; r <= belowRow; r++ {
				for c := leftCol; c <= rightCol; c++ {
					if r == rowIndex && c == colIndex {
						continue
					}
					if grid[r][c] == "@" {
						rollsSeen++
					}
				}
			}

			if rollsSeen < 4 {
				rollsRemoved++
				grid[rowIndex][colIndex] = "."
			}
		}
	}
	return rollsRemoved, grid
}
