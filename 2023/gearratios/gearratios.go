package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var matrix [][]rune

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Too few arguments\nUsage: %s [sample_file]", os.Args[0])
	}

	content := loadSample(os.Args[1])
	scanner := bufio.NewScanner(strings.NewReader(content))
	parseSample(scanner)

	// matrix[0] = ['4', '6', '7', '.', '.', '1', '1', '4', '.', '.']
	// matrix[1] = ['.', '.', '.', '*', '.', '.', '.', '.', '.', '.']
	// matrix[2] = ['.', '.', '3', '5', '.', '.', '6', '3', '3', '.']
	// etc...

	sum := 0
	for row := 0; row < len(matrix); row++ {
		currentNum := 0
		startCol := -1

		for col := 0; col < len(matrix[row]); col++ {
			if unicode.IsDigit(matrix[row][col]) {

				// Start of a number or continuing a number
				if startCol == -1 {
					startCol = col
				}

				// Build the number
				digit, _ := strconv.Atoi(string(matrix[row][col]))
				currentNum = currentNum*10 + digit
			} else if startCol != -1 {

				// End of a number - check its surroundings
				if checkSurroundings(matrix, row, startCol, col-1) {
					sum += currentNum
				}

				// Reset for next number
				currentNum = 0
				startCol = -1
			}
		}

		// Check if number ends at end of line
		if startCol != -1 {
			if checkSurroundings(matrix, row, startCol, len(matrix[row])-1) {
				sum += currentNum
			}
		}
	}
	
	fmt.Printf("Sum of part numbers: %d\n", sum)
}

func loadSample(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to load file: %s", err)
	}
	return string(content)
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
