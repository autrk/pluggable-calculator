/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/autrk/pluggable-calculator/internal/operation"
	extism "github.com/extism/go-sdk"
	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "substract",
	Short: "Subtracts number from the first number",
	Long:  `Performs subtraction by subtracting subsequent number from the first number.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run:   runSubstractCommand,
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

func runSubstractCommand(cmd *cobra.Command, args []string) {
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

	name := args[0]

	// Validierung des --type Flags
	validTypes := map[string]bool{
		"zig":  true,
		"rust": true,
		"ts":   true,
	}

	if !validTypes[pluginType] {
		fmt.Printf("Invalid plugin type: %s\n", pluginType)
		fmt.Println("Valid types are: zig, rust, ts")
		return
	}

	wasmFile := "./plugins/wasm/substract.wasm"

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFile,
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{EnableWasi: true}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	data := []byte(name)
	exit, out, err := plugin.Call("greet", data)
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	response := string(out)
	fmt.Println(response)
}
