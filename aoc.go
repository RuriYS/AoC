package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RuriYS/AoC/solutions/cubeconundrum"
	"github.com/RuriYS/AoC/solutions/gearratios"
	"github.com/RuriYS/AoC/solutions/trebuchet"
	"github.com/spf13/cobra"
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
		inputPath = filepath.Join("solutions", s.solutionDir, inputPath)
	}

	return s.run(inputPath)
}

var solutions = map[string]*PuzzleSolution{
	"trebuchet":     {run: trebuchet.Run, solutionDir: "trebuchet"},
	"cubeconundrum": {run: cubeconundrum.Run, solutionDir: "cubeconundrum"},
	"gearratios":    {run: gearratios.Run, solutionDir: "gearratios"},
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "AoC <puzzle> [flags]",
		Short: "Advent of Code solutions runner",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = false
	rootCmd.CompletionOptions.DisableNoDescFlag = true

	// Add completion command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion [bash|zsh|fish]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:
  $ source <(aoc completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ aoc completion bash > /etc/bash_completion.d/aoc
  # macOS:
  $ aoc completion bash > /usr/local/etc/bash_completion.d/aoc

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ aoc completion zsh > "${fpath[1]}/_aoc"

  # You will need to start a new shell for this setup to take effect.

Fish:
  $ aoc completion fish | source

  # To load completions for each session, execute once:
  $ aoc completion fish > ~/.config/fish/completions/aoc.fish
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			}
		},
	})

	// Add command for each puzzle
	for name, solution := range solutions {
		puzzleName := name // Create a new variable to avoid closure issues
		puzzleSolution := solution

		cmd := &cobra.Command{
			Use:   puzzleName + " <input-file>",
			Short: fmt.Sprintf("Run the %s puzzle", puzzleName),
			Args:  cobra.MinimumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				return puzzleSolution.Run(args)
			},
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				if len(args) == 0 {
					// When completing the input file argument, look in both current directory
					// and solution directory for .txt files
					var files []string

					// Look in current directory
					matches, _ := filepath.Glob("*.txt")
					files = append(files, matches...)

					// Look in solution directory
					solutionPath := filepath.Join("solutions", puzzleSolution.solutionDir)
					if matches, err := filepath.Glob(filepath.Join(solutionPath, "*.txt")); err == nil {
						// Strip the path prefix for cleaner completion
						for _, match := range matches {
							files = append(files, filepath.Base(match))
						}
					}

					return files, cobra.ShellCompDirectiveDefault
				}
				return nil, cobra.ShellCompDirectiveDefault
			},
		}
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
