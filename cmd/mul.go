/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/autrk/pluggable-calculator/internal/operation"
	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "Multiplies two numbers",
	Long: `Performs multiplication on two numbers.`,
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		a, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}

		b, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			panic(err)
		}

		sum := operation.Multiply(a, b)
		fmt.Printf("Multiplication of %s is %.3f\n", args, sum)
	},
}

func init() {
	rootCmd.AddCommand(mulCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
