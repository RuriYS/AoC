package trebuchet

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(inputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
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

	fmt.Printf("Sum of digits: %d\n", sumDigits)
	fmt.Printf("Sum of both: %d\n", sumBoth)
	return nil
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
