package scratchcards

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(inputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	points, totalCards, err := readSample(string(content))
	if err != nil {
		return err
	}

	fmt.Printf("Total points: %d\n", points)
	fmt.Printf("Total cards: %d\n", totalCards)
	return nil
}

func readSample(content string) (int, int, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	var lines []string
	var points int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	// Create instances to track each card
	cardInstances := make([]int, len(lines))
	for i := range cardInstances {
		cardInstances[i] = 1
	}

	for cardNum, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return 0, 0, fmt.Errorf("invalid line format: %s", line)
		}

		numbers := strings.Split(parts[1], "|")
		winningNumbers := strings.Fields(numbers[0])
		currentNumbers := strings.Fields(numbers[1])

		// Find matches
		numberMap := make(map[int]bool)
		for _, numStr := range winningNumbers {
			num, _ := strconv.Atoi(numStr)
			numberMap[num] = true
		}

		matches := 0
		for _, numStr := range currentNumbers {
			num, _ := strconv.Atoi(numStr)
			if numberMap[num] {
				matches++
			}
		}

		// Sum the points, then for each instance of current card, create copies of next cards
		if matches > 0 {
			points += 1 << (matches - 1)
			currentInstances := cardInstances[cardNum]
			for i := 1; i <= matches; i++ {
				nextCard := cardNum + i
				if nextCard < len(cardInstances) {
					cardInstances[nextCard] += currentInstances
				}
			}
		}
	}

	totalCards := 0
	for _, instances := range cardInstances {
		totalCards += instances
	}

	return points, totalCards, nil
}
