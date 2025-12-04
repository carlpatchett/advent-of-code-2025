package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	aoc_helpers "github.com/carlpatchett/advent-of-code-2025/inputs"
)

func main() {

	lines, err := aoc_helpers.LoadInputsAsStrings("inputs.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load inputs:", err)
		os.Exit(1)
	}

	invalidCodes := make([]string, 0)

	parts := strings.Split(lines[0], ",")
	for _, part := range parts {
		codes := strings.Split(part, "-")
		left, _ := strconv.Atoi(codes[0])
		right, _ := strconv.Atoi(codes[1])

		for i := left; i <= right; i++ {
			isInvalid, err := isInvalid(i)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error checking code:", err)
				os.Exit(1)
			}
			if isInvalid {
				invalidCodes = append(invalidCodes, strconv.Itoa(i))
				fmt.Println("Code is invalid:", i)
			}
		}
	}

	result := 0
	for _, code := range invalidCodes {
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to convert code to int:", err)
			os.Exit(1)
		}

		result += codeInt
	}

	fmt.Println("Sum of invalid codes:", result)
}

func isInvalid(code int) (bool, error) {
	s := strconv.Itoa(code)

	/* Checking if the string is made of two identical halves

	if len(s)%2 != 0 {
		return false, nil
	}

	half := len(s) / 2
	return s[:half] == s[half:], nil
	*/
	concat := s + s
	trimmed := concat[1 : len(concat)-1]

	// Check if s is a substring of trimmed
	return strings.Contains(trimmed, s), nil
}
