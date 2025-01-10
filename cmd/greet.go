/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets a person by their name",
	Long: `The greet command welcomes a person by their provided name. Additionally, you can specify
which WebAssembly (WASM) plugin to use for generating the greeting. Available plugin types
are Zig, Rust, and TypeScript.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:  runGreetCommand,
}

var pluginType string

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().StringVarP(
		&pluginType, // Variable zur Speicherung des Flag-Werts
		"type",      // Flag-Name
		"t",         // Kurzform
		"zig",       // Standardwert
		`Specifies which WebAssembly (WASM) plugin to use for the greeting.
Available options are:
- zig: Uses the Zig-implemented plugin.
- rust: Uses the Rust-implemented plugin.
- ts: Uses the TypeScript-implemented plugin.`,
	)
}

func runGreetCommand(cmd *cobra.Command, args []string) {
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

	var wasmFile string

	switch pluginType {
	case "zig":
		wasmFile = "./plugins/wasm/greet-zig.wasm"
	case "rust":
		wasmFile = "./plugins/wasm//greet-rust.wasm"
	case "ts":
		wasmFile = "./plugins/wasm/greet-ts.wasm"
	}

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
