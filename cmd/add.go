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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds two numbers",
	Long:  `Performs addition on two numbers.`,
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

		sum := operation.Add(a, b)
		fmt.Printf("Addition of %s is %.3f\n", args, sum)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
