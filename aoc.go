package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/RuriYS/AoC/solutions/cubeconundrum"
	"github.com/RuriYS/AoC/solutions/gearratios"
	"github.com/RuriYS/AoC/solutions/trebuchet"
)

type PuzzleSolution struct {
	run         func(string) error
	solutionDir string
}

func (s *PuzzleSolution) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide an input file")
	}

	inputPath := args[0]
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		solutionPath := filepath.Join("solutions", s.solutionDir, inputPath)
		if _, err := os.Stat(solutionPath); err == nil {
			inputPath = solutionPath
		}
	}

	return s.run(inputPath)
}

var solutions = map[string]*PuzzleSolution{
	"trebuchet":     {run: trebuchet.Run, solutionDir: "trebuchet"},
	"cubeconundrum": {run: cubeconundrum.Run, solutionDir: "cubeconundrum"},
	"gearratios":    {run: gearratios.Run, solutionDir: "gearratios"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc <puzzle> [args...]")
		fmt.Println("Available puzzles:")
		for name := range solutions {
			fmt.Printf("  - %s\n", name)
		}
		os.Exit(1)
	}

	puzzleName := os.Args[1]
	solution, exists := solutions[puzzleName]
	if !exists {
		log.Fatalf("Unknown puzzle: %s", puzzleName)
	}

	if err := solution.Run(os.Args[2:]); err != nil {
		log.Fatal(err)
	}
}
