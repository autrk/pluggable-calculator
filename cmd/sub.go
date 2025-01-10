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

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtracts number from the first number",
	Long: `Performs subtraction by subtracting subsequent number from the first number.`,
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

		sum := operation.Subtract(a, b)
		fmt.Printf("Substraction of %s is %.3f\n", args, sum)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
