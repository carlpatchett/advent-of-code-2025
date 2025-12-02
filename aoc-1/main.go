package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	lines, err := loadInput("inputs.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load inputs:", err)
		os.Exit(1)
	}

	fmt.Println("Loaded", len(lines), "lines from inputs.txt")

	dial := Dial{Current: 50}
	passes := dial.decodeInput(lines)
	fmt.Println("Total passes:", passes)
}

func loadInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

type Dial struct {
	Current int
}

func (d *Dial) decode(input []string) int {

	total := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		if line[0] == 'L' {
			numTurns, _ := strings.CutPrefix(line, "L")
			numTurnsInt, _ := strconv.Atoi(numTurns)
			for range numTurnsInt {
				d.turnLeft()

				// Click per number
				if d.Current == 0 {
					total++
				}
			}

			// Clicks after turns
			if d.Current == 0 {
				total++
			}

		} else {
			numTurns, _ := strings.CutPrefix(line, "R")
			numTurnsInt, _ := strconv.Atoi(numTurns)
			for range numTurnsInt {
				d.turnRight()

				if d.Current == 0 {
					total++
				}
			}
		}
	}

	return total
}

func (d *Dial) turnLeft() {
	if d.Current == 99 {
		d.Current = 0
		return
	}

	d.Current++
}

func (d *Dial) turnRight() {
	if d.Current == 0 {
		d.Current = 99
		return
	}

	d.Current--
}
func (d *Dial) decodeInput(input []string) int {
	return d.decode(input)
}
