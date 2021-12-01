package cmd

import (
	"embed"
	"fmt"
	"github.com/maxslarsson/adventOfCodeTemplate/days/day01"
	"github.com/maxslarsson/adventOfCodeTemplate/input"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var inputs embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adventOfCode2021 <day>",
	Short: "My solutions to Advent of Code 2021, written in Go",
	//	Long: `A longer description that spans multiple lines and likely contains
	//examples and usage of using your application. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		day, err := strconv.Atoi(args[0])

		if err != nil {
			return err
		}

		inputPath := fmt.Sprintf("days/day%02d/input.txt", day)
		inputData, err := inputs.ReadFile(inputPath)
		if err != nil {
			return err
		}
		input := input.FromString(string(inputData))

		part1, part2 := getDay(day)

		if part1 != nil {
			fmt.Println("Running Part 1")
			start := time.Now()
			res := part1(input)
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Println(res)
			fmt.Println("Took", elapsed)
		}

		if part2 != nil {
			fmt.Println("Running Part 2")
			start := time.Now()
			res := part2(input)
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Println(res)
			fmt.Println("Took", elapsed)
		}

		return nil
	},
}

type DayFn = func(*input.Input) string

func getDay(day int) (DayFn, DayFn) {
	switch day {
	case 1:
		return day01.A, nil
	default:
		return nil, nil
	}
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(inputFiles embed.FS) {
	inputs = inputFiles
	cobra.CheckErr(rootCmd.Execute())
}