package main

import (
	"fmt"
	"strconv"

	aoc_helpers "github.com/carlpatchett/advent-of-code-2025/inputs"
)

func main() {
	lines, _ := aoc_helpers.LoadInputsAsStrings("inputs.txt")

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	for i := range lines {
		for len(lines[i]) < maxLen {
			lines[i] += " "
		}
	}

	nRows := len(lines)

	// Track numbers and operator for current problem
	problemNumbers := []int{}
	var currentOperator byte

	grandTotal := 0

	// Process columns from right to left
	for ci := maxLen - 1; ci >= 0; ci-- {
		operatorChar := lines[nRows-1][ci]

		// Check if this is a separator column (all spaces)
		allSpaces := true
		for ri := range nRows {
			if lines[ri][ci] != ' ' {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			// End of a problem - calculate result
			if currentOperator != 0 && len(problemNumbers) > 0 {
				result := problemNumbers[0]
				fmt.Printf("Problem: %d", result)
				for i := 1; i < len(problemNumbers); i++ {
					fmt.Printf(" %c %d", currentOperator, problemNumbers[i])
					switch currentOperator {
					case '+':
						result += problemNumbers[i]
					case '*':
						result *= problemNumbers[i]
					}
				}
				fmt.Printf(" = %d\n", result)
				grandTotal += result
			}

			problemNumbers = []int{}
			currentOperator = 0
			continue
		}

		if operatorChar == '+' || operatorChar == '*' {
			currentOperator = operatorChar
		}

		numStr := ""
		for ri := range lines[:nRows-1] {
			ch := lines[ri][ci]
			if ch != ' ' {
				numStr += string(ch)
			}
		}

		if numStr != "" {
			val, _ := strconv.Atoi(numStr)
			problemNumbers = append(problemNumbers, val)
		}
	}

	if currentOperator != 0 && len(problemNumbers) > 0 {
		result := problemNumbers[0]
		fmt.Printf("Problem: %d", result)
		for i := 1; i < len(problemNumbers); i++ {
			fmt.Printf(" %c %d", currentOperator, problemNumbers[i])
			switch currentOperator {
			case '+':
				result += problemNumbers[i]
			case '*':
				result *= problemNumbers[i]
			}
		}
		fmt.Printf(" = %d\n", result)
		grandTotal += result
	}

	fmt.Println("\nGrand total is", grandTotal)
}
