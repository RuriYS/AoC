package gearratios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var matrix [][]rune

type NumberLocation struct {
	value    int
	row      int
	startCol int
	endCol   int
}

func Run(inputFile string) error {
	content, err := loadSample(inputFile)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	parseSample(scanner)

	// Store all number locations and track part 1 sum
	numbers := make([]NumberLocation, 0)
	part1Sum := 0

	// First pass: find all numbers and calculate part 1
	for row := 0; row < len(matrix); row++ {
		currentNum := 0
		startCol := -1

		for col := 0; col < len(matrix[row]); col++ {
			if unicode.IsDigit(matrix[row][col]) {
				if startCol == -1 {
					startCol = col
				}
				digit, _ := strconv.Atoi(string(matrix[row][col]))
				currentNum = currentNum*10 + digit
			} else if startCol != -1 {
				// Store number location
				numLoc := NumberLocation{
					value:    currentNum,
					row:      row,
					startCol: startCol,
					endCol:   col - 1,
				}
				numbers = append(numbers, numLoc)

				// Part 1 calculation
				if checkSurroundings(matrix, row, startCol, col-1) {
					part1Sum += currentNum
				}

				currentNum = 0
				startCol = -1
			}
		}

		// Handle number at end of line
		if startCol != -1 {
			numLoc := NumberLocation{
				value:    currentNum,
				row:      row,
				startCol: startCol,
				endCol:   len(matrix[row]) - 1,
			}
			numbers = append(numbers, numLoc)

			if checkSurroundings(matrix, row, startCol, len(matrix[row])-1) {
				part1Sum += currentNum
			}
		}
	}

	// Part 2: Find gear ratios
	part2Sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == '*' {
				// Find adjacent numbers
				adjacentNums := findAdjacentNumbers(numbers, row, col)
				if len(adjacentNums) == 2 {
					// Calculate gear ratio
					part2Sum += adjacentNums[0].value * adjacentNums[1].value
				}
			}
		}
	}

	fmt.Printf("Sum of part numbers: %d\n", part1Sum)
	fmt.Printf("Sum of gear ratios: %d\n", part2Sum)
	return nil
}

func loadSample(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to load file: %s", err)
	}
	return string(content), nil
}

func parseSample(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}
}

func checkSurroundings(matrix [][]rune, row, startCol, endCol int) bool {
	for r := max(0, row-1); r <= min(len(matrix)-1, row+1); r++ {
		for c := max(0, startCol-1); c <= min(len(matrix[0])-1, endCol+1); c++ {
			if isSymbol(matrix[r][c]) {
				return true
			}
		}
	}
	return false
}

func findAdjacentNumbers(numbers []NumberLocation, row, col int) []NumberLocation {
	adjacent := make([]NumberLocation, 0)

	for _, num := range numbers {
		// Check if the number is adjacent to the position
		if isAdjacent(num, row, col) {
			adjacent = append(adjacent, num)
		}
	}

	return adjacent
}

// isAdjacent checks if a number is adjacent to a position
func isAdjacent(num NumberLocation, row, col int) bool {
	// Check if the position is within one space of the number
	return num.row >= row-1 &&
		num.row <= row+1 &&
		col >= num.startCol-1 &&
		col <= num.endCol+1
}

func isSymbol(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
