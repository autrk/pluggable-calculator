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

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divides the first number by the subsequent number",
	Long:  `Performs division by dividing the first number by the subsequent number.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		a, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}

		b, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			panic(err)
		}

		sum, err := operation.Divide(a, b)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Division of %s is %.3f\n", args, sum)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
