package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/p886/5-million-steps/calculator"
	"github.com/p886/5-million-steps/representer"
	"github.com/spf13/cobra"
)

// RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "./5-million-steps [steps so far]",
	Short: "Am I on track for 5 Million Steps in 2020?",
	Long:  `This little CLI app tells me if I'm on track for my goal of 5 Million Steps in 2020`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currentSteps, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error converting input steps: %v", err)
		}
		result := calculator.Calculate(currentSteps)
		textRepresentation := representer.AsText(result)
		fmt.Println(textRepresentation)
	},
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
