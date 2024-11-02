package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var wordToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide an input file")
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))

	var sumDigits, sumBoth int

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			digits, both := parseLine(line)
			sumDigits += getCalibrationValue(digits)
			sumBoth += getCalibrationValue(both)
		}
	}

	log.Printf("Part 1 - Sum of digits: %d", sumDigits)
	log.Printf("Part 2 - Sum of both: %d", sumBoth)
}

func getCalibrationValue(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	first, last := nums[0], nums[len(nums)-1]
	return first*10 + last
}

func parseLine(line string) ([]int, []int) {
	var digits, both []int

	for i := 0; i < len(line); i++ {
		// Check for digit first
		if c := line[i]; c >= '0' && c <= '9' {
			num := int(c - '0')
			both = append(both, num)
			digits = append(digits, num)
			continue
		}

		// Check for word numbers
		for word, num := range wordToInt {
			if strings.HasPrefix(line[i:], word) {
				both = append(both, num)
				// Shouldn't break - handle overlapping cases like "oneight"
			}
		}
	}

	return digits, both
}
