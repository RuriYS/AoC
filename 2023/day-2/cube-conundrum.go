package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// represents the number of cubes of each color
type ColorCount struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

// Game represents a single game with its ID and rolls
type Game struct {
	GameID int            `json:"gameId"`
	Rolls  [][]ColorCount `json:"rolls"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Too few arguments\nUsage: %s [sample_file]", os.Args[0])
	}

	content, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	games, err := readSample(string(content))
	if err != nil {
		log.Fatalf("Failed to parse games: %v", err)
	}

	sum := 0
	power := 0

	for id, game := range games {
		if isGamePossible(game, 12, 13, 14) {
			sum += id
		}

		red, green, blue := getMinimumCubes(game)
		power += red * green * blue
	}
	fmt.Printf("Part 1 - Sum of possible game IDs: %d\n", sum)
	fmt.Printf("Part 2 - Sum of powers: %d\n", power)
}

func readSample(content string) (map[int]Game, error) {
	games := make(map[int]Game)
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Split "Game 1: ..." into ["Game 1", "..."]
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Extract game ID
		gameIDStr := strings.TrimPrefix(parts[0], "Game ")
		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			return nil, fmt.Errorf("invalid game ID: %s", gameIDStr)
		}

		// Split rolls (separated by semicolons)
		rollStrings := strings.Split(parts[1], "; ")
		rolls := make([][]ColorCount, len(rollStrings))

		for i, rollStr := range rollStrings {
			// Each roll only has one ColorCount in this case
			colorCount := ColorCount{}

			// Split color counts (e.g., "8 green, 6 blue, 1 red")
			colorParts := strings.Split(rollStr, ", ")
			for _, colorPart := range colorParts {
				// Split count and color (e.g., "8 green")
				countColor := strings.Split(colorPart, " ")
				if len(countColor) != 2 {
					return nil, fmt.Errorf("invalid color count format: %s", colorPart)
				}

				count, err := strconv.Atoi(countColor[0])
				if err != nil {
					return nil, fmt.Errorf("invalid count: %s", countColor[0])
				}

				// Assign count to appropriate color
				switch countColor[1] {
				case "red":
					colorCount.Red = count
				case "green":
					colorCount.Green = count
				case "blue":
					colorCount.Blue = count
				default:
					return nil, fmt.Errorf("invalid color: %s", countColor[1])
				}
			}

			rolls[i] = []ColorCount{colorCount}
		}

		games[gameID] = Game{
			GameID: gameID,
			Rolls:  rolls,
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return games, nil
}

func isGamePossible(game Game, maxRed, maxGreen, maxBlue int) bool {
	for _, rolls := range game.Rolls {
		for _, count := range rolls {
			if count.Red > maxRed || count.Green > maxGreen || count.Blue > maxBlue {
				return false
			}
		}
	}
	return true
}

func getMinimumCubes(game Game) (red, green, blue int) {
	for _, roll := range game.Rolls {
		for _, c := range roll {
			red, green, blue = max(red, c.Red), max(green, c.Green), max(blue, c.Blue)
		}
	}
	return
}
