package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	aoc_helpers "github.com/carlpatchett/advent-of-code-2025/inputs"
)

func main() {
	input, err := aoc_helpers.LoadInputsAsStrings("inputs.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load inputs:", err)
		os.Exit(1)
	}
	ranges := make([][2]int, 0)

	for _, line := range input {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			min, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			max, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

			ranges = append(ranges, [2]int{min, max})
		}
	}

	if len(ranges) == 0 {
		fmt.Println(0)
		return
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })

	var total int64 = 0
	curStart := ranges[0][0]
	curEnd := ranges[0][1]

	for i := 1; i < len(ranges); i++ {
		s := ranges[i][0]
		e := ranges[i][1]
		if s <= curEnd+1 { // overlapping
			if e > curEnd {
				curEnd = e
			}
		} else {
			total += int64(curEnd - curStart + 1)
			curStart = s
			curEnd = e
		}
	}

	total += int64(curEnd - curStart + 1)

	fmt.Println(total)
}
