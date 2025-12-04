package main

import (
	"strconv"

	aoc_helpers "github.com/carlpatchett/advent-of-code-2025/inputs"
)

const K = 12 // digits to consider

func main() {
	input, _ := aoc_helpers.LoadInputsAsStrings("inputs.txt")

	total := 0
	for _, line := range input {
		best := make([]int, K+1) // best possible digits
		for i := 1; i <= K; i++ {
			best[i] = -1
		}
		best[0] = 0

		for i := 0; i < len(line); i++ {
			digit, _ := strconv.Atoi(string(line[i]))
			for l := K; l >= 1; l-- {
				if best[l-1] >= 0 {
					candidate := best[l-1]*10 + digit
					if candidate > best[l] {
						best[l] = candidate
					}
				}
			}
		}

		val := 0
		if best[K] >= 0 {
			val = best[K]
		} else {
			val = 0
		}

		total += val
	}

	println("Total:", total)
}
